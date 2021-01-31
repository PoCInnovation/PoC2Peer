package network

import (
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/data"
)

// RequestMsg holds a have message data payload
type RequestMsg struct {
	Start data.ChunkID
	End   data.ChunkID
}
