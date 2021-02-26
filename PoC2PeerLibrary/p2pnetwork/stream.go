package p2pnetwork

import (
	"bufio"
	json "encoding/json"
	inet "github.com/libp2p/go-libp2p-core/network"
)

// WrappedStream wraps a libp2p stream. We encode/decode whenever we
// write/read from a stream, so we can just carry the encoders
// and bufios with us
type WrappedStream struct {
	Stream inet.Stream
	Enc    *json.Encoder
	Dec    *json.Decoder
	W      *bufio.Writer
	R      *bufio.Reader
}

// WrapStream takes a stream and complements it with r/w bufios and
// decoder/encoder. In order to write raw storage to the stream we can use
// wrap.w.Write(). To encode something into it we can wrap.enc.Encode().
// Finally, we should wrap.w.Flush() to actually send the storage. Handling
// incoming storage works similarly with wrap.r.Read() for raw-reading and
// wrap.dec.Decode() to decode.
func WrapStream(s inet.Stream) *WrappedStream {
	reader := bufio.NewReader(s)
	writer := bufio.NewWriter(s)
	dec := json.NewDecoder(reader)
	enc := json.NewEncoder(writer)
	return &WrappedStream{
		Stream: s,
		R:      reader,
		W:      writer,
		Enc:    enc,
		Dec:    dec,
	}
}

func (s *WrappedStream) Close() error {
	return s.Stream.Close()
}
