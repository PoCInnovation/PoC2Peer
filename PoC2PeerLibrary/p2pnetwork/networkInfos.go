package p2pnetwork

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type NetworkInfos struct {
	Ip     string
	IpPub  string
	NWPort int
}

func NewNetworkInfos(ip string, port int) NetworkInfos {
	url1 := "https://api.ipify.org?format=text"
	resp, err := http.Get(url1)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ipPub, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Public IP is: %s\n", ipPub)
	return NetworkInfos{
		Ip:     ip,
		IpPub:  string(ipPub),
		NWPort: port,
	}
}

func (n *NetworkInfos) IP() string {
	return n.Ip
}

func (n *NetworkInfos) URL() string {
	return fmt.Sprintf("%s:%d", n.Ip, n.NWPort)
}

func (n *NetworkInfos) PubIP() string {
	return n.IpPub
}

func (n *NetworkInfos) Port() int {
	return n.NWPort
}
