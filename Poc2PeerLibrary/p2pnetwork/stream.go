package p2pnetwork

import (
	"bufio"
	inet "github.com/libp2p/go-libp2p-core/network"

	//inet "github.com/libp2p/go-libp2p-net"
	//multicodec "github.com/multiformats/go-multicodec"
	//json "github.com/multiformats/go-multicodec/json"
	json "encoding/json"
)

// WrappedStream wraps a libp2p stream. We encode/decode whenever we
// write/read from a stream, so we can just carry the encoders
// and bufios with us
type WrappedStream struct {
	stream inet.Stream
	enc    *json.Encoder
	dec    *json.Decoder
	//enc    multicodec.Encoder
	//dec    multicodec.Decoder
	w *bufio.Writer
	r *bufio.Reader
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

	// Note that if these change, then the Encode/Decode functions for Msg
	// may no longer get called, which may mess up the codec for Msg.Data
	dec := json.NewDecoder(reader)
	enc := json.NewEncoder(writer)
	return &WrappedStream{
		stream: s,
		r:      reader,
		w:      writer,
		enc:    enc,
		dec:    dec,
	}
}
