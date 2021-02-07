package core

import "C"
import (
	"bufio"
	"context"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/data"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/network"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
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
	Node   host.Host
	peerid peer.ID
	peerip string
	ctx    context.Context
}

func NewLibP2p(ip string) (*LibP2pCore, error) {
	libCore := new(LibP2pCore)

	//Request Peer Id from Http Endpoint
	requestedPeerId, err := libCore.RequestPermanentPeerID()
	if err != nil {
		return nil, err
	}

	// One peerId is acquired from Permanent peer, intialize p2p node.
	err = libCore.InitNode()
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
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", 4000)),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
		libp2p.NATPortMap(),
	}

}

func (c *LibP2pCore) RequestPermanentPeerID() (string, error) {
	res, err := http.Get(httpEndpoint)
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
	p2paddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d/p2p/%s", 5000, requestedPeerId))
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
	s, err := c.Node.NewStream(c.ctx, c.peerid, "/echo/1.0.0")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(s)
	writeRandomStrings(writer)
	return nil
}

func writeRandomStrings(writer *bufio.Writer) error {
	for i := 0; i != 10; i += 1 {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", "coucou"))
		if err != nil {
			return err
		}
		if err = writer.Flush(); err != nil {
			return err

		}
		time.Sleep(100 * time.Millisecond)
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

// makeBasicHost creates a LibP2P host with a random peer ID listening on the
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
func SendRequest(start data.ChunkID, end data.ChunkID, remote network.PeerID) error {
	log.Println("SendReq Chunk %v-%v, to %v, on %v", start, end, remote)
	h := network.RequestMsg{Start: start, End: end}
	m := network.Msg{Op: network.Request, Data: h}
	//d := Datagram{Msgs: []Msg{m}}
	_ = network.Datagram{Msgs: []network.Msg{m}}
	//return p.sendDatagram(d, ours)
	return nil
}
