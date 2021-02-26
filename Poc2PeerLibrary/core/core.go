package core

import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/protocol"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
	"github.com/golang/glog"
	inet "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type LibP2pCore struct {
	//N            p2pnetwork.Network
	N       *p2pnetwork.P2PNetwork
	infos   p2pnetwork.NetworkInfos
	tracker p2pnetwork.Tracker
	//LocalStorage *storage.P2PStorage
	LocalStorage storage.LocalStorage
}

// NewLibP2P creates a LibP2P host with a random peer ID listening on the
// given multiaddress.
func NewLibP2P(infos p2pnetwork.NetworkInfos, prot string) (core *LibP2pCore, err error) {
	network, err := p2pnetwork.NewLibp2pNetwork(infos, prot)
	if err != nil {
		return nil, err
	}
	localStorage := storage.NewP2PStorage()
	return &LibP2pCore{N: network, infos: infos, LocalStorage: localStorage}, nil
}

// MakeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress.
func NewP2PPeer(tracker p2pnetwork.Tracker, infos p2pnetwork.NetworkInfos, prot string) (*LibP2pCore, error) {
	libCore, err := NewLibP2P(infos, prot)
	if err != nil {
		return nil, err
	}

	libCore.tracker = tracker

	//Request Peer Id from Http Endpoint
	err = libCore.RequestPermanentPeerID()
	if err != nil {
		return nil, err
	}
	libCore.SetStreamHandlers()
	return libCore, nil
}

func (c *LibP2pCore) ID() string {
	return c.N.ID().String()
}
func (c *LibP2pCore) RequestPermanentPeerID() error {
	res, err := http.Get(c.tracker.HTTPURL() + "/ID")
	if err != nil {
		return err
	} else if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Http Endpoint returned wrong status: %d", res.StatusCode)
	}
	byteID, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return c.DecodePermanentPeerID(string(byteID))
}

func (c *LibP2pCore) DecodePermanentPeerID(requestedPeerId string) error {
	// The following extracts target's the peer ID from the given multiaddress
	p2paddr, err := ma.NewMultiaddr(
		fmt.Sprintf("/ip4/%s/tcp/%d/p2p/%s", c.tracker.IP(), c.tracker.Port(), requestedPeerId),
	)
	if err != nil {
		return nil
	}
	pid, err := p2paddr.ValueForProtocol(ma.P_P2P)
	if err != nil {
		return nil
	}
	peerid, err := peer.Decode(pid)
	if err != nil {
		return nil
	}
	// Decapsulate the /p2p/<peerID> part from the target
	// /ip4/<a.b.c.d>/p2p/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/p2p/%s", pid))
	targetAddr := p2paddr.Decapsulate(targetPeerAddr)

	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	c.N.AddAddrs(peerid, []ma.Multiaddr{targetAddr})
	return nil
}

func (c *LibP2pCore) Launch() error {
	p, err := c.N.FirstPeer()
	if err != nil {
		log.Fatal(err)
	}
	c.SetStreamHandlers()
	//c.N.SetDatagramHandler(c.HandleDatagram)

	s, err := c.N.Connect(p)
	//_, err = c.N.Connect(p)
	if err != nil {
		log.Fatal(err)
	}
	reqFile := storage.FileHashTmp(1)
	h := protocol.RequestChunks{
		File:  reqFile,
		Start: 0, End: 1,
	}
	// log.Printf("Sending DataExchange : %v", h)
	m := protocol.Msg{Op: protocol.Request, Data: h}
	d := protocol.NewDataGram(m)
	if err = c.SendDatagram(d, s.Stream.Conn().RemotePeer()); err != nil {
		return err
	}
	time.Sleep(time.Second * 25)
	data, err := c.LocalStorage.GetFileDatas(reqFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
	//c.Receive(s)

	//// Start a stream with the destination.
	//// Multiaddress of the destination peer is fetched from the peerstore using 'peerid'.
	//s, err := c.N.Host.NewStream(context.Background(), p.(peer.ID), protocol.FileTransferProtocol)
	//if err != nil {
	//	glog.Error(err)
	//	return err
	//}
	//writeRandomStrings(bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)))
	return nil
}

func (c *LibP2pCore) SetStreamHandlers() error {
	c.N.SetDatagramHandler(c.HandleDatagram)
	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerid'.
	//s, err := c.Node.NewStream(c.ctx, c.peerid, protocol.FileTransferProtocol)
	//if err != nil {
	//	glog.Error(err)
	//	return err
	//}
	//c.Receive(s)
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
			return
		}
		log.Println("reading new datagrams")
		//c.HandleDatagram(&v, s.Conn().ID())
	}
}

func (c *LibP2pCore) HandleDatagram(d *protocol.Datagram, pid p2pnetwork.PeerID) error {
	if len(d.Msgs) == 0 {
		return errors.New("Datagram receiver has no message")
	}
	for _, msg := range d.Msgs {
		switch msg.Op {
		case protocol.Data:
			log.Println("handling DATA message datagram")
			err := msg.HandleDataExchange(c.LocalStorage)
			if err != nil {
				return err
			}
		case protocol.Request:
			log.Printf("handling REQUEST message datagram\n")
			chunks, err := msg.HandleRequest(c.LocalStorage)
			if err != nil {
				log.Println("ERROR WHEN HANDLING REQUEST message datagram")
				return err
			}
			log.Println("finished processing REQUEST message datagram")
			return c.SendDatagram(chunks, pid)
		}
		//log.Println(msg)
	}
	return nil
}

func (c *LibP2pCore) SendDatagram(d *protocol.Datagram, pid p2pnetwork.PeerID) error {
	s, err := c.N.Connect(pid)
	defer s.Close()
	if err != nil {
		return err
	}
	err = s.Enc.Encode(&d)
	if err != nil {
		return err
	}
	err = s.W.Flush()
	if err != nil {
		return err
	}
	return nil
}

//func writeRandomStrings(writer *bufio.ReadWriter) error {
//	for i := 0; i != 10; i += 1 {
//		_, err := writer.WriteString(fmt.Sprintf("%s:%d\n", "coucou", i))
//		if err != nil {
//			return err
//		}
//		if err = writer.Flush(); err != nil {
//			return err
//		}
//		time.Sleep(100 * time.Millisecond)
//	}
//	fmt.Println("messages Written")
//	for i := 0; i != 10; i += 1 {
//		str, err := writer.ReadString('\n')
//		if err != nil {
//			return err
//		}
//		log.Printf("read: %s\n", str)
//	}
//
//	return nil
//}
//
//func writeData(rw *bufio.Writer) error {
//	stdReader := bufio.NewReader(os.Stdin)
//	for {
//		fmt.Print("> ")
//		sendData, err := stdReader.ReadString('\n')
//		if err != nil {
//			return err
//		}
//		_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
//		if err != nil {
//			return err
//		}
//		err = rw.Flush()
//		if err != nil {
//			return err
//		}
//	}
//}
//
//
//// SendRequest sends a request for the chunk range to the remote peer on the swarm
//func SendRequest(start storage.ChunkID, end storage.ChunkID, remote protocol.PeerID) error {
//	log.Println("SendReq Chunk %v-%v, to %v, on %v", start, end, remote)
//	h := protocol.RequestChunks{Start: start, End: end}
//	m := protocol.Msg{Op: protocol.Request, Data: h}
//	//d := Datagram{Msgs: []Msg{m}}
//	_ = protocol.Datagram{Msgs: []protocol.Msg{m}}
//	//return p.sendDatagram(d, ours)
//	return nil
//}
