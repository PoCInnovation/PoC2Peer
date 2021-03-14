package protocol

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
	"log"
)

// MsgData holds the storage payload of a message
type MsgData interface{}

// Msg holds a protocol message
type Msg struct {
	Op   Opcode
	Data MsgData
}

//func (m Msg) HandleMsg(pStorage storage.LocalStorage) error {
//	switch m.Op {
//	case Have:
//		log.Println("handling HAVE message datagram")
//		//err := msg.HandleDataExchange(c.LocalStorage)
//		//if err != nil {
//		//	return err
//		//}
//	case Data:
//		log.Println("handling DATA message datagram")
//		err := m.HandleDataExchange(pStorage)
//		if err != nil {
//			return err
//		}
//	case Request:
//		log.Printf("handling REQUEST message datagram\n")
//		chunks, err := m.HandleRequest(pStorage
//		if err != nil {
//			log.Println("ERROR WHEN HANDLING REQUEST message datagram")
//			return err
//		}
//		log.Println("finished processing REQUEST message datagram")
// TODO find a way to add protocol (interface ?) to send back datagram
//		return c.network.SendDatagram(chunks, pid)
//	}
//}

func (m *Msg) HandleHave(pid PeerID, lStorage storage.LocalStorage, pStorage storage.PeerStorage) (*Datagram, error) {
	log.Println("handling Have Request")
	have, ok := m.Data.(HaveMsg)
	if !ok {
		return nil, fmt.Errorf("message got DataExchange op Code but could not convert to DataExchange\nreceived: %v", m)
	}
	switch have.Type {
	case HaveRequest:
		chunks, err := lStorage.GetChunkIDsInStorage(have.File)
		if err != nil {
			//	TODO: do smth
			return nil, err
		}
		return NewDataGram(Msg{Op: Have, Data: HaveMsg{File: have.File, Type: HaveResponse, Chunks: chunks}}), nil
	case HaveResponse:
		//log.Println(have.Chunks)
		err := pStorage.AddFileChunksForPeer(pid, have.File, have.Chunks)
		return nil, err
	default:
		return nil, fmt.Errorf("Have got Unknown Type: %v", have.Type)
	}
}

func (m *Msg) HandleDataExchange(pStorage storage.LocalStorage) error {
	exch, ok := m.Data.(DataExchange)
	if !ok {
		return fmt.Errorf("message got DataExchange op Code but could not convert to DataExchange\nreceived: %v", m)
	}
	if len(exch.Chunks) > 0 {
		log.Printf("Received Data for file {%v} -> Chunk : %v to %v\n", exch.File, exch.Chunks[0].Id, exch.Chunks[len(exch.Chunks)-1].Id)
	} else {
		log.Printf("Received Data for file {%v} -> But no chunks requested\n", exch.File)
	}
	return pStorage.AddReceivedFileChunks(exch.File, exch.Chunks)
}

func (m *Msg) HandleRequest(pStorage storage.LocalStorage) (*Datagram, error) {
	req, ok := m.Data.(RequestChunks)
	if !ok {
		return nil, fmt.Errorf("message got DataExchange op Code but could not convert to RequestChunks\nreceived: %v", m)
	}
	data, err := pStorage.GetRequestedChunks(req.File, req.IDs)
	if err != nil {
		log.Printf("HERE: %v\n", err)
		//TODO: better way to send back error
		return NewDataGram(Msg{Op: Error, Data: req}), nil
	}
	nm := Msg{Op: Data, Data: DataExchange{File: req.File, Chunks: data}}
	return NewDataGram(nm), nil
}

// msgAux is an auxiliary struct that looks like Msg except it has
// a []byte to store the incoming gob for MsgData
// (see marshal/unmarshal functions on Msg)
type msgAux struct {
	Op   Opcode
	Data []byte
}

// Decode handles the deserializing of a message.
//
// We can't get away with off-the-shelf JSON, because
// we're using an interface type for MsgData, which causes problems
// on the decode side.
func (m *Msg) UnmarshalJSON(b []byte) error {
	// Use builtin json to unmarshall into aux
	var aux msgAux
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	// The Op field in aux is already what we want for m.Op
	m.Op = aux.Op

	// decode the gob in aux.Data and put it in m.Data
	dec := gob.NewDecoder(bytes.NewBuffer(aux.Data))
	switch aux.Op {
	//case Handshake:
	//	var h HandshakeMsg
	//	err := dec.Decode(&h)
	//	if err != nil {
	//		return errors.New("failed to decode HandshakeMsg")
	//	}
	//	m.Data = h
	case Have:
		var h HaveMsg
		err = dec.Decode(&h)
		if err != nil {
			return errors.New("failed to decode HaveMsg")
		}
		m.Data = h
	case Request:
		var r RequestChunks
		err = dec.Decode(&r)
		if err != nil {
			return errors.New("failed to decode RequestChunks")
		}
		m.Data = r
	case Data:
		var r DataExchange
		err = dec.Decode(&r)
		if err != nil {
			return errors.New("failed to decode DataExchange")
		}
		m.Data = r
	default:
		return errors.New("failed to decode message storage")
	}

	return nil
}

// Encode handles the serializing of a message.
//
// See note above Decode for the reason for the custom Encode
func (m *Msg) MarshalJSON() ([]byte, error) {
	// Encode m.Data into a gob
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	switch m.Data.(type) {
	//case HandshakeMsg:
	//	gob.Register(HandshakeMsg{})
	//	err := enc.Encode(m.Data.(HandshakeMsg))
	//	if err != nil {
	//		return nil, fmt.Errorf("Failed to marshal HandshakeMsg: %v", err)
	//	}
	case HaveMsg:
		// TOdo: Move in init function ?
		gob.Register(HaveMsg{})
		err := enc.Encode(m.Data.(HaveMsg))
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal HaveMsg: %v", err)
		}
	case RequestChunks:
		// TOdo: Move in init function ?
		gob.Register(RequestChunks{})
		err := enc.Encode(m.Data.(RequestChunks))
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal RequestChunks: %v", err)
		}
	case DataExchange:
		// TOdo: Move in init function ?
		gob.Register(DataExchange{})
		err := enc.Encode(m.Data.(DataExchange))
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal DataExchange: %v", err)
		}
	default:
		return nil, errors.New("failed to marshal message storage")
	}

	// build an aux and marshal using built-in json
	aux := msgAux{Op: m.Op, Data: b.Bytes()}
	return json.Marshal(aux)
}

//// SendData sends the chunk range in a storage message
//func SendData(start storage.ChunkID, end storage.ChunkID, remote PeerID, sid SwarmID) error {
//	glog.Infof("SendData Chunks %d-%d, to %v, on %v", start, end, remote)
//	storage, err := swarm.DataFromLocalChunks(start, end)
//	if err != nil {
//		return err
//	}
//	h := storage.DataExchange{Start: start, End: end, Data: storage}
//	m := Msg{Op: Data, Data: h}
//	d := Datagram{ChanID: c.theirs, Msgs: []Msg{m}}
//	return p.sendDatagram(d, ours)
//}
//
//func (p *Ppspp) handleData(cid ChanID, m Msg, remote PeerID) error {
//	glog.Infof("handleData from %v", remote)
//	c, ok1 := p.chans[cid]
//	if !ok1 {
//		return fmt.Errorf("handleData could not find chan %v", cid)
//	}
//	sid := c.sw
//	swarm, ok2 := p.swarms[sid]
//	if !ok2 {
//		return fmt.Errorf("handleData could not find %v", sid)
//	}
//	d, ok3 := m.Data.(storage.DataExchange)
//	if !ok3 {
//		return MsgError{c: cid, m: m, info: "could not convert to DataExchange"}
//	}
//	glog.Infof("recvd storage %d-%d from %v on %v", d.Start, d.End, remote, sid)
//	// TODO: skipping integrity check
//	if err := swarm.AddLocalChunks(d.Start, d.End, d.Data); err != nil {
//		return err
//	}
//	// Invoke the storage handler if we have one
//	if swarm.dataHandler != nil {
//		go swarm.dataHandler(d)
//	}
//	// Send haves to all peers in the swarm
//	for r := range swarm.chans {
//		if err := p.SendHave(d.Start, d.End, r, sid); err != nil {
//			return err
//		}
//	}
//	return nil
//}
