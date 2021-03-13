package core

import "C"
import (
	"errors"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/protocol"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"os"
	"sync"

	//"github.com/libp2p/go-libp2p-core/peer"
	//ma "github.com/multiformats/go-multiaddr"
	"io/ioutil"
	"log"
	//"net/http"
	"time"
)

type LibP2pCore struct {
	network      p2pnetwork.Network
	infos        p2pnetwork.NetworkInfos
	trackers     []p2pnetwork.Tracker
	LocalStorage storage.LocalStorage
	PeerStorage  storage.PeerStorage
}

// NewLibP2P creates a LibP2P host with a random peer ID listening on the
// given p2pnetwork.NetworkInfos.
func NewLibP2P(infos p2pnetwork.NetworkInfos, prot string) (core *LibP2pCore, err error) {
	network, err := p2pnetwork.NewLibp2pNetwork(infos, prot)
	if err != nil {
		return nil, err
	}
	localStorage := storage.NewP2PStorage()
	peerStorage := storage.NewP2PRemoteStorage()
	return &LibP2pCore{
		network:      network,
		infos:        infos,
		LocalStorage: localStorage,
		PeerStorage:  peerStorage,
	}, nil
}

// NewP2PPeer: creates a LibP2P host with a random peer ID listening on the
// given multiaddress.
// TODO: remove once Tracker functionnal
func NewP2PPeer(trackers []p2pnetwork.Tracker, infos p2pnetwork.NetworkInfos, prot string) (*LibP2pCore, error) {
	libCore, err := NewLibP2P(infos, prot)
	if err != nil {
		return nil, err
	}

	libCore.trackers = trackers
	//libCore.PeerStorage = storage.NewP2PRemoteStorage()

	//Request Peer Id from Http Endpoint
	//err = libCore.getPeerList()
	//if err != nil {
	//	return nil, err
	//}
	libCore.SetDefaultStreamHandlers()
	return libCore, nil
}

// MakeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress.
func NewP2PPermanentPeer(trackers []p2pnetwork.Tracker, infos p2pnetwork.NetworkInfos, prot string) (*LibP2pCore, error) {
	libCore, err := NewLibP2P(infos, prot)
	if err != nil {
		return nil, err
	}
	libCore.trackers = trackers
	libCore.SetDefaultStreamHandlers()
	fmt.Println(libCore.ID())
	return libCore, nil
}

// ID: return a string representin the ID of this Peer
func (c *LibP2pCore) ID() string {
	return c.network.ID().String()
}

func (c *LibP2pCore) getPeerList() []p2pnetwork.PeerInfos {
	log.Println("Requesting Peers from trackers...")
	var peers []p2pnetwork.PeerInfos
	pid := c.ID()
	wg := sync.WaitGroup{}
	for _, tracker := range c.trackers {
		wg.Add(1)
		go func() {
			err := tracker.Ping()
			if err != nil {
				log.Println(fmt.Errorf("Ping tracker {%s} failed: %v", tracker.URL(), err))
				wg.Done()
				return
			}
			err = tracker.AddPeer(pid, c.infos.URL())
			if err != nil {
				log.Println(fmt.Errorf("AddRemotePeer for tracker {%s} failed: %v", tracker.URL(), err))
			}
			newPeers, err := tracker.Peers()
			if err != nil {
				wg.Done()
				log.Println(fmt.Errorf("Peers for tracker {%s} failed: %v", tracker.URL(), err))
				return
			}
			peers = append(peers, newPeers...)
			wg.Done()
		}()
	}
	wg.Wait()
	//log.Println("Peers Requested from tracker !")
	log.Println("Peers Requested from tracker !", peers)
	return removeDuplicates(peers, c.network.ID())
}

func removeDuplicates(s []p2pnetwork.PeerInfos, self p2pnetwork.PeerID) []p2pnetwork.PeerInfos {
	// Use of empty struct to optimize memory instead of boolean
	// https://dave.cheney.net/2014/03/25/the-empty-struct
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v.ID]; ok || v.ID == self.String() {
			log.Println("Removed: ", v.ID)
			continue
		}
		seen[v.ID] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

func (c LibP2pCore) RemovePeerFromTrackers() error {
	log.Println("Removing peer from trackers...")
	pid := c.ID()
	wg := sync.WaitGroup{}
	for _, tracker := range c.trackers {
		go func() {
			wg.Add(1)
			err := tracker.Ping()
			if err != nil {
				log.Println(fmt.Errorf("Ping tracker {%s} failed: %v", tracker.URL(), err))
				wg.Done()
				return
			}
			err = tracker.RemovePeer(pid)
			if err != nil {
				log.Println(fmt.Errorf("AddRemotePeer for tracker {%s} failed: %v", tracker.URL(), err))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("Peer removed from trackers !")
	return nil
}

func (c *LibP2pCore) Close() error {
	if err := c.RemovePeerFromTrackers(); err != nil {
		log.Println(err)
	}
	return c.network.Close()
}

//func (c *LibP2pCore) getPeerList() error {
//	res, err := http.Get(c.trackers.HTTPURL() + "/ID")
//	if err != nil {
//		return err
//	} else if res.StatusCode != http.StatusOK {
//		return fmt.Errorf("Http Endpoint returned wrong status: %d", res.StatusCode)
//	}
//	byteID, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		return err
//	}
//	return c.AddRemotePeer(string(byteID))
//}
//
//func (c *LibP2pCore) AddRemotePeer(requestedPeerId string) error {
//	// The following extracts target's the peer ID from the given multiaddress
//	p2paddr, err := ma.NewMultiaddr(
//		fmt.Sprintf("/ip4/%s/tcp/%d/p2p/%s", c.trackers.IP(), c.trackers.Port(), requestedPeerId),
//	)
//	if err != nil {
//		return nil
//	}
//	pid, err := p2paddr.ValueForProtocol(ma.P_P2P)
//	if err != nil {
//		return nil
//	}
//	peerid, err := peer.Decode(pid)
//	if err != nil {
//		return nil
//	}
//	log.Println("pid : ", pid)
//	log.Println("peer: ", peerid)
//	// Decapsulate the /p2p/<peerID> part from the target
//	// /ip4/<a.b.c.d>/p2p/<peer> becomes /ip4/<a.b.c.d>
//	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/p2p/%s", pid))
//	targetAddr := p2paddr.Decapsulate(targetPeerAddr)
//
//	// We have a peer ID and a targetAddr so we add it to the peerstore
//	// so LibP2P knows how to contact it
//	c.network.AddAddrs(peerid, []ma.Multiaddr{targetAddr})
//	return nil
//}

// TODO: Add Interface for RemotePeer (GetMultiAddr() ?)
func (c *LibP2pCore) AddRemotePeer(remotePeer p2pnetwork.PeerInfos) error {
	// The following extracts target's the peer ID from the given multiaddress
	p2paddr, err := ma.NewMultiaddr(
		fmt.Sprintf("/ip4/%s/%s/%d/p2p/%s", remotePeer.IP, remotePeer.Transport, remotePeer.Port, remotePeer.ID),
	)
	if err != nil {
		return err
	}
	pid, err := p2paddr.ValueForProtocol(ma.P_P2P)
	if err != nil {
		return err
	}
	peerid, err := peer.Decode(pid)
	if err != nil {
		return err
	}
	// Decapsulate the /p2p/<peerID> part from the target
	// /ip4/<a.b.c.d>/p2p/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/p2p/%s", pid))
	targetAddr := p2paddr.Decapsulate(targetPeerAddr)

	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	c.network.AddAddrs(peerid, []ma.Multiaddr{targetAddr})
	return nil
}

func (c *LibP2pCore) Launch(file string) error {
	log.Println("Launching peer: ", c.ID())
	lst := c.getPeerList()
	log.Println(lst)
	for _, peer := range lst {
		if err := c.AddRemotePeer(peer); err != nil {
			log.Fatal(err)
		}
	}
	//p, err := c.network.FirstPeer()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//c.SetDefaultStreamHandlers()
	//c.network.SetDatagramHandler(c.HandleDatagram)

	//s, err := c.network.Connect(p)
	////_, err = c.network.Connect(p)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//reqFile := storage.FileHashTmp(1)
	//h := protocol.RequestChunks{
	//	P2PFile: reqFile,
	//	IDs:  []storage.ChunkID{0, 1, 2, 3},
	//}
	//// log.Printf("Sending DataExchange : %v", h)
	//m := protocol.Msg{Op: protocol.Request, Data: h}
	//d := protocol.NewDataGram(m)
	//if err = c.network.SendDatagram(d, s.Stream.Conn().RemotePeer()); err != nil {
	//	return err
	//}
	//time.Sleep(time.Second * 5)
	//data, err := c.LocalStorage.GetFileData(reqFile)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(string(data))
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Can't read file %s: %v\n", file, err)
		return err
	}
	hash := storage.NewHashFromFile(content)

	fmt.Printf("%x\n", hash)
	data, err := c.RequestFile(hash)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("test_file.mp3", data, os.ModePerm)
	//log.Println(string(data))

	//c.Receive(s)

	//// Start a stream with the destination.
	//// Multiaddress of the destination peer is fetched from the peerstore using 'peerid'.
	//s, err := c.network.Host.NewStream(context.Background(), p.(peer.ID), protocol.FileTransferProtocol)
	//if err != nil {
	//	glog.Error(err)
	//	return err
	//}
	//writeRandomStrings(bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)))
	return nil
}

// SetDefaultStreamHandlers initailise p2p lib with defaults function to send and handle datagrams
func (c *LibP2pCore) SetDefaultStreamHandlers() error {
	c.network.SetDatagramHandler(c.HandleDatagram)
	//c.network.SetDatagramSender(c.HandleDatagram)
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
			resp, err := msg.HandleHave(pid, c.LocalStorage, c.PeerStorage)
			if err != nil {
				return err
			}
			if resp != nil {
				return c.network.SendDatagram(resp, pid)
			}
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
			return c.network.SendDatagram(chunks, pid)
		}
		//log.Println(msg)
	}
	return nil
}

// TODO: TO MODIFY
func (c *LibP2pCore) RequestFile(fileID storage.FileHash) ([]byte, error) {
	datas, err := c.LocalStorage.GetFileData(fileID)
	if err == storage.FILENOTFOUND {
		log.Println("Requesting files to peers")
		c.network.RequestFileToPeers(fileID, c.PeerStorage)
		time.Sleep(time.Second * 2)
		datas, err = c.LocalStorage.GetFileData(fileID)
	}
	if err != nil {
		return nil, err
	}
	return datas, err
}
