package storage

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type LocalStorageConfig struct {
	ChunkSize int
}

type P2PStorage struct {
	Config     LocalStorageConfig
	LocalFiles map[string][]*Chunk
}

// AddFile Add a file to local storage. Return the hashed file when successfull
func (s *P2PStorage) AddFile(fileData []byte) ([]byte, error) {
	hash := newHashFromFile(fileData)
	key := string(hash)
	if _, ok := s.LocalFiles[key]; ok {
		return nil, errors.New("Trying to add existing file")
	}
	s.LocalFiles[key] = FileDataToChunks(fileData, s.Config.ChunkSize)
	return hash, nil
}

// GetChunks Search for requested chunk with file hash
func (s *P2PStorage) GetChunks(hash []byte, start, end ChunkID) ([]*Chunk, error) {
	data, ok := s.LocalFiles[string(hash)]
	if ok {
		return nil, fmt.Errorf("Requested file is not in storage: {%x}", hash)
	}
	dataLen := uint32(len(data))
	if dataLen < uint32(end) || start > end {
		return nil, fmt.Errorf("Requested Chuncks out of range: [ len: %d, start: %v, end: %v", dataLen, start, end)
	}
	return data[start:end], nil
}

// GetChunks Search for requested chaunk with file hash
func (s *P2PStorage) GetDataFromLocalChunk(hash []byte, start, end ChunkID) ([]byte, error) {
	var dataLen int
	chunks, err := s.GetChunks(hash, start, end)
	if err != nil {
		return nil, err
	}
	for _, chunk := range chunks {
		dataLen += len(chunk.B)
	}
	data := make([]byte, dataLen)
	dataLen = 0
	for _, chunk := range chunks {
		copy(data[dataLen:], chunk.B)
		dataLen += len(chunk.B)
	}
	return data, nil
}

func newHashFromFile(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
