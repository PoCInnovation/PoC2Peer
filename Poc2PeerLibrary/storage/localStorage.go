package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"sync"
)

type LocalStorage interface {
	AddFile(fileData []byte) (FileID, error)
	AddReceivedFileChunks(hash FileID, chunks []Chunk) error
	GetChunks(hash FileID, start, end ChunkID) ([]Chunk, error)
	GetFileDatas(hash FileID) ([]byte, error)
}

type FileID interface {
	String() string
	//Decode() FileHash
}

type FileHash []byte

func (h FileHash) String() string {
	return string(h)
}

func (h FileHash) Decode() FileHash {
	return []byte(h)
	he, err := hex.DecodeString(h.String())
	if err != nil {
		log.Printf("decoding filehash failed")
		return []byte("")
	}
	fmt.Println(h)
	return FileHash(he)
	//return fmt.Sprintf("%x", h)
}

const LocalStorageSize = 100

type LocalStorageConfig struct {
	ChunkSize int
}

type P2PStorage struct {
	sync.Mutex
	Config LocalStorageConfig
	//LocalFiles map[string][]*Chunk
	LocalFiles map[string]map[ChunkID]*Chunk
}

func NewP2PStorage() LocalStorage {
	return &P2PStorage{
		Config:     LocalStorageConfig{ChunkSize: LocalStorageSize},
		LocalFiles: make(map[string]map[ChunkID]*Chunk),
	}
}

type FileHashTmp int

func (f FileHashTmp) String() string {
	return fmt.Sprintf("%d", f)
}

var tmp FileHashTmp = 1

func (s *P2PStorage) Dump() {
	s.Lock()
	for _, file := range s.LocalFiles {
		for _, chunk := range file {
			log.Println(chunk)
		}
	}
	s.Unlock()
}

// AddFile Add a file to local storage. Return the hashed file when successfull
func (s *P2PStorage) AddFile(fileData []byte) (FileID, error) {
	//hash := NewHashFromFile(fileData)
	hash := tmp
	key := hash.String()
	s.Lock()
	if _, ok := s.LocalFiles[key]; ok {
		s.Unlock()
		return nil, errors.New("Trying to add existing file")
	}
	file := FileDataToChunks(fileData, s.Config.ChunkSize)
	s.LocalFiles[key] = make(map[ChunkID]*Chunk, len(file))
	for i, chunk := range file {
		//log.Printf("Adding Chunk whith ID: %v\nFile: %v\nBytes: %v\n", chunk.Id, hash, chunk.B)
		s.LocalFiles[key][chunk.ID()] = &file[i]
	}
	tmp += 1
	//s.LocalFiles[key] = FileDataToChunks(fileData, s.Config.ChunkSize)
	s.Unlock()
	//s.Dump()
	return hash, nil
}

// AddFile Add a file to local storage. Return the hashed file when successfull
func (s *P2PStorage) AddReceivedFileChunks(hash FileID, chunks []Chunk) error {
	s.Lock()
	key := hash.String()
	if _, ok := s.LocalFiles[key]; !ok {
		s.LocalFiles[key] = make(map[ChunkID]*Chunk)
	}
	for i, chunk := range chunks {
		//log.Printf("Addind Chunk nb: %v | Data: %s\n", chunk.Id, string(chunk.B))
		// TODO: Throw error if Chunk already in storage ?
		s.LocalFiles[key][chunk.Id] = &chunks[i]
		//log.Printf("Added Chunk: %v \n", *s.LocalFiles[key][chunk.Id])
	}
	//for key, chunk := range chunks {
	//	fmt.Printf("in storage: %v with value:%s\n", key, string(chunk.B))
	//}
	s.Unlock()
	//s.LocalFiles[key] = FileDataToChunks(fileData, s.Config.ChunkSize)
	return nil
}

// GetChunks Search for requested chunk with file hash
func (s *P2PStorage) GetChunks(hash FileID, start, end ChunkID) ([]Chunk, error) {
	s.Lock()
	file, ok := s.LocalFiles[hash.String()]
	if !ok {
		s.Unlock()
		//return nil, fmt.Errorf("Requested file is not in storage: {%x}", hash)
		return nil, fmt.Errorf("Requested file is not in storage: {%s}", hash.String())
	}
	if start > end {
		s.Unlock()
		return nil, fmt.Errorf("Requested Chuncks out of range: [ start: %v, end: %v ]", start, end)
	}
	data := make([]Chunk, end-start+1)
	for i := 0; start <= end; i, start = i+1, start+1 {
		log.Printf("Searching for block with id: %v", start)
		chunk, ok := file[start]
		if !ok {
			s.Unlock()
			return nil, fmt.Errorf("Requested Chuncks not in local storage: [ start: %v, end: %v ]", start, end)
		}
		data[i] = *chunk
		copy(data[i].B, chunk.B)
	}
	s.Unlock()
	return data, nil
}

// GetChunks Search for requested chunk with file hash
func (s *P2PStorage) GetFileDatas(hash FileID) ([]byte, error) {
	var dataLen int

	s.Lock()
	chunks, ok := s.LocalFiles[hash.String()]
	if !ok {
		s.Unlock()
		return nil, errors.New("File not in storage")
	}
	for i := 0; ; i += 1 {
		chunk, ok1 := chunks[ChunkID(i)]
		if !ok1 {
			break
		}
		if chunkLen := len(chunk.B); chunkLen == 0 {
			break
		} else {
			dataLen += chunkLen
		}
	}
	data := make([]byte, dataLen)
	dataLen = 0
	for i := 0; ; i += 1 {
		chunk, ok1 := chunks[ChunkID(i)]
		if !ok1 {
			break
		}
		if chunkLen := len(chunk.B); chunkLen == 0 {
			break
		} else {
			copy(data[dataLen:], chunk.B)
			dataLen += chunkLen
		}
	}
	s.Unlock()
	return data, nil
}

//// GetChunks Search for requested chunk with file hash
//func (s *P2PStorage) GetChunks(hash FileHash, start, end ChunkID) ([]Chunk, error) {
//	data, ok := s.LocalFiles[string(hash)]
//	if ok {
//		return nil, fmt.Errorf("Requested file is not in storage: {%x}", hash)
//	}
//	dataLen := uint32(len(data))
//	if dataLen < uint32(end) || start > end {
//		return nil, fmt.Errorf("Requested Chuncks out of range: [ len: %d, start: %v, end: %v", dataLen, start, end)
//	}
//	return data[start:end], nil
//	//dataLen := uint32(len(data))
//	//if dataLen < uint32(end) || start > end {
//	//	return nil, fmt.Errorf("Requested Chuncks out of range: [ len: %d, start: %v, end: %v", dataLen, start, end)
//	//}
//	//return data[start:end], nil
//}

// GetDataFromLocalChunks Search for requested chunks with file hash and aggregate them in bytes
func (s *P2PStorage) GetDataFromLocalChunks(hash FileHash, start, end ChunkID) ([]byte, error) {
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

func NewHashFromFile(data []byte) FileHash {
	hash := sha256.Sum256(data)
	return hash[:]
}
