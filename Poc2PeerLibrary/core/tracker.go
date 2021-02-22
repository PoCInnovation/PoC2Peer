package core

import (
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
)

type Tracker interface {
	IP() string
	URL() string
	Port() int
	// TODO: remove once swarm
	HTTPIP() string
	HTTPPort() int
	HTTPURL() string
}

// TODO: change once swarm
type httpTracker struct {
	secure  bool
	tracker *p2pnetwork.NetworkInfos
	http    *p2pnetwork.NetworkInfos
}

// NewHttpTracker Create a new htttp tracker (temporary type before fuguring libp2pswarm)
// Allows Node to request permanent peerID from htttp server
func NewHttpTracker(trackerIP string, trackerPort int, httpIP string, httpPort int, secure bool) Tracker {
	return &httpTracker{
		secure:  secure,
		tracker: p2pnetwork.NewNetworkInfos(trackerIP, trackerPort),
		http:    p2pnetwork.NewNetworkInfos(httpIP, httpPort),
	}
}

// IP Return Tracker ip like <192.168.0.31>
func (t httpTracker) IP() string {
	return t.tracker.IP()
}

// Port Return Tracker port like <5001>
func (t httpTracker) Port() int {
	return t.tracker.Port()
}

// URL Return Tracker url like <192.168.0.31:5001>
func (t httpTracker) URL() string {
	if t.secure {
		return fmt.Sprintf("https://%s", t.tracker.URL())
	}
	return fmt.Sprintf("http://%s", t.tracker.URL())
}

func (t httpTracker) HTTPIP() string {
	return t.http.IP()
}

func (t httpTracker) HTTPPort() int {
	return t.http.Port()
}

func (t httpTracker) HTTPURL() string {
	if t.secure {
		return fmt.Sprintf("https://%s", t.http.URL())
	}
	return fmt.Sprintf("http://%s", t.http.URL())
}
