package kotlinHandler

import (
	"encoding/hex"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
	"log"
)

var Lib *core.LibP2pCore

//type FileDatas struct {
//	Cursor int
//	Data   []byte
//}
//
//type Storage map[string]FileDatas
//
//type SoundBuffer struct {
//	Storage        Storage
//	CurrentTrackID string
//}
//
//var buffer = SoundBuffer{
//	make(Storage),
//	"",
//}

func InitP2PLibrary(infos p2pnetwork.NetworkInfos, trackers []p2pnetwork.Tracker) (err error) {
	Lib, err = core.NewP2PPeer(trackers, infos, "tcp")
	Lib.Launch()
	return err
}

func CloseP2PLibrary() {
	Lib.Close()
}

func Open(ID string) int {
	he, err := hex.DecodeString(ID)
	if err != nil {
		log.Println("decoding filehash failed")
		return -1
	}
	dataLength, err := Lib.InitRequestFile(storage.FileHash(he))
	if err != nil {
		log.Println(err)
		return 0
	}
	return dataLength
}

func Read(buf []byte, sourcePos, destPos, readLength int, ID string) ([]byte, error) {
	he, err := hex.DecodeString(ID)
	if err != nil {
		log.Println("decoding filehash failed")
		return nil, err
	}
	data, err := Lib.RequestFile(storage.FileHash(he))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var endOffset int
	if sourcePos+readLength > len(data) {
		endOffset = len(data)
	} else {
		endOffset = sourcePos + readLength
	}
	log.Printf("Reading from {%d} to {%d}\n", sourcePos, endOffset)
	log.Println(data[sourcePos:endOffset])

	return data[sourcePos:endOffset], nil

	//
	//var endOffset int
	//if sourcePos+readLength > len(data) {
	//	endOffset = len(data)
	//} else {
	//	endOffset = sourcePos + readLength
	//}
	//log.Printf("Reading from {%d} to {%d}\n", sourcePos, endOffset)
	//
	//var inBuf = bytes.NewBuffer(data[sourcePos:endOffset])
	//var outBuf = bytes.NewBuffer(buf[destPos:])
	//log.Println(inBuf.String())
	//l, err := io.Copy(outBuf, inBuf)
	//if err != nil {
	//	log.Println(err)
	//	return -1
	//}
	//return l

	//var endOffset int
	//if sourcePos+readLength > len(data) {
	//	endOffset = len(data)
	//} else {
	//	endOffset = sourcePos + readLength
	//}
	//copy(buf, data[sourcePos:endOffset])
	//return len(buf)

}

func Close(ID string) {
	he, err := hex.DecodeString(ID)
	if err != nil {
		log.Printf("decoding filehash failed")
		return
	}
	Lib.LocalStorage.DeleteData(storage.FileHash(he))
}
