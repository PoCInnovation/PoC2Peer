package protocol

import (
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
)

type DataExchange struct {
	File  []byte
	Start storage.ChunkID
	End   storage.ChunkID
	Data  []byte
}

// SendData sends the chunk range in a storage message
func GetDataForRequestedChunks(storage storage.P2PStorage, hash []byte, start, end storage.ChunkID) error {
	//glog.Infof("SendData Chunks %d-%d, to %v, on %v", start, end, remote)
	//data, err := storage.GetDataFromLocalChunk(hash, start, end)
	//if err != nil {
	//	return err
	//}
	//h := DataExchange{Start: start, End: end, Data: data}
	//m := Msg{Op: Data, Data: h}
	//d := NewDataGram(m)
	//return p.sendDatagram(d, ours)
	return nil
}
