package main

import (
	"bufio"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/kotlinHandler"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/p2pnetwork"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type prompt struct {
	interrupt chan os.Signal
	msg       chan string
	reader    *bufio.Reader
}

func NewPrompt() *prompt {
	p := new(prompt)
	p.reader = bufio.NewReader(os.Stdin)

	go p.readInput()
	p.interrupt = make(chan os.Signal)
	p.msg = make(chan string)
	signal.Notify(p.interrupt, syscall.SIGINT)
	return p
}

func (p *prompt) readInput() {
	for {
		if in, err := p.reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				p.msg <- "exit"
			} else {
				p.msg <- ""
			}
		} else {
			p.msg <- in
		}
	}
}

func (p *prompt) GetInput() string {
	select {
	case in := <-p.msg:
		return in
	case <-p.interrupt:
		fmt.Println("")
		return ""
	}
}

func (p *prompt) Close() {
	close(p.interrupt)
	close(p.msg)
	signal.Reset()
	fmt.Println("Stopping Shell.")
}

const (
	defaultPort = 5000
	//defaultIP = "78.197.6.119"
	defaultIP = "192.168.0.6"
)

func main() {
	trackers, err := p2pnetwork.ParseTrackerInfos(".")
	if err != nil {
		log.Fatal(err)
	}

	err = kotlinHandler.InitP2PLibrary(p2pnetwork.NewNetworkInfos(defaultIP, defaultPort), trackers)
	if err != nil {
		return
	}
	kotlinHandler.Open("30269e6812313d78c89adc1688e1fdd73d76a79cb2951c0818668c0b96558f02")
	buf := make([]byte, 10016084)
	kotlinHandler.Read(buf, 0, 100, 100, "30269e6812313d78c89adc1688e1fdd73d76a79cb2951c0818668c0b96558f02")
	defer kotlinHandler.CloseP2PLibrary()
}

//func main() {
//	port := flag.Int("p", defaultPort, "Port for P2P server")
//	ip := flag.String("i", defaultIP, "Port for P2P server")
//	file := flag.String("f", "", "file to request at lib init")
//	flag.Parse()
//
//	trackers, err := p2pnetwork.ParseTrackerInfos(".")
//	if err != nil {
//		log.Fatal(err)
//	}
//	lib, err := core.NewP2PPeer(trackers, p2pnetwork.NewNetworkInfos(*ip, *port), "tcp")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer lib.Close()
//	if err = lib.Launch(); err != nil {
//		log.Println(err)
//		return
//	}
//	if *file != "" {
//		if err = lib.TestFile(*file); err != nil {
//			log.Println(err)
//			return
//		}
//	}
//
//	p := NewPrompt()
//	defer p.Close()
//
//	for {
//		fmt.Printf("→ ")
//		input := strings.TrimSuffix(p.GetInput(), "\n")
//		cmd := strings.Fields(input)
//
//		switch {
//		case len(cmd) == 0 || cmd[0] == "exit":
//			return
//		case cmd[0] == "continue":
//		case cmd[0] == "load":
//			if err = loadfiles(lib, cmd[1:]); err != nil {
//				log.Println(err)
//			}
//		default:
//			if err = lib.TestFile(input); err != nil {
//				log.Println(err)
//			}
//		}
//	}
//}

func loadfiles(lib *core.LibP2pCore, files []string) error {
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("Can't read file %s: %v\n", file, err)
		}
		hash, err := lib.LocalStorage.AddFile(content)
		if err != nil {
			return err
		}
		fmt.Printf("File Hashed: %x\n", hash)
	}
	return nil
}
