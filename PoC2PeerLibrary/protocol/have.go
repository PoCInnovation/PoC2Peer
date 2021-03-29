package protocol

import (
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
	"log"
)

const (
	HaveRequest = iota
	HaveResponse
)

// HaveMsg holds a have message storage payload
type HaveMsg struct {
	File     storage.FileHash
	Type     int
	FileSize int
	Chunks   []storage.ChunkID
}

func (m *Msg) HandleHave(pid PeerID, lStorage storage.LocalStorage, pStorage storage.PeerStorage) (*Datagram, error) {
	log.Println("handling Have Request")
	have, ok := m.Data.(HaveMsg)
	if !ok {
		return nil, fmt.Errorf("message got DataExchange op Code but could not convert to DataExchange\nreceived: %v", m)
	}
	switch have.Type {
	case HaveRequest:
		log.Printf("have request for file: %v\n", have.File.Decode())
		chunks, sz, err := lStorage.FileInfos(have.File)
		if err != nil {
			return nil, err
		}
		return NewDataGram(Msg{Op: Have, Data: HaveMsg{File: have.File, Type: HaveResponse, Chunks: chunks, FileSize: sz}}), nil
	case HaveResponse:
		err := pStorage.UpdateFileInfos(pid, have.File, -1, have.Chunks)
		return nil, err
	default:
		return nil, fmt.Errorf("Have got Unknown Type: %v", have.Type)
	}
}
