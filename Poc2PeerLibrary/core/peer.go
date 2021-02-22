package core

import "C"
import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/protocol"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
	"github.com/golang/glog"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	inet "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const httpEndpoint = "http://0.0.0.0:5001/ID"

type LibP2pCore struct {
	N            p2pnetwork.Network
	Node         host.Host
	peerid       peer.ID
	ctx          context.Context
	infos        *p2pnetwork.NetworkInfos
	tracker      Tracker
	LocalStorage storage.P2PStorage
	streams      map[peer.ID]p2pnetwork.WrappedStream
}

// NewLibP2P creates a LibP2P host with a random peer ID listening on the
// given multiaddress.
func NewLibP2P(infos *p2pnetwork.NetworkInfos) (*LibP2pCore, error) {
	libCore := &LibP2pCore{
		infos: infos,
	}

	// One peerId is acquired from Permanent peer, intialize p2p node.
	if err := libCore.InitNode(); err != nil {
		return nil, err
	}
	return libCore, nil
}

// MakeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress.
func NewP2PPeer(tracker Tracker, infos *p2pnetwork.NetworkInfos) (*LibP2pCore, error) {
	libCore, err := NewLibP2P(infos)
	if err != nil {
		return nil, err
	}

	libCore.tracker = tracker

	//Request Peer Id from Http Endpoint
	requestedPeerId, err := libCore.RequestPermanentPeerID()
	if err != nil {
		return nil, err
	}

	if err = libCore.DecodePermanentPeerID(requestedPeerId); err != nil {
		return nil, err
	}
	return libCore, nil
}

func (c *LibP2pCore) InitNode() (err error) {
	c.ctx = context.Background()

	opts := c.getNodeOptions()

	if c.Node, err = libp2p.New(c.ctx, opts...); err != nil {
		return err
	}
	return nil
}

func (c *LibP2pCore) getNodeOptions() []libp2p.Option {
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		log.Fatal(err)
	}
	return []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/tcp/%d", c.infos.ip, c.infos.Port())),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
		libp2p.DefaultTransports,
		libp2p.NATPortMap(),
	}
}

func (c *LibP2pCore) RequestPermanentPeerID() (string, error) {
	res, err := http.Get(c.tracker.HTTPURL() + "/ID")
	if err != nil {
		return "", err
	} else if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Http Endpoint returned wrong status: %d", res.StatusCode)
	}
	byteID, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(byteID), nil
}

func (c *LibP2pCore) DecodePermanentPeerID(requestedPeerId string) error {
	// The following code extracts target's the peer ID from the
	// given multiaddress
	p2paddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d/p2p/%s", c.tracker.IP(), c.tracker.Port(), requestedPeerId))
	if err != nil {
		return nil
	}
	//pid, err := p2paddr.ValueForProtocol(ma.P_IPFS)
	pid, err := p2paddr.ValueForProtocol(ma.P_P2P)
	if err != nil {
		return nil
	}
	// Decapsulate the /p2p/<peerID>
	// /ip4/<a.b.c.d>/p2p/<peer> becomes /ip4/<a.b.c.d>
	c.peerid, err = peer.Decode(pid)
	if err != nil {
		return nil
	}
	// Decapsulate the /p2p/<peerID> part from the target
	// /ip4/<a.b.c.d>/p2p/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/p2p/%s", pid))
	targetAddr := p2paddr.Decapsulate(targetPeerAddr)

	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	c.Node.Peerstore().AddAddr(c.peerid, targetAddr, peerstore.PermanentAddrTTL)
	return nil
}

func (c *LibP2pCore) Launch() error {
	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerid'.
	s, err := c.Node.NewStream(c.ctx, c.peerid, protocol.FileTransferProtocol)
	if err != nil {
		glog.Error(err)
		return err
	}
	c.Receive(s)
	//writeRandomStrings(bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)))
	return nil
}

func (c *LibP2pCore) SetStreamHandlers() error {
	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerid'.
	s, err := c.Node.NewStream(c.ctx, c.peerid, protocol.FileTransferProtocol)
	if err != nil {
		glog.Error(err)
		return err
	}
	c.Receive(s)
	//writeRandomStrings(bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)))
	return nil
}

func (c *LibP2pCore) Receive(s inet.Stream) {
	dec := json.NewDecoder(s)
	for {
		var v protocol.Datagram
		err := dec.Decode(&v)
		if err != nil {
			glog.Error(err)
			continue
		}
		log.Println("reading new datagrams")
		c.HandleDatagram(&v)
	}
}

func (c *LibP2pCore) HandleDatagram(d *protocol.Datagram) error {
	if len(d.Msgs) == 0 {
		return errors.New("Datagram receiver has no message")
	}
	for _, msg := range d.Msgs {
		log.Println(msg)
	}
	return nil
}

func (c *LibP2pCore) SendDatagram(d *protocol.Datagram) error {
	if len(d.Msgs) == 0 {
		return errors.New("Datagram receiver has no message")
	}
	for _, msg := range d.Msgs {
		glog.Info(msg)
	}
	return nil
}

func writeRandomStrings(writer *bufio.ReadWriter) error {
	for i := 0; i != 10; i += 1 {
		_, err := writer.WriteString(fmt.Sprintf("%s:%d\n", "coucou", i))
		if err != nil {
			return err
		}
		if err = writer.Flush(); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("messages Written")
	for i := 0; i != 10; i += 1 {
		str, err := writer.ReadString('\n')
		if err != nil {
			return err
		}
		log.Printf("read: %s\n", str)
	}

	return nil
}

func writeData(rw *bufio.Writer) error {
	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			return err
		}
		_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
		if err != nil {
			return err
		}
		err = rw.Flush()
		if err != nil {
			return err
		}
	}
}

// MakeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress.
func MakeBasicHost(ctx context.Context, opts []libp2p.Option) (host.Host, error) {
	basicHost, err := libp2p.New(ctx, opts...)
	if err != nil {
		return nil, err
	}
	fmt.Println(basicHost.ID().String())

	// Build host multiaddress
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/p2p/%s", basicHost.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	addr := basicHost.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)
	log.Printf("I am %s\n", fullAddr)
	return basicHost, nil
}

// SendRequest sends a request for the chunk range to the remote peer on the swarm
func SendRequest(start storage.ChunkID, end storage.ChunkID, remote protocol.PeerID) error {
	log.Println("SendReq Chunk %v-%v, to %v, on %v", start, end, remote)
	h := protocol.RequestChunks{Start: start, End: end}
	m := protocol.Msg{Op: protocol.Request, Data: h}
	//d := Datagram{Msgs: []Msg{m}}
	_ = protocol.Datagram{Msgs: []protocol.Msg{m}}
	//return p.sendDatagram(d, ours)
	return nil
}
