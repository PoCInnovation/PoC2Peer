module github.com/PoCInnovation/PoC2Peer/Poc2PeerServer

go 1.15

require (
	github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary v0.0.0
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/libp2p/go-libp2p v0.13.0 // indirect
	github.com/libp2p/go-libp2p-core v0.8.5 // indirect
	github.com/multiformats/go-multiaddr v0.3.1
)

replace github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary => ../Poc2PeerLibrary
