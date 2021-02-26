package p2pnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/protocol"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	libp2ppeer "github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
	"io"
	"log"
)

// Network handles interactions with the underlying protocol
type Network interface {

	//// SendDatagram sends a datagram to the remote peer
	//SendDatagram(d Datagram, remote PeerID) error
	//
	// Connect connects to the remote peer and creates any io resources necessary for the connection
	//Connect(remote PeerID) error
	Connect(remote PeerID) (network.Stream, error)

	//// Disconnect disconnects from the remote peer and destroys any io resources created for the connection
	//Disconnect(remote PeerID) error
	//
	// ID returns the ID of this peer
	ID() PeerID

	// TODO: Remove when swarm functionnal
	// ID returns the ID of this peer
	FirstPeer() (PeerID, error)

	// SetDatagramHandler sets the function that will be called on receipt of a new datagram
	// f gets called every time a new Datagram is received.
	SetDatagramHandler(func(*protocol.Datagram, PeerID) error)

	// AddAddrs adds multiaddresses for the remote peer to this peer's store
	AddAddrs(id PeerID, addrs []ma.Multiaddr)

	// Addrs returns multiaddresses for this peer
	Addrs() []ma.Multiaddr
}

// PeerID identifies a peer
type PeerID interface {
	String() string
}

type P2PNetwork struct {
	// all of this peer's streams, indexed by a global? peer.ID
	streams map[PeerID]*WrappedStream
	Host    host.Host
}

func InitPeer(infos NetworkInfos, prot string, ctx context.Context) (host.Host, error) {
	opts := basicPeerOptions(infos, prot)
	node, err := libp2p.New(ctx, opts...)
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
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/%s/%d", infos.ip, prot, infos.Port())),
		libp2p.Identity(priv),
		libp2p.DefaultTransports,
		libp2p.NATPortMap(),
	}
}

func NewLibp2pNetwork(infos NetworkInfos, prot string) (*P2PNetwork, error) {
	streams := make(map[PeerID]*WrappedStream)
	h, err := InitPeer(infos, prot, context.Background())
	if err != nil {
		return nil, err
	}
	return &P2PNetwork{streams: streams, Host: h}, nil
}

// Connect creates a stream from p to the peer at id and sets a stream handler
func (n *P2PNetwork) Connect(id PeerID) (*WrappedStream, error) {
	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerid'.
	stream, err := n.Host.NewStream(context.Background(), id.(libp2ppeer.ID), protocol.FileTransferProtocol)
	if err != nil {
		return nil, err
	}
	ws := WrapStream(stream)
	n.streams[id] = ws
	//n.SetDatagramHandler(protocol.HandleDatagram)
	return ws, nil
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
				log.Fatal(err)
			}
			if err = handler(d, remote); err != nil {
				log.Fatal("Datagram Handler err", err)
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
	log.Printf("%v decoded datagram %v\n", n.ID(), d)
	return &d, nil
}

func (n *P2PNetwork) ID() PeerID {
	return n.Host.ID()
}

func (n *P2PNetwork) FirstPeer() (PeerID, error) {
	addrs := n.Host.Peerstore().Peers()
	if len(addrs) < 1 {
		return libp2ppeer.ID(""), errors.New("No Peers yet")
	}
	return addrs[1], nil
}

func (n *P2PNetwork) AddAddrs(remote PeerID, addrs []ma.Multiaddr) {
	n.Host.Peerstore().AddAddrs(remote.(libp2ppeer.ID), addrs, ps.PermanentAddrTTL)
}

func (n *P2PNetwork) Addrs() []ma.Multiaddr {
	return n.Host.Addrs()
}
