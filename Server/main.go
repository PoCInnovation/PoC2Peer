package main

import (
	"P2PServer/core"
	"flag"
	"fmt"
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

func main() {

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


	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	P2PPort := flag.Int("lp", DefaultP2PPort, "wait for incoming connections")
	HttpPort := flag.Int("lh", DefaultHttpPort, "wait for incoming connections")
	//
	//priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//opts := []libp2p.Option{
	//	libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", *P2PPort)),
	//	libp2p.Identity(priv),
	//	libp2p.DefaultTransports,
	//	libp2p.DefaultMuxers,
	//	libp2p.DefaultSecurity,
	//	libp2p.NATPortMap(),
	//}
	//
	//basicHost, err := libp2p.New(ctx, opts...)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////_ , err = libp2p.New(ctx, opts...)
	//
	//fmt.Println(basicHost.ID().Pretty())
	//fmt.Println(basicHost.Addrs())
	//
	//for {
	//
	//}


	//fmt.Printf("%v  %T | %v %T", *P2PPort, *P2PPort, *HttpPort, *HttpPort)
	s, err := core.NewP2PServer(*HttpPort, *P2PPort)
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Run(); err != nil {
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
