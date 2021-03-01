package kotlinHandler

import (
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
	"log"
)

var Lib *core.LibP2pCore

type FileDatas struct {
	Cursor int
	Data   []byte
}

type Storage map[string]FileDatas

type SoundBuffer struct {
	Storage        Storage
	CurrentTrackID string
}

var buffer = SoundBuffer{
	make(Storage),
	"",
}

func InitP2PLibrary(infos p2pnetwork.NetworkInfos, trackers []p2pnetwork.Tracker) (err error) {
	Lib, err = core.NewP2PPeer(trackers, infos, "tcp")
	return err
}

func RequestFile(ID string) []byte {
	data, err := Lib.RequestFile(storage.FileHash(ID))
	if err != nil {
		log.Fatal(err)
	}
	var cursor int
	if buffer.CurrentTrackID != ID {
		if cur, ok := buffer.Storage[buffer.CurrentTrackID]; ok {
			cur.Cursor = 0
			buffer.Storage[buffer.CurrentTrackID] = cur
		}
		buffer.CurrentTrackID = ID
		buffer.Storage[buffer.CurrentTrackID] = FileDatas{
			Cursor: cursor,
			Data:   data,
		}
		return data
	} else {
		cur, ok := buffer.Storage[buffer.CurrentTrackID]
		if !ok {
			log.Fatal("ERROR")
		}
		data = data[cur.Cursor:]
		cur.Cursor += len(data)
		buffer.Storage[buffer.CurrentTrackID] = cur
		return data
	}
	return []byte("")
}
