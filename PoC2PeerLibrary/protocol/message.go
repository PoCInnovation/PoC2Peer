package protocol

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
)

// MsgData holds the storage payload of a message
type MsgData interface{}

// Msg holds a protocol message
type Msg struct {
	Op   Opcode
	Data MsgData
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
	case Sync:
		var r SyncMsg
		err = dec.Decode(&r)
		if err != nil {
			return errors.New("failed to decode SyncMsg")
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
	case SyncMsg:
		// TOdo: Move in init function ?
		gob.Register(SyncMsg{})
		err := enc.Encode(m.Data.(SyncMsg))
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
