package protocol

import (
	"fmt"
	"github.com/PoCInnovation/PoC2Peer/PoC2PeerLibrary/storage"
	"log"
)

const (
	SyncRequest = iota
	SyncResponse
)

type SyncMsg struct {
	Type  int
	Files []FileInfos
}

// FileInfos holds Informations About a file in Local Storage
type FileInfos struct {
	File     storage.FileHash
	Size     int
	ChunkIds []storage.ChunkID
}

func (c FileInfos) String() string {
	return fmt.Sprintf("P2PFile: %v, Size : %d, Id's: %v", c.File, c.Size, c.ChunkIds)
}

func (m *Msg) HandleSync(pid PeerID, lStorage storage.LocalStorage, pStorage storage.PeerStorage) (*Datagram, error) {
	log.Println("handling Sync Request")
	sync, ok := m.Data.(SyncMsg)
	if !ok {
		return nil, fmt.Errorf("message got DataExchange op Code but could not convert to DataExchange\nreceived: %v", m)
	}
	switch sync.Type {
	case SyncRequest:
		files := lStorage.FilesList()
		if len(files) <= 0 {
			log.Println("No files in storage for sync msg")
			return nil, nil
		}
		fileInfosList := make([]FileInfos, len(files))
		for i, fileHash := range files {
			ids, filesz, err := lStorage.FileInfos(fileHash)
			if err != nil {
				return nil, fmt.Errorf("Error when getting file infos in sync Request: %v", err)
			}
			fileInfosList[i] = FileInfos{
				File:     fileHash,
				Size:     filesz,
				ChunkIds: ids,
			}
		}
		return NewDataGram(Msg{Op: Sync, Data: SyncMsg{Type: SyncResponse, Files: fileInfosList}}), nil
	case SyncResponse:
		syncRqst := make([]Msg, len(sync.Files))
		for i, fileInfo := range sync.Files {
			pStorage.UpdateFileInfos(pid, fileInfo.File, fileInfo.Size, fileInfo.ChunkIds)
			syncRqst[i] = Msg{Op: Request, Data: RequestChunks{
				File: fileInfo.File,
				IDs:  fileInfo.ChunkIds[:len(fileInfo.ChunkIds)/4],
			}}
		}
		return NewDataGram(syncRqst...), nil
	default:
		return nil, fmt.Errorf("Have got Unknown Type: %v", sync.Type)
	}
}
