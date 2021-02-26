package protocol

import "github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"

const (
	HaveRequest = iota
	HaveResponse
)

// HaveMsg holds a have message storage payload
type HaveMsg struct {
	File   storage.FileHash
	Type   int
	Chunks []storage.ChunkID
}
