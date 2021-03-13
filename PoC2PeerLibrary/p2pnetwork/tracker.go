package p2pnetwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Tracker interface {
	IP() string
	URL() string
	Port() int
	Ping() error
	AddPeer(pid, peerURL string) error
	Peers() ([]PeerInfos, error)
	RemovePeer(id string) error
}

const (
	TrackerFileName           = "trackers.p2p"
	trackerPingEndpoint       = "health"
	trackerAddPeerEndpoint    = "peers/add"
	trackerRemovePeerEndpoint = "peers/del"
	trackerListPeersEndpoint  = "peers"
)

type PeerInfos struct {
	IP        string
	Port      int
	ID        string
	Transport string
}

type HttpTracker struct {
	Secure  bool
	Tracker NetworkInfos
}

// NewHttpTracker Create a new htttp tracker
// Allows Node to request permanent otherPeers from http server
func NewHttpTracker(trackerIP string, trackerPort int, secure bool) Tracker {
	return &HttpTracker{
		Secure:  secure,
		Tracker: NewNetworkInfos(trackerIP, trackerPort),
		//Http:    NewNetworkInfos(httpIP, httpPort),
	}
}

func ParseTrackerInfos(strorageDir string) ([]Tracker, error) {
	bytes, err := ioutil.ReadFile(strorageDir + "/" + TrackerFileName)
	if err != nil {
		return nil, err
	}
	var httptrackers []HttpTracker
	if err = json.Unmarshal(bytes, &httptrackers); err != nil {
		return nil, err
	}
	trackers := make([]Tracker, len(httptrackers))
	for i := range httptrackers {
		trackers[i] = httptrackers[i]
	}
	return trackers, nil
}

// IP: Return Tracker NetworkInfos.Ip (Ex: <192.168.0.31>)
func (t HttpTracker) IP() string {
	return t.Tracker.IP()
}

// Port: Return Tracker NetworkInfos.NWPort (Ex: <5001>)
func (t HttpTracker) Port() int {
	return t.Tracker.Port()
}

// URL: Return Tracker URL like (Ex: <192.168.0.31:5001>)
func (t HttpTracker) URL() string {
	if t.Secure {
		return fmt.Sprintf("https://%s", t.Tracker.URL())
	}
	return fmt.Sprintf("http://%s", t.Tracker.URL())
}

// Ping: Check if Tracker is alive
func (t HttpTracker) Ping() error {
	url := fmt.Sprintf("%s/%s", t.URL(), trackerPingEndpoint)
	req, err := http.Get(url)
	if err != nil {
		return err
	}
	if req.StatusCode != http.StatusOK {
		response, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		return errors.New(string(response))
	}
	return nil
}

// AddPeer: Request Tracker to add a peer
func (t HttpTracker) AddPeer(pid, peerURL string) error {
	url := fmt.Sprintf("%s/%s", t.URL(), trackerAddPeerEndpoint)
	b, err := json.Marshal(&PeerInfos{IP: peerURL, ID: pid})
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)
	req, err := http.Post(url, "application/json", buf)
	if err != nil {
		return err
	}
	if req.StatusCode != http.StatusOK {
		response, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		return errors.New(string(response))
	}
	return nil
}

// Peers: Get peers from Tracker.
func (t HttpTracker) Peers() ([]PeerInfos, error) {
	url := fmt.Sprintf("%s/%s", t.URL(), trackerListPeersEndpoint)
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if req.StatusCode != http.StatusOK {
		response, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(response))
	}
	var peers []PeerInfos
	response, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &peers)
	if err != nil {
		return nil, err
	}
	//log.Println(peers)
	return peers, nil
}

// TODO: Modify
// AddPeer: Request Tracker to remove a peer
func (t HttpTracker) RemovePeer(id string) error {
	url := fmt.Sprintf("%s/%s/%s", t.URL(), trackerRemovePeerEndpoint, id)
	req, err := http.Post(url, "", nil)
	if err != nil {
		return err
	}
	if req.StatusCode != http.StatusOK {
		response, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		return errors.New(string(response))
	}
	return nil
}
