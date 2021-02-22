package p2pnetwork

import (
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// Network handles interactions with the underlying protocol
type Network interface {

	//// SendDatagram sends a datagram to the remote peer
	//SendDatagram(d Datagram, remote PeerID) error
	//
	//// Connect connects to the remote peer and creates any io resources necessary for the connection
	//Connect(remote PeerID) error
	//
	//// Disconnect disconnects from the remote peer and destroys any io resources created for the connection
	//Disconnect(remote PeerID) error
	//
	//// ID returns the ID of this peer
	//ID() PeerID
	//
	//// SetDatagramHandler sets the function that will be called on receipt of a new datagram
	//// f gets called every time a new Datagram is received.
	//SetDatagramHandler(f func(*Datagram, PeerID) error)
	//
	//// AddAddrs adds multiaddresses for the remote peer to this peer's store
	//AddAddrs(id PeerID, addrs []ma.Multiaddr)
	//
	//// Addrs returns multiaddresses for this peer
	//Addrs() []ma.Multiaddr
}

type libp2pNetwork struct {
	// all of this peer's streams, indexed by a global? peer.ID
	streams map[peer.ID]*WrappedStream
	h       host.Host
}

func newLibp2pNetwork(port int) (*libp2pNetwork, error) {
	streams := make(map[peer.ID]*WrappedStream)
	h, err := newBasicHost(port)
	if err != nil {
		return nil, err
	}
	return &libp2pNetwork{streams: streams, h: h}, nil
}

func (n *libp2pNetwork) ID() PeerID {
	return n.h.ID()
}

func (n *libp2pNetwork) AddAddrs(remote PeerID, addrs []ma.Multiaddr) {
	n.h.Peerstore().AddAddrs(remote.(p2ppeer.ID), addrs, ps.PermanentAddrTTL)
}

func (n *libp2pNetwork) Addrs() []ma.Multiaddr {
	return n.h.Addrs()
}
