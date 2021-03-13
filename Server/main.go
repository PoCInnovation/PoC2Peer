package main

import (
	"flag"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
	"github.com/PoCInnovation/PoC2Peer/Server/core"
	_ "github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	DefaultP2PPort  = 5000
	DefaultHttpPort = 5001
)

//ifaces, err := net.Interfaces()
//// handle err
//for _, i := range ifaces {
//	addrs, err := i.Addrs()
//	if err != nil {
//		log.Fatal(err)
//	}
//	// handle err
//	for _, addr := range addrs {
//		var ip net.IP
//		switch v := addr.(type) {
//		case *net.IPNet:
//			ip = v.IP
//			fmt.Println(ip)
//		case *net.IPAddr:
//			ip = v.IP
//		}
//		// process IP address
//	}
//}

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	P2PPort := flag.Int("lp", DefaultP2PPort, "Port for P2P server")
	HttpPort := flag.Int("lh", DefaultHttpPort, "Port for http server")
	file := flag.String("f", "", "file to parse at server init")
	trackerPath := flag.String("tp", ".", fmt.Sprintf("Path of tracker file (%s)", p2pnetwork.TrackerFileName))
	flag.Parse()
	trackers, err := p2pnetwork.ParseTrackerInfos(*trackerPath)
	if err != nil {
		log.Fatal(err)
	}
	if *file == "" {
		log.Fatal("No input file, use -f (will change later)")
	}

	s, err := core.NewP2PServer(trackers, *HttpPort, *P2PPort)
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Run(*file); err != nil {
		log.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
	// shut the node down
	if err := s.Close(); err != nil {
		panic(err)
	}
	//makeRoutedHost(*P2PPort, convertPeers([]string{"/ip4/0.0.0.0/tcp/5000"}))
}

func convertPeers(peers []string) []peer.AddrInfo {
	pinfos := make([]peer.AddrInfo, len(peers))
	for i, addr := range peers {
		maddr := ma.StringCast(addr)
		p, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			log.Fatalln(err)
		}
		pinfos[i] = *p
	}
	return pinfos
}
