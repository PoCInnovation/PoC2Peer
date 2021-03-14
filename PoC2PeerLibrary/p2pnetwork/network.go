package p2pnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/protocol"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	libp2ppeer "github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-core/peerstore"
	libp2pprot "github.com/libp2p/go-libp2p-core/protocol"
	ma "github.com/multiformats/go-multiaddr"
	"io"
	"log"
	"time"
)

// Network handles interactions with the underlying protocol
type Network interface {

	// SendDatagram sends a datagram to the remote peer
	SendDatagram(d *protocol.Datagram, remote PeerID) error

	// Connect connects to the remote peer and creates any io resources necessary for the connection
	//Connect(remote PeerID) error
	//Connect(remote PeerID) (*WrappedStream, error)
	Connect(remote PeerID, protocol string) (*WrappedStream, error)

	// Disconnect disconnects from the remote peer and destroys any io resources created for the connection
	Disconnect(remote PeerID) error

	// ID returns the ID of this peer
	ID() PeerID

	//// TODO: Remove when swarm functionnal
	//// ID returns the ID of this peer
	//FirstPeer() (PeerID, error)

	// SetDatagramHandler sets the function that will be called on receipt of a new datagram
	// f gets called every time a new Datagram is received.
	SetDatagramHandler(func(*protocol.Datagram, PeerID) error)

	// AddAddrs adds multiaddresses for the remote peer to this peer's store
	AddAddrs(id PeerID, addrs []ma.Multiaddr)

	// Addrs returns multiaddresses for this peer
	Addrs() []ma.Multiaddr

	// Close close the network
	Close() error

	// TODO: move in protocol ?
	// Peers return all connected peers
	Peers() []PeerID
	RequestFileToPeers(file storage.FileHash, remoteStorage storage.PeerStorage) (int, error)
}

// PeerID identifies a peer
type PeerID interface {
	String() string
}

type P2PPeerID string

func (id P2PPeerID) String() string {
	return string(id)
}

type P2PNetwork struct {
	// all of this peer's streams, indexed by a global? peer.ID
	streams map[PeerID]*WrappedStream
	Host    host.Host
	ctx     context.Context
}

func InitPeer(infos NetworkInfos, prot string, ctx context.Context) (host.Host, error) {
	opts := basicPeerOptions(infos, prot)
	node, err := libp2p.New(ctx, opts...)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return nil, err
	}
	return node, nil
}

func basicPeerOptions(infos NetworkInfos, prot string) []libp2p.Option {
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		log.Fatal(err)
	}
	return []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/%s/%d", prot, infos.Port())),
		libp2p.Identity(priv),
		libp2p.DefaultTransports,
		libp2p.NATPortMap(),
	}
}

func NewLibp2pNetwork(infos NetworkInfos, prot string) (*P2PNetwork, error) {
	//func NewLibp2pNetwork(infos NetworkInfos, prot string) (Network, error) {
	network := &P2PNetwork{
		streams: make(map[PeerID]*WrappedStream),
		ctx:     context.Background(),
	}
	h, err := InitPeer(infos, prot, network.ctx)
	if err != nil {
		return nil, err
	}
	network.Host = h
	return network, nil
}

// Connect creates a stream from p to the peer at id and sets a stream handler
//func (n *P2PNetwork) Connect(id PeerID) (*WrappedStream, error) {
func (n *P2PNetwork) Connect(id PeerID, protocol string) (*WrappedStream, error) {
	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerid'.
	stream, err := n.Host.NewStream(n.ctx, id.(libp2ppeer.ID), libp2pprot.ID(protocol))
	if err != nil {
		return nil, err
	}
	ws := WrapStream(stream)
	n.streams[id] = ws
	return ws, nil
}

func (n *P2PNetwork) Disconnect(remote PeerID) error {
	s, ok := n.streams[remote]
	if !ok {
		return fmt.Errorf("Peer is not connected, ID: %v", remote)
	}
	err := s.Close()
	delete(n.streams, remote)
	return err
}

func (n *P2PNetwork) SetDatagramHandler(handler func(*protocol.Datagram, PeerID) error) {
	n.Host.SetStreamHandler(protocol.FileTransferProtocol, func(s network.Stream) {
		remote := PeerID(s.Conn().RemotePeer())
		log.Printf("%s received a stream from %s\n", n.ID(), remote)
		defer s.Close()
		ws := WrapStream(s)
		for {
			d, err := n.receiveDatagram(ws)
			if err == io.EOF {
				log.Printf("%v received EOF\n", n.ID())
				break
			}
			if err != nil {
				log.Println(err)
				break
			}
			if err = handler(d, remote); err != nil {
				log.Fatal("Datagram Handler err", err)
				break
			}
		}
		log.Printf("%v handled stream\n", n.ID())
	})
}

// receiveDatagram reads and decodes a datagram from the stream
func (n *P2PNetwork) receiveDatagram(ws *WrappedStream) (*protocol.Datagram, error) {
	log.Printf("%v received Datagram\n", n.ID())
	if ws == nil {
		return nil, fmt.Errorf("%v receiveDatagram on nil *WrappedStream", n.ID())
	}
	var d protocol.Datagram
	if err := ws.Dec.Decode(&d); err != nil {
		return nil, err
	}
	log.Printf("%v decoded datagram\n", n.ID())
	return &d, nil
}

func (n *P2PNetwork) SendDatagram(d *protocol.Datagram, pid PeerID) (err error) {
	// TODO: Check Protocol State ?
	s, ok := n.streams[pid]
	if !ok {
		s, err = n.Connect(pid, protocol.FileTransferProtocol)
	}
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

func (n *P2PNetwork) ID() PeerID {
	return n.Host.ID()
}

func (n *P2PNetwork) Close() error {
	for peer, stream := range n.streams {
		log.Printf("Closing Stream for: %v", peer)
		stream.Close()
	}
	return n.Host.Close()
}

func (n *P2PNetwork) FirstPeer() (PeerID, error) {
	for _, addr := range n.Host.Peerstore().Peers() {
		if addr != n.Host.ID() {
			fmt.Println(addr)
			return addr, nil
		}
	}
	return nil, errors.New("No Peers yet")
}

func (n *P2PNetwork) AddAddrs(remote PeerID, addrs []ma.Multiaddr) {
	n.Host.Peerstore().AddAddrs(remote.(libp2ppeer.ID), addrs, ps.PermanentAddrTTL)
}

func (n *P2PNetwork) Addrs() []ma.Multiaddr {
	return n.Host.Addrs()
}

func (n *P2PNetwork) Peers() []PeerID {
	peers := n.Host.Peerstore().Peers()
	ret := make([]PeerID, len(peers))
	for i, id := range peers {
		ret[i] = id
	}
	return ret
}

func (n *P2PNetwork) RequestFileToPeers(file storage.FileHash, remoteStorage storage.PeerStorage) (int, error) {
	d1 := protocol.NewDataGram(protocol.Msg{Op: protocol.Have, Data: protocol.HaveMsg{File: file, Type: protocol.HaveRequest}})
	for _, peer := range n.Peers() {
		if peer.String() == n.ID().String() {
			continue
		}
		log.Println("Sending HAVE REQUEST: ", peer)
		err := n.SendDatagram(d1, peer)
		if err != nil {
			log.Fatal(err)
		}
	}
	time.Sleep(time.Second * 2)
	ls, err := remoteStorage.GetPeersFileChunks(file)
	if err != nil {
		return 0, err
	}
	var nbChunk int
	for _, chunks := range ls {
		if len(chunks) > nbChunk {
			nbChunk = len(chunks)
		}
	}
	for peer, chunks := range ls {
		d2 := protocol.NewDataGram(protocol.Msg{Op: protocol.Request, Data: protocol.RequestChunks{File: file, IDs: chunks}})
		log.Println("Requesting to peer: ", peer)
		err := n.SendDatagram(d2, peer)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nbChunk * storage.LocalStorageSize, nil
}
