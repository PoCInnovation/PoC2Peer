package protocol

import (
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
	"log"
)

type DataExchange struct {
	File   storage.FileHash
	Chunks []storage.Chunk
}

func (m *Msg) HandleDataExchange(pStorage storage.LocalStorage) error {
	exch, ok := m.Data.(DataExchange)
	if !ok {
		return fmt.Errorf("message got DataExchange op Code but could not convert to DataExchange\nreceived: %v", m)
	}
	if len(exch.Chunks) > 0 {
		log.Printf("Received Data for file {%v} -> Chunk : %v to %v\n", exch.File.Decode(), exch.Chunks[0].Id, exch.Chunks[len(exch.Chunks)-1].Id)
	} else {
		log.Printf("Received Data for file {%v} -> But no chunks requested\n", exch.File.Decode())
	}
	return pStorage.AddFileChunks(exch.File, exch.Chunks)
}
