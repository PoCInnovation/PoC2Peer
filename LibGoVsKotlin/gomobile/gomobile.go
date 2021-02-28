package gomobile

import (
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/kotlinHandler"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
)

const httpEndpoint = "http://192.168.0.31:5001/ID"

//func LaunchP2P(localIP, ip string, port int) error {
//	tracker := p2pnetwork.NewHttpTracker("192.168.0.31", 5000, "192.168.0.31", 5001, false)
//	lib, err := core.NewP2PPeer(tracker, p2pnetwork.NewNetworkInfos(localIP, 4000), "tcp")
//	if err != nil {
//		return err
//	}
//	err = lib.Launch()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//

func LaunchP2P(localIP, ip string, port int) error {
	tracker := p2pnetwork.NewHttpTracker("192.168.0.31", 5000, "192.168.0.31", 5001, false)
	err := kotlinHandler.InitP2PLibrary(p2pnetwork.NewNetworkInfos(localIP, port), []p2pnetwork.Tracker{tracker})
	return err
}

func Read(ID int) []byte {
	return kotlinHandler.RequestFile(ID)
}
