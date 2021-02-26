package protocol

import "github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"

// HaveRequest holds a have message storage payload
type HaveRequest struct {
	File storage.FileHashTmp
}

// HaveResponse holds a have message storage payload
type HaveResponse struct {
	HaveRequest
	Chunks []storage.ChunkID
}
