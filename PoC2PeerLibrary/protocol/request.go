package protocol

import (
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
)

// RequestChunks holds a have message storage payload
type RequestChunks struct {
	File storage.FileHash
	IDs  []storage.ChunkID
}

func (c RequestChunks) String() string {
	return fmt.Sprintf("P2PFile: %v, Id's: %v", c.File, c.IDs)
}

func (m *Msg) HandleRequest(pStorage storage.LocalStorage) (*Datagram, error) {
	req, ok := m.Data.(RequestChunks)
	if !ok {
		return nil, fmt.Errorf("message got DataExchange op Code but could not convert to RequestChunks\nreceived: %v", m)
	}
	data, err := pStorage.GetRequestedChunks(req.File, req.IDs)
	if err != nil {
		//TODO: better way to send back error
		return NewDataGram(Msg{Op: Error, Data: req}), nil
	}
	nm := Msg{Op: Data, Data: DataExchange{File: req.File, Chunks: data}}
	return NewDataGram(nm), nil
}
