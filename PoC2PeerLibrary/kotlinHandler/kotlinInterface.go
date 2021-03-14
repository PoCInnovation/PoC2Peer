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

func Open(ID string) int {
	he, err := hex.DecodeString(ID)
	if err != nil {
		log.Printf("decoding filehash failed")
		return -1
	}
	dataLength, err := Lib.InitRequestFile(storage.FileHash(he))
	if err != nil {
		log.Println(err)
		return 0
	}
	return dataLength
}

func Read(buf []byte, sourcePos, destPos, readLength int, ID string) int {
	he, err := hex.DecodeString(ID)
	if err != nil {
		log.Printf("decoding filehash failed")
		return -1
	}
	data, err := Lib.RequestFile(storage.FileHash(he))
	if err != nil {
		return 0
	}
	var endOffset int
	if sourcePos+readLength > len(data) {
		endOffset = len(data)
	} else {
		endOffset = sourcePos + readLength
	}
	copy(buf, data[sourcePos:endOffset])
	return len(buf)
}

func Close(ID string) {
	he, err := hex.DecodeString(ID)
	if err != nil {
		log.Printf("decoding filehash failed")
		return
	}
	Lib.LocalStorage.DeleteData(storage.FileHash(he))
}
