package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/p2pnetwork"
	"log"
	"os"
)

func main() {
	flag.Parse()
	tracker := core.NewHttpTracker("192.168.0.31", 5000, "192.168.0.31", 5001, false)
	lib, err := core.NewP2PPeer(tracker, p2pnetwork.NewNetworkInfos("0.0.0.0", 4000))
	if err != nil {
		log.Fatal(err)
	}
	err = lib.Launch()
	if err != nil {
		log.Fatal(err)
	}
	////if len(os.Args) < 2 {
	////	log.Fatal(errors.New("missing port"))
	////}
	//peerID := flag.String("i", "", "wait for incoming connections")
	//flag.Parse()
	////if *peerID == "" {
	////	log.Fatal(errors.New("no peerID"))
	////}
	//
	//priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//opts := []libp2p.Option{
	//	libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", 4000)),
	//	libp2p.Identity(priv),
	//	libp2p.DisableRelay(),
	//}
	//
	//ctx := context.Background()
	//node, err := libp2p.New(ctx,
	//	opts...,
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//res, err := http.Get(httpEndpoint)
	//if err != nil {
	//	log.Fatal(err)
	//} else if res.StatusCode != http.StatusOK {
	//	log.Fatal("res not OK")
	//}
	//byteID, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//*peerID = string(byteID)
	//
	//
	//// The following code extracts target's the peer ID from the
	//// given multiaddress
	//ipfsaddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d/p2p/%s", 5000, *peerID))
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	////pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	//pid, err := ipfsaddr.ValueForProtocol(multiaddr.P_P2P)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	////peerid, err := peer.Decode(pid)
	//peerid, err := peer.Decode(*peerID)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println("PeerId", peerid)
	//
	//// Decapsulate the /ipfs/<peerID> part from the target
	//// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
	//targetPeerAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", pid))
	//fmt.Println(targetPeerAddr)
	////targetPeerAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/p2p/%s", 5000, *peerID))
	//targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)
	//fmt.Println(targetAddr)
	//
	//// We have a peer ID and a targetAddr so we add it to the peerstore
	//// so LibP2P knows how to contact it
	//node.Peerstore().AddAddr(peerid, targetAddr, peerstore.PermanentAddrTTL)
	//
	//// Start a stream with the destination.
	//// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
	//s, err := node.NewStream(ctx, peerid, "/echo/1.0.0")
	//if err != nil {
	//	panic(err)
	//}
	//writer := bufio.NewWriter(s)
	//for {
	//	writeData(writer)
	//}
}

func writeData(rw *bufio.Writer) error {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			return err
		}
		_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
		if err != nil {
			return err
		}
		err = rw.Flush()
		if err != nil {
			return err
		}
	}
}
