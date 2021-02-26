package gomobile

import (
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
	"log"
)

var Lib *core.LibP2pCore

type SoundBuffer []byte

func InitP2PLibrary(infos p2pnetwork.NetworkInfos, trackers []p2pnetwork.HttpTracker) (err error) {
	for _, t := range trackers {
		tracker := p2pnetwork.NewHttpTracker(t.IP(), t.Port(), t.HTTPIP(), t.HTTPPort(), false)
		Lib, err = core.NewP2PPeer(tracker, infos, "tcp")
		if err == nil {
			break
		}
		log.Println(err)
	}
	if err != nil {
		log.Fatal(err)
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
//	return "string -> Maxime, je suis le meilleur prends moi chez Skillz STP\n"
//}
//
//func CallByteArray() []byte {
//	return []byte("byte array -> Maxime, je suis le meilleur prends moi chez Skillz STP\n")
//}
//
//func CallIntArray() []uint {
//	return []uint{1, 2, 3}
//}
//
//func CallInterfaceInt() interface{} {
//	return []int{1, 2, 3}
//}
//
//func CallMap() map[string]int {
//	return map[string]int{"lol": 1, "mdr": 2, "xd": 3}
//}
//
//func CallInterfaceMap() interface{} {
//	return map[string]int{"lol": 1, "mdr": 2, "xd": 3}
//}
//
//func CallPackageVariableStringified() string {
//	return fmt.Sprintf("%v\n", packageVar)
//}
