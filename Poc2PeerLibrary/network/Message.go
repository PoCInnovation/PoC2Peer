package network

// Opcode identifies the type of message
type Opcode uint8

// MsgData holds the data payload of a message
type MsgData interface{}

// Msg holds a protocol message
type Msg struct {
	Op   Opcode
	Data MsgData
}
