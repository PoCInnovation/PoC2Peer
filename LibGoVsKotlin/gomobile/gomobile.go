package gomobile

import (
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/kotlinHandler"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
)

const httpEndpoint = "http://192.168.0.31:5001/ID"

func LaunchP2P(localIP, ip string, port int) error {
	tracker := p2pnetwork.NewHttpTracker("192.168.0.31", 5000, "192.168.0.31", 5001, false)
	err := kotlinHandler.InitP2PLibrary(p2pnetwork.NewNetworkInfos(localIP, port), []p2pnetwork.Tracker{tracker})
	return err
}

func Read(ID string) []byte {
	return kotlinHandler.RequestFile(ID)
}
