# Core

## Package Description

## Types
```go
package core

type LibP2pCore struct {
	network      p2pnetwork.Network
	infos        p2pnetwork.NetworkInfos
	trackers     []p2pnetwork.Tracker
	LocalStorage storage.LocalStorage
	PeerStorage  storage.PeerStorage
}
```
## Misc
