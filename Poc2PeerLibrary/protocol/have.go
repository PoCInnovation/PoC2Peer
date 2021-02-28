package protocol

import "github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"

const (
	HaveRequest = iota
	HaveResponse
)

// HaveMsg holds a have message storage payload
type HaveMsg struct {
	File   storage.FileHashTmp
	Type   int
	Chunks []storage.ChunkID
}
