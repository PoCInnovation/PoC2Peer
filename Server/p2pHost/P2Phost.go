package p2pHost

import (
	"context"
	"fmt"
	libcore "github.com/PoCInnovation/PoPoc2PeerLibrary/core"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
)

func NewP2PHost(ip, prot string, listenPort int) (host.Host, error) {
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		return nil, err
	}
	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/%s/%d", ip, prot, listenPort)),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
		libp2p.DefaultTransports,
		libp2p.NATPortMap(),
	}
	return libcore.MakeBasicHost(context.Background(), opts)
}
