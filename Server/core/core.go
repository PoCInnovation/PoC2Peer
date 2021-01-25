package core

import (
	"P2PServer/httpHost"
	"P2PServer/p2pHost"
	"bufio"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"log"
)

type P2PServer struct {
	P2PHost  host.Host
	HTTPHost *httpHost.Host
}

func NewP2PServer(HttpPort, P2PPort int) (*P2PServer, error) {
	p2pServer, err := p2pHost.NewP2PHost("127.0.0.1", "tcp", P2PPort)
	if err != nil {
		return nil, err
	}
	httpServer, err := httpHost.NewHttpProvider(HttpPort, p2pServer.ID().String())
	if err != nil {
		return nil, err
	}
	return &P2PServer{P2PHost: p2pServer, HTTPHost: httpServer}, nil
}

func (s *P2PServer) Run() error {
	go func() {
		if err := s.HTTPHost.Run("0.0.0.0:5001"); err != nil {
			log.Fatal(err)
		}
	}()
	s.P2PHost.SetStreamHandler("/echo/1.0.0", func(s network.Stream) {
		log.Println("Got a new stream!")
		if err := doEcho(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	return nil
}

func (s *P2PServer) Close() error {
	if err := s.P2PHost.Close(); err != nil {
		return err
	}
	return nil
}

// doEcho reads a line of data from a stream and writes it back
func doEcho(s network.Stream) error {
	for {
		buf := bufio.NewReader(s)
		str, err := buf.ReadString('\n')
		if err != nil {
			return err
		}

		log.Printf("read: %s\n", str)
		//_, err = s.Write([]byte(str))
	}
}
