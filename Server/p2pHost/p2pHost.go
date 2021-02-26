package p2pHost

import (
	p2pcore "github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
)

func NewP2PHost(ip, prot string, listenPort int) (*p2pcore.LibP2pCore, error) {
	infos := p2pnetwork.NewNetworkInfos(ip, listenPort)
	return p2pcore.NewP2PPermanentPeer(nil, infos, prot)
}
