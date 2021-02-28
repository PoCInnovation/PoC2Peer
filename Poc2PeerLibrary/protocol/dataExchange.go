package protocol

import (
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
)

type DataExchange struct {
	File   storage.FileHashTmp
	Chunks []storage.Chunk
}
