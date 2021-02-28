package core

import (
	"fmt"
	p2pcore "github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/core"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerServer/httpHost"
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerServer/p2pHost"
	"io/ioutil"
	"log"
)

type P2PServer struct {
	P2PHost  *p2pcore.LibP2pCore
	HTTPHost *httpHost.Host
}

func NewP2PServer(HttpPort, P2PPort int) (*P2PServer, error) {
	p2pServer, err := p2pHost.NewP2PHost("0.0.0.0", "tcp", P2PPort)
	if err != nil {
		return nil, err
	}
	httpServer, err := httpHost.NewHttpProvider(HttpPort, p2pServer.ID())
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
	if err := s.P2PHost.SetStreamHandlers(); err != nil {
		return err
	}
	//size := 400
	//chunks := make([]byte, size)
	//for n := 0; n < size; {
	//	data := fmt.Sprintf("Chunk: %4d|", n)
	//	copy(chunks[n:], data)
	//	n += len(data)
	//}
	//hash, err := s.P2PHost.LocalStorage.AddFile(chunks)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("File Hashed: %v\n", hash)
	files := []string{
		"tests/testsfile",
	}
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("Can't read file %s: %v\n", file, err)
		}
		hash, err := s.P2PHost.LocalStorage.AddFile(content)
		if err != nil {
			return err
		}
		fmt.Printf("File Hashed: %v\n", hash)
	}
	//s.P2PHost.N.Host.SetStreamHandler(protocol.FileTransferProtocol, func(stream network.Stream) {
	//	log.Println("Got a new stream!")
	//	fmt.Println(stream.ID())
	//	s1, err := s.P2PHost.N.Connect(stream.Conn().RemotePeer())
	//	if err != nil {
	//		return
	//	}
	//	defer s1.Stream.Close()
	//
	//
	//	size := 600000
	//	batch := size / 10
	//	chunks:= make([]storage.Chunk, size)
	//	for n := 0; n < size; n += 1 {
	//		data := fmt.Sprintf("Chunk: %d", n)
	//		chunks[n] = storage.NewChunkFromData(storage.ChunkID(n), len(data), []byte(data))
	//	}
	//	for i := 0; i < size; i += batch {
	//		toSend := make([]storage.Chunk, batch)
	//		copy(toSend, chunks)
	//		chunks = chunks[batch:]
	//		h := protocol.DataExchange{File: []byte("File_1"), Start: 0, End: storage.ChunkID(batch - 1), Chunks: toSend}
	//		//log.Printf("Sending DataExchange : %v", h)
	//		m := protocol.Msg{Op: protocol.Data, Data: h}
	//		d := protocol.NewDataGram(m)
	//		err = s1.Enc.Encode(&d)
	//		if err != nil {
	//			log.Println(err)
	//			break
	//		}
	//		err = s1.W.Flush()
	//		if err != nil {
	//			log.Println(err)
	//			break
	//		}
	//	}
	//	log.Printf("stream served: %v\n", stream.ID())
	//	stream.Close()
	//})
	//s.P2PHost.N.Host.SetStreamHandler(protocol.FileTransferProtocol, func(stream network.Stream) {
	//	log.Println("Got a new stream!")
	//	fmt.Println(stream.ID())
	//	s1, err := s.P2PHost.N.Connect(stream.Conn().RemotePeer())
	//	if err != nil {
	//		return
	//	}
	//	defer s1.Stream.Close()
	//
	//
	//	size := 600000
	//	batch := size / 10
	//	chunks:= make([]storage.Chunk, size)
	//	for n := 0; n < size; n += 1 {
	//		data := fmt.Sprintf("Chunk: %d", n)
	//		chunks[n] = storage.NewChunkFromData(storage.ChunkID(n), len(data), []byte(data))
	//	}
	//	for i := 0; i < size; i += batch {
	//		toSend := make([]storage.Chunk, batch)
	//		copy(toSend, chunks)
	//		chunks = chunks[batch:]
	//		h := protocol.DataExchange{File: []byte("File_1"), Start: 0, End: storage.ChunkID(batch - 1), Chunks: toSend}
	//		//log.Printf("Sending DataExchange : %v", h)
	//		m := protocol.Msg{Op: protocol.Data, Data: h}
	//		d := protocol.NewDataGram(m)
	//		err = s1.Enc.Encode(&d)
	//		if err != nil {
	//			log.Println(err)
	//			break
	//		}
	//		err = s1.W.Flush()
	//		if err != nil {
	//			log.Println(err)
	//			break
	//		}
	//	}
	//	log.Printf("stream served: %v\n", stream.ID())
	//	stream.Close()
	//})

	//for i := 0; i < 2; i += 1 {
	//	size := 2000
	//	chunks:= make([]storage.Chunk, size)
	//	for n := 0; n < size; n += 1 {
	//		data := fmt.Sprintf("Chunk: %d", n)
	//		chunks[n] = storage.NewChunkFromData(storage.ChunkID(n), len(data), []byte(data))
	//	}
	//	h := protocol.DataExchange{File: []byte("File_1"), Start: 0, End: 10, Chunks: chunks}
	//	log.Printf("Sending DataExchange : %v", h)
	//	m := protocol.Msg{Op: protocol.Data, Data: h}
	//	d := protocol.NewDataGram(m)
	//	err = ns.Enc.Encode(&d)
	//	if err != nil {
	//		log.Println(err)
	//		break
	//	}
	//	err = ns.W.Flush()
	//	if err != nil {
	//		log.Println(err)
	//		break
	//	}
	//	time.Sleep(time.Second)
	//}

	return nil
}

func (s *P2PServer) Close() error {
	if err := s.P2PHost.N.Close(); err != nil {
		return err
	}
	return nil
}
