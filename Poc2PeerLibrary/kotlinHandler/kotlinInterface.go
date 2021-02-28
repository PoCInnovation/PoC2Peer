package kotlinHandler

import (
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
	"log"
)

var Lib *core.LibP2pCore

type FileDatas struct {
	Cursor int
	Data   []byte
}

type Storage map[int]FileDatas

type SoundBuffer struct {
	Storage        Storage
	CurrentTrackID int
}

var buffer = SoundBuffer{
	make(Storage),
	-1,
}

func InitP2PLibrary(infos p2pnetwork.NetworkInfos, trackers []p2pnetwork.Tracker) (err error) {
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

func RequestFile(ID int) []byte {
	data, err := Lib.RequestFile(storage.FileHashTmp(ID))
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
