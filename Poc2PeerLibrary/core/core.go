package core

import "C"
import (
	"errors"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/protocol"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type LibP2pCore struct {
	N            p2pnetwork.Network
	infos        p2pnetwork.NetworkInfos
	tracker      p2pnetwork.Tracker
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

// MakeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress.
func NewP2PPermanentPeer(tracker p2pnetwork.Tracker, infos p2pnetwork.NetworkInfos, prot string) (*LibP2pCore, error) {
	libCore, err := NewLibP2P(infos, prot)
	if err != nil {
		return nil, err
	}
	libCore.tracker = tracker
	libCore.SetStreamHandlers()
	fmt.Println(libCore.ID())
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
	//p, err := c.N.FirstPeer()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//c.SetStreamHandlers()
	//c.N.SetDatagramHandler(c.HandleDatagram)

	//s, err := c.N.Connect(p)
	////_, err = c.N.Connect(p)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//reqFile := storage.FileHashTmp(1)
	//h := protocol.RequestChunks{
	//	File: reqFile,
	//	IDs:  []storage.ChunkID{0, 1, 2, 3},
	//}
	//// log.Printf("Sending DataExchange : %v", h)
	//m := protocol.Msg{Op: protocol.Request, Data: h}
	//d := protocol.NewDataGram(m)
	//if err = c.N.SendDatagram(d, s.Stream.Conn().RemotePeer()); err != nil {
	//	return err
	//}
	//time.Sleep(time.Second * 5)
	//data, err := c.LocalStorage.GetFileDatas(reqFile)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(string(data))

	data, err := c.RequestFile(storage.FileHashTmp(1))
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
	//c.N.SetDatagramSender(c.HandleDatagram)
	return nil
}

func (c *LibP2pCore) HandleDatagram(d *protocol.Datagram, pid p2pnetwork.PeerID) error {
	if len(d.Msgs) == 0 {
		return errors.New("Datagram receiver has no message")
	}
	for _, msg := range d.Msgs {
		// TODO: in message.go
		//if err := msg.HandleMsg() ; err != nil {
		//	return err
		//}
		switch msg.Op {
		case protocol.Have:
			log.Println("handling HAVE message datagram")
			//err := msg.HandleDataExchange(c.LocalStorage)
			//if err != nil {
			//	return err
			//}
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
			return c.N.SendDatagram(chunks, pid)
		}
		//log.Println(msg)
	}
	return nil
}

// TODO: TO MODIFY
func (c *LibP2pCore) RequestFile(fileID storage.FileHashTmp) ([]byte, error) {
	datas, err := c.LocalStorage.GetFileDatas(fileID)
	if err == storage.FILENOTFOUND {
		log.Println("Requesting files to peers")
		c.N.RequestFileToPeers(fileID)
		time.Sleep(time.Second * 1)
		datas, err = c.LocalStorage.GetFileDatas(fileID)
	}
	if err != nil {
		return nil, err
	}
	return datas, err
}

//func (c *LibP2pCore) SendDatagram(d *protocol.Datagram, pid p2pnetwork.PeerID) error {
//	s, err := c.N.Connect(pid)
//	defer s.Close()
//	if err != nil {
//		return err
//	}
//	err = s.Enc.Encode(&d)
//	if err != nil {
//		return err
//	}
//	err = s.W.Flush()
//	if err != nil {
//		return err
//	}
//	return nil
//}
