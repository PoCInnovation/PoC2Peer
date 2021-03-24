# P2PNetwork

## Package Description

## Interfaces
```go

// Network handles interactions with the underlying protocol
type Network interface {

	// SendDatagram sends a datagram to the remote peer
	SendDatagram(d *protocol.Datagram, remote PeerID) error

	// Connect connects to the remote peer and creates any io resources necessary for the connection
	Connect(remote PeerID, protocol string) (*WrappedStream, error)

	// Disconnect disconnects from the remote peer and destroys any io resources created for the connection
	Disconnect(remote PeerID) error

	// ID returns the ID of this peer
	ID() PeerID

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

```

```go
// PeerID identifies a peer
type PeerID interface {
	String() string
}
```

## Types
```go
type P2PPeerID string
```

```go
type P2PNetwork struct {
	// all of this peer's streams, indexed by a peer.ID
	streams map[PeerID]*WrappedStream
	Host    host.Host
	ctx     context.Context
}
```

## Misc
