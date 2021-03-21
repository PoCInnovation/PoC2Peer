package protocol

import (
	"fmt"
)

// Opcode identifies the type of message
type Opcode uint8

// From the RFC:
//   +----------+------------------+
//   | Msg Type | Description      |
//   +----------+------------------+
//   | 0        | HANDSHAKE        |
//   | 1        | DATA             |
//   | 2        | ACK              |
//   | 3        | HAVE             |
//   | 4        | INTEGRITY        |
//   | 5        | PEX_RESv4        |
//   | 6        | PEX_REQ          |
//   | 7        | SIGNED_INTEGRITY |
//   | 8        | REQUEST          |
//   | 9        | CANCEL           |
//   | 10       | CHOKE            |
//   | 11       | UNCHOKE          |
//   | 12       | PEX_RESv6        |
//   | 13       | PEX_REScert      |
//   | 14-254   | Unassigned       |
//   | 255      | Reserved         |
//   +----------+------------------+

const (
	Handshake Opcode = 0
	Data      Opcode = 1
	Have      Opcode = 3
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
