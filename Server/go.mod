module github.com/PoCInnovation/PoC2Peer/Server

go 1.15

require (
	github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/yuin/goldmark v1.3.2 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/tools v0.1.0 // indirect
)

replace github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary => ../PoC2PeerLibrary
