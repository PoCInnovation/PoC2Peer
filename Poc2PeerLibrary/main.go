package main

import (
	"flag"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()
	tracker := p2pnetwork.NewHttpTracker("192.168.0.31", 5000, "192.168.0.31", 5001, false)
	lib, err := core.NewP2PPeer(tracker, p2pnetwork.NewNetworkInfos("0.0.0.0", 4000), "tcp")
	if err != nil {
		log.Fatal(err)
	}

	err = lib.Launch()
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
	lib.N.Close()
	// shut the node down
	//if err := s.Close(); err != nil {
	//	panic(err)
	//}
	//makeRoutedHost(*P2PPort, convertPeers([]string{"/ip4/0.0.0.0/tcp/5000"}))

}
