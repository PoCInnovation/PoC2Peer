package gomobile

import (
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/gomobile"
	"io/ioutil"
	"log"
	"net/http"
)

var Global = gomobile.SoundBuffer("SALUT JE SUIS UN ARRAY DE BYTES")
var Lib core.LibP2pCore

func ReadBuffer() []byte {
	return Global.Read()
}

const httpEndpoint = "http://192.168.0.31:5001/ID"

func GetID() string {
	res, err := http.Get(httpEndpoint)
	if err != nil {
		return err.Error()
	} else if res.StatusCode != http.StatusOK {
		return fmt.Sprintf("Http Endpoint returned wrong status: %d\n", res.StatusCode)
	}
	byteID, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(byteID)
}

func LaunchP2P() int {
	lib, err := core.NewLibP2p("0.0.0.0")
	if err != nil {
		log.Println(err)
		return 84
	}
	err = lib.Launch()
	if err != nil {
		log.Println(err)
		return 84
	}
	return 0
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
