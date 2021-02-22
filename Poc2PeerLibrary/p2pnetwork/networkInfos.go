package p2pnetwork

import "fmt"

type NetworkInfos struct {
	ip   string
	port int
}

func NewNetworkInfos(ip string, port int) *NetworkInfos {
	return &NetworkInfos{
		ip:   ip,
		port: port,
	}
}
func (n *NetworkInfos) IP() string {
	return n.ip
}

func (n *NetworkInfos) URL() string {
	return fmt.Sprintf("%s:%d", n.ip, n.port)
}

func (n *NetworkInfos) Port() int {
	return n.port
}
