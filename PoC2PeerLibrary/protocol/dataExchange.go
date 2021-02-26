package protocol

import (
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
)

type DataExchange struct {
	File   storage.FileHash
	Chunks []storage.Chunk
}
