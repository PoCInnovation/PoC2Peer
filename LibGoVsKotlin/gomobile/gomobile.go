package gomobile

import (
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/kotlinHandler"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
)

const (
	httpEndpoint = "http://192.168.0.31:5001/ID"
	TrackerIP    = "163.172.143.105"
	TrackerPort  = 3000
)

func LaunchP2P(localIP, ip string, port int) error {
	tracker := p2pnetwork.NewHttpTracker(TrackerIP, TrackerPort, false)
	err := kotlinHandler.InitP2PLibrary(p2pnetwork.NewNetworkInfos(localIP, port), []p2pnetwork.Tracker{tracker})
	return err
}

func Open(ID string) error {
	return kotlinHandler.Open(ID)
}

func Read(buf []byte, offset int, readLength int, ID string) int {
	return kotlinHandler.Read(buf, offset, readLength, ID)
}

//func Read(ID string) []byte {
//	return kotlinHandler.RequestFile(ID)
//}
