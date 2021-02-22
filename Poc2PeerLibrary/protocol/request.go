package protocol

import (
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
)

// RequestChunks holds a have message storage payload
type RequestChunks struct {
	File  []byte
	Start storage.ChunkID
	End   storage.ChunkID
}
