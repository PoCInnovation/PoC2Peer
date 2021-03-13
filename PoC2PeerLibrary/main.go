package main

import (
	"flag"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	file := flag.String("f", "", "file to request at lib init")
	flag.Parse()
	//tracker := p2pnetwork.NewHttpTracker("192.168.0.31", 5001, false)
	//lib, err := core.NewP2PPeer([]p2pnetwork.Tracker{tracker}, p2pnetwork.NewNetworkInfos("0.0.0.0", 4000), "tcp")
	//lib, err := core.NewP2PPeer([]p2pnetwork.Tracker{tracker}, p2pnetwork.NewNetworkInfos("0.0.0.0", 4000), "tcp")
	if *file == "" {
		log.Fatal("No file to request, use -f (will change later)")
	}

	trackers, err := p2pnetwork.ParseTrackerInfos(".")
	if err != nil {
		log.Fatal(err)
	}
	lib, err := core.NewP2PPeer(trackers, p2pnetwork.NewNetworkInfos("192.168.0.6", 4000), "tcp")
	if err != nil {
		log.Fatal(err)
	}

	err = lib.Launch(*file)
	if err != nil {
		log.Println(err)
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
	lib.Close()
}
