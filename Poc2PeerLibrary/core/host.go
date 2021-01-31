package core

import (
	"context"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/data"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/network"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	ma "github.com/multiformats/go-multiaddr"
	"log"
)

type LibP2pCore struct {
	host.Host
}

// makeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress. It won't encrypt the connection if insecure is true.
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
