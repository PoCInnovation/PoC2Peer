package protocol

import (
	"github.com/PoCInnovation/PoC2Peer/Poc2PeerLibrary/storage"
	"log"
)

type DataExchange struct {
	File   storage.FileHashTmp
	Start  storage.ChunkID
	End    storage.ChunkID
	Chunks []storage.Chunk
}

func (e DataExchange) AddReceivedDatasToStorage(pStorage storage.LocalStorage) error {
	log.Printf("Received Chunks for file : %v", e.File)
	return pStorage.AddReceivedFileChunks(e.File, e.Chunks)
}

func NewDataExchangeFromRequested(file storage.FileHashTmp, chunks []storage.Chunk) DataExchange {
	start := chunks[0].Id
	end := chunks[len(chunks)-1].Id
	return DataExchange{File: file, Start: start, End: end, Chunks: chunks}
}

// SendData sends the chunk range in a storage message
func GetDataForRequestedChunks(storage storage.P2PStorage, hash []byte, start, end storage.ChunkID) error {
	//glog.Infof("SendData Chunks %d-%d, to %v, on %v", start, end, remote)
	//data, err := storage.GetDataFromLocalChunks(hash, start, end)
	//if err != nil {
	//	return err
	//}
	//h := DataExchange{Start: start, End: end, Data: data}
	//m := Msg{Op: Data, Data: h}
	//d := NewDataGram(m)
	//return p.sendDatagram(d, ours)
	return nil
}
