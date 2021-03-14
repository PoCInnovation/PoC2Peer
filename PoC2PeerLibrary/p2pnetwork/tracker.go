package p2pnetwork

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Tracker interface {
	IP() string
	URL() string
	Port() int
	Ping() error
	AddPeer(peerID, peerIP string, peerPort int) error
	Peers() ([]PeerInfos, error)
	RemovePeer(id string) error
}

const (
	TrackerFileName = "trackers.p2p"
	//trackerPingEndpoint       = "health"
	//trackerAddPeerEndpoint    = "peers/add"
	//trackerRemovePeerEndpoint = "peers/del"
	//trackerListPeersEndpoint  = "peers"
	trackerPingEndpoint       = "health"
	trackerAddPeerEndpoint    = "addPeer"
	trackerRemovePeerEndpoint = "deletePeer"
	trackerListPeersEndpoint  = "peerList"
)

// TODO: Modify after Greg's modifs
type PeerInfos struct {
	Idpeer   string
	Ippeer   string
	Portpeer int
}

//type PeerInfos struct {
//	IP        string
//	Port      int
//	ID        string
//	Transport string
//}

func (i PeerInfos) ID() string {
	return i.Idpeer
}

func (i PeerInfos) IP() string {
	return i.Ippeer
}

func (i PeerInfos) Port() string {
	return fmt.Sprintf("%d", i.Portpeer)
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
		i += 1
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
	log.Println("Pring finished: ", t.URL())
	if req.StatusCode != http.StatusOK {
		log.Println("Pring finished: ", t.URL())
		response, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		return errors.New(string(response))
	}
	log.Println("Pring finished: ", t.URL())
	return nil
}

// AddPeer: Request Tracker to add a peer
func (t HttpTracker) AddPeer(peerID, peerIP string, peerPort int) error {
	// TODO: Modify after Greg's modifs
	url1 := fmt.Sprintf("%s/%s", t.URL(), trackerAddPeerEndpoint)
	u, err := url.Parse(url1)
	if err != nil {
		return err
	}
	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return err
	}
	values.Set("idpeer", peerID)
	values.Set("ippeer", peerIP)
	values.Set("port", fmt.Sprintf("%d", peerPort))
	u.RawQuery = values.Encode()
	req, err := http.Get(fmt.Sprintf("%v", u))
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

	//url := fmt.Sprintf("%s/%s", t.URL(), trackerAddPeerEndpoint)
	//b, err := json.Marshal(&PeerInfos{IP: peerIP, ID: peerID})
	//if err != nil {
	//	return err
	//}
	//buf := bytes.NewBuffer(b)
	//req, err := http.Post(url, "application/json", buf)
	//if err != nil {
	//	return err
	//}
	//if req.StatusCode != http.StatusOK {
	//	response, err := ioutil.ReadAll(req.Body)
	//	if err != nil {
	//		return err
	//	}
	//	return errors.New(string(response))
	//}
	//return nil
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
func (t HttpTracker) RemovePeer(pid string) error {
	// TODO: Modify after Greg's modifs
	url1 := fmt.Sprintf("%s/%s", t.URL(), trackerRemovePeerEndpoint)
	u, err := url.Parse(url1)
	if err != nil {
		return err
	}
	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return err
	}
	values.Set("idpeer", pid)
	u.RawQuery = values.Encode()
	req, err := http.Get(fmt.Sprintf("%v", u))
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

	//url := fmt.Sprintf("%s/%s/%s", t.URL(), trackerRemovePeerEndpoint, id)
	//req, err := http.Post(url, "", nil)
	//if err != nil {
	//	return err
	//}
	//if req.StatusCode != http.StatusOK {
	//	response, err := ioutil.ReadAll(req.Body)
	//	if err != nil {
	//		return err
	//	}
	//	return errors.New(string(response))
	//}
	//return nil
}
