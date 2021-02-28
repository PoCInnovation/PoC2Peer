package protocol

import (
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
)

// RequestChunks holds a have message storage payload
type RequestChunks struct {
	File storage.FileHashTmp
	IDs  []storage.ChunkID
}

func (c RequestChunks) String() string {
	return fmt.Sprintf("File: %v, Id's: %v", c.File, c.IDs)
}
