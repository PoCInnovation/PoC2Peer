package p2pnetwork

import "fmt"

type NetworkInfos struct {
	Ip     string
	NWPort int
}

func NewNetworkInfos(ip string, port int) NetworkInfos {
	return NetworkInfos{
		Ip:     ip,
		NWPort: port,
	}
}

func (n *NetworkInfos) IP() string {
	return n.Ip
}

func (n *NetworkInfos) URL() string {
	return fmt.Sprintf("%s:%d", n.Ip, n.NWPort)
}

func (n *NetworkInfos) Port() int {
	return n.NWPort
}
