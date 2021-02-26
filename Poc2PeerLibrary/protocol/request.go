package protocol

import (
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
)

// RequestChunks holds a have message storage payload
type RequestChunks struct {
	File  storage.FileHashTmp
	Start storage.ChunkID
	End   storage.ChunkID
}

func (c RequestChunks) String() string {
	return fmt.Sprintf("File: %v, Start: %v, End: %v", c.File, c.Start, c.End)
}
