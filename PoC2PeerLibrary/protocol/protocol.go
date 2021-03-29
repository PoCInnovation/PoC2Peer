package protocol

import (
	"fmt"
)

// Opcode identifies the type of message
type Opcode uint8

const (
	Handshake Opcode = 0
	Data      Opcode = 1
	Have      Opcode = 3
	Sync      Opcode = 4
	Request   Opcode = 8
	Error     Opcode = 255
)

const (
	FileTransferProtocol = "/p2p/files/0.0.0"
	HandshakeProtocol    = "/p2p/handshake/0.0.0"
)

// PeerID identifies a peer
// To avoid cycle Import
type PeerID interface {
	String() string
}

type Protocol interface {
	HandleDatagram(d *Datagram, id PeerID)
	SetDatagramSender(f func(Datagram, PeerID) error)
}

// Datagram holds a protocol datagram
type Datagram struct {
	Msgs []Msg
}

func NewDataGram(m ...Msg) *Datagram {
	return &Datagram{Msgs: m}
}

// MsgError is an error that happens while handling an incoming message
type MsgError struct {
	m    Msg
	info string
}

func (e MsgError) Error() string {
	return fmt.Sprintf("message error : %v\n%v", e.info, e.m)
}
