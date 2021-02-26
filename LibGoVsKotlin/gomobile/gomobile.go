package gomobile

import (
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/gomobile"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
)

var Global = gomobile.SoundBuffer("SALUT JE SUIS UN ARRAY DE BYTES")
var Lib core.LibP2pCore

func ReadBuffer() []byte {
	return Global.Read()
}

const httpEndpoint = "http://192.168.0.31:5001/ID"

func LaunchP2P(localIP, ip string, port int) error {
	tracker := p2pnetwork.NewHttpTracker("192.168.0.31", 5000, "192.168.0.31", 5001, false)
	lib, err := core.NewP2PPeer(tracker, p2pnetwork.NewNetworkInfos(localIP, 4000), "tcp")
	if err != nil {
		return err
	}
	err = lib.Launch()
	if err != nil {
		return err
	}
	return nil
}

//// callback
//var jc JavaCallback
//
//type JavaCallback interface {
//	SendString(string)
//}
//
//type TestStruct struct {
//	Str string
//}
//
//var packageVar = TestStruct{Str: "InsidePackagevar"}
//
//func RegisterJavaCallback(c JavaCallback) {
//	jc = c
//}
//
//func TestCall() {
//	for i := 0; i < 100; i++ {
//		time.Sleep(100 * time.Millisecond)
//		jc.SendString(fmt.Sprintln("Counting... ", i))
//	}
//}
//
//func CallString() string {
//    return "string -> Maxime, je suis le meilleur prends moi chez Skillz STP\n"
//}
//
//func CallByteArray() []byte {
//    return []byte("byte array -> Maxime, je suis le meilleur prends moi chez Skillz STP\n")
//}
//
//func CallIntArray() []uint {
//    return []uint{1, 2, 3}
//}
//
//func CallInterfaceInt() interface{} {
//    return []int{1, 2, 3}
//}
//
//func CallMap() map[string]int {
//    return map[string]int{"lol": 1, "mdr":2, "xd": 3}
//}
//
//func CallInterfaceMap() interface{} {
//    return map[string]int{"lol": 1, "mdr":2, "xd": 3}
//}
//
//func CallPackageVariableStringified() string {
//    return fmt.Sprintf("%v\n", packageVar)
//}
//
