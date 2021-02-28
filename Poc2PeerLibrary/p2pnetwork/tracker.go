package p2pnetwork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Tracker interface {
	IP() string
	URL() string
	Port() int
	PeerID() PeerID
	// TODO: remove once swarm
	HTTPIP() string
	HTTPPort() int
	HTTPURL() string
}

const trackerFileName = "trackers.p2p"

// TODO: change once swarm
type HttpTracker struct {
	Secure  bool
	Tracker NetworkInfos
	Http    NetworkInfos
	peerid  PeerID
}

// NewHttpTracker Create a new htttp tracker (temporary type before fuguring libp2pswarm)
// Allows Node to request permanent peerID from htttp server
func NewHttpTracker(trackerIP string, trackerPort int, httpIP string, httpPort int, secure bool) Tracker {
	return &HttpTracker{
		Secure:  secure,
		Tracker: NewNetworkInfos(trackerIP, trackerPort),
		Http:    NewNetworkInfos(httpIP, httpPort),
	}
}

func ParseTrackerInfos(strorageDir string) ([]HttpTracker, error) {
	bytes, err := ioutil.ReadFile(strorageDir + "/" + trackerFileName)
	if err != nil {
		return nil, err
	}
	var trackers []HttpTracker
	if err = json.Unmarshal(bytes, &trackers); err != nil {
		return nil, err
	}
	return trackers, nil
}

// IP Return Tracker ip like <192.168.0.31>
func (t HttpTracker) IP() string {
	return t.Tracker.IP()
}

// Port Return Tracker port like <5001>
func (t HttpTracker) Port() int {
	return t.Tracker.Port()
}

// Port Return Tracker port like <5001>
func (t HttpTracker) PeerID() PeerID {
	return t.peerid
}

// URL Return Tracker url like <192.168.0.31:5001>
func (t HttpTracker) URL() string {
	if t.Secure {
		return fmt.Sprintf("https://%s", t.Tracker.URL())
	}
	return fmt.Sprintf("http://%s", t.Tracker.URL())
}

func (t HttpTracker) HTTPIP() string {
	return t.Http.IP()
}

func (t HttpTracker) HTTPPort() int {
	return t.Http.Port()
}

func (t HttpTracker) HTTPURL() string {
	if t.Secure {
		return fmt.Sprintf("https://%s", t.Http.URL())
	}
	return fmt.Sprintf("http://%s", t.Http.URL())
}
