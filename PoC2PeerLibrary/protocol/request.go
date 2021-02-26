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
	return fmt.Sprintf("File: %v, Id's: %v", c.File, c.IDs)
}
