package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"log"
	"os"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/multiformats/go-multiaddr"
)

// makeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress. It won't encrypt the connection if insecure is true.
//func makeBasicHost(listenPort int, insecure bool, randseed int64) (host.Host, error) {
//
//	// If the seed is zero, use real cryptographic randomness. Otherwise, use a
//	// deterministic randomness source to make generated keys stay the same
//	// across multiple runs
//	var r io.Reader
//	if randseed == 0 {
//		r = rand.Reader
//	} else {
//		r = mrand.New(mrand.NewSource(randseed))
//	}
//
//	// Generate a key pair for this host. We will use it at least
//	// to obtain a valid host ID.
//	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
//	if err != nil {
//		return nil, err
//	}
//
//	opts := []libp2p.Option{
//		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
//		libp2p.Identity(priv),
//		libp2p.DisableRelay(),
//	}
//
//	if insecure {
//		opts = append(opts, libp2p.NoSecurity)
//	}
//
//	basicHost, err := libp2p.New(context.Background(), opts...)
//	if err != nil {
//		return nil, err
//	}
//
//	// Build host multiaddress
//	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))
//
//	// Now we can build a full multiaddress to reach this host
//	// by encapsulating both addresses:
//	addr := basicHost.Addrs()[0]
//	fullAddr := addr.Encapsulate(hostAddr)
//	log.Printf("I am %s\n", fullAddr)
//	if insecure {
//		log.Printf("Now run \"./echo -l %d -d %s -insecure\" on a different terminal\n", listenPort+1, fullAddr)
//	} else {
//		log.Printf("Now run \"./echo -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
//	}
//
//	return basicHost, nil
//}

func main() {
	//if len(os.Args) < 2 {
	//	log.Fatal(errors.New("missing port"))
	//}
	peerID := flag.String("i", "", "wait for incoming connections")
	flag.Parse()
	if *peerID == "" {
		log.Fatal(errors.New("no peerID"))
	}
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		log.Fatal(err)
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", 4000)),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
	}

	ctx := context.Background()
	node, err := libp2p.New(ctx,
		opts...,
	//libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
	//libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/5000"),
	)
	if err != nil {
		log.Fatal(err)
	}

	//node.SetStreamHandler("/echo/1.0.0", func(s network.Stream) {
	//	log.Println("Got a new stream!")
	//	if err := doEcho(s); err != nil {
	//		log.Println(err)
	//		s.Reset()
	//	} else {
	//		s.Close()
	//	}
	//})

	// The following code extracts target's the peer ID from the
	// given multiaddress
	ipfsaddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/p2p/%s", 5000, *peerID))
	if err != nil {
		log.Fatalln(err)
	}

	//pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	pid, err := ipfsaddr.ValueForProtocol(multiaddr.P_P2P)
	if err != nil {
		log.Fatalln(err)
	}

	peerid, err := peer.Decode(pid)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("PeerId", peerid)

	// Decapsulate the /ipfs/<peerID> part from the target
	// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", pid))
	fmt.Println(targetPeerAddr)
	//targetPeerAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/p2p/%s", 5000, *peerID))
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)
	fmt.Println(targetAddr)

	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	node.Peerstore().AddAddr(peerid, targetAddr, peerstore.PermanentAddrTTL)
	//node.SetStreamHandler("/echo/1.0.0", func(s network.Stream) {
	//	log.Println("Got a new stream!")
	//	if err := doEcho(s); err != nil {
	//		log.Println(err)
	//		s.Reset()
	//	} else {
	//		s.Close()
	//	}
	//})

	//// Turn the destination into a multiaddr.
	//maddr, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/QmdTu7tmERainuy3Cm3RmYAwRXwgTyxC1duGpESJ3nXVPW")
	//if err != nil {
	//	log.Fatalln("multiaddr.NewMultiaddr" + err.Error())
	//}
	//
	//// Extract the peer ID from the multiaddr.
	//info, err := peer.AddrInfoFromP2pAddr(maddr)
	//if err != nil {
	//	log.Fatalln("peer.AddrInfoFromP2pAddr" + err.Error())
	//}
	//
	//// Add the destination's peer multiaddress in the peerstore.
	//// This will be used during connection and stream creation by libp2p.
	//node.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
	s, err := node.NewStream(ctx, peerid, "/echo/1.0.0")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(s)
	for {
		writeData(writer)
		time.Sleep(time.Second * 5)
	}

	//// Create a buffered stream so that read and writes are non blocking.
	//rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	//
	/////ip4/127.0.0.1/tcp/61790/ipfs/
	//addr, err := multiaddr.NewMultiaddr("ip4/127.0.0.1/tcp/" + os.Args[1])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//target, err := peer.AddrInfoFromP2pAddr(addr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = node.Connect(ctx, *target)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//node.
	//s, err := target.NewStream(ctx, peerid, "/echo/1.0.0")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//_, err = s.Write([]byte("Hello, world!\n"))
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//out, err := ioutil.ReadAll(s)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("read reply: %q\n", out)

	//// LibP2P code uses golog to log messages. They log with different
	//// string IDs (i.e. "swarm"). We can control the verbosity level for
	//// all loggers with:
	//golog.SetAllLoggers(golog.LogLevel(gologging.INFO)) // Change to DEBUG for extra info
	//
	//// Parse options from the command line
	//listenF := flag.Int("l", 0, "wait for incoming connections")
	//target := flag.String("d", "", "target target to dial")
	//insecure := flag.Bool("insecure", false, "use an unencrypted connection")
	//seed := flag.Int64("seed", 0, "set random seed for peerID generation")
	//flag.Parse()
	//
	//if *listenF == 0 {
	//	log.Fatal("Please provide a port to bind on with -l")
	//}
	//
	//// Make a host that listens on the given multiaddress
	//ha, err := makeBasicHost(*listenF, *insecure, *seed)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Set a stream handler on host A. /echo/1.0.0 is
	//// a user-defined protocol name.
	//ha.SetStreamHandler("/echo/1.0.0", func(s network.Stream) {
	//	log.Println("Got a new stream!")
	//	if err := doEcho(s); err != nil {
	//		log.Println(err)
	//		s.Reset()
	//	} else {
	//		s.Close()
	//	}
	//})
	//
	//if *target == "" {
	//	log.Println("listening for connections")
	//	select {} // hang forever
	//}
	/**** This is where the listener code ends ****/

	// The following code extracts target's the target ID from the
	// given multiaddress
	//ipfsaddr, err := ma.NewMultiaddr(*target)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//peerid, err := target.Decode(pid)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//// Decapsulate the /ipfs/<peerID> part from the target
	//// /ip4/<a.b.c.d>/ipfs/<target> becomes /ip4/<a.b.c.d>
	//targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", pid))
	//targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)
	//
	//// We have a target ID and a targetAddr so we add it to the peerstore
	//// so LibP2P knows how to contact it
	////ha.Peerstore().AddAddr(peerid, targetAddr, peerstore.PermanentAddrTTL)
	//
	//log.Println("opening stream")
	//// make a new stream from host B to host A
	//// it should be handled on host A by the handler we set above because
	//// we use the same /echo/1.0.0 protocol
	//
	//fmt.Println("libp2p node address:", addrs[0])
	//
	//s, err := ha.NewStream(context.Background(), peerid, "/echo/1.0.0")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//_, err = s.Write([]byte("Hello, world!\n"))
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//out, err := ioutil.ReadAll(s)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("read reply: %q\n", out)
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

// doEcho reads a line of data a stream and writes it back
func doEcho(s network.Stream) error {
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	log.Printf("read: %s\n", str)
	_, err = s.Write([]byte(str))
	return err
}
