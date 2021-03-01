package storage

import (
	"errors"
	"fmt"
	"sync"
)

type LocalStorage interface {
	AddFile(fileData []byte) (FileID, error)
	AddReceivedFileChunks(hash FileID, chunks []Chunk) error
	GetRequestedChunks(hash FileID, ids []ChunkID) ([]Chunk, error)
	GetChunkIDsInStorage(hash FileID) ([]ChunkID, error)
	GetFileData(hash FileID) ([]byte, error)
}

type FileID interface {
	String() string
}

const LocalStorageSize = 100

type LocalStorageConfig struct {
	ChunkSize int
}

type P2PStorage struct {
	// Map in not thread safe by design, so we have to use a mutex
	sync.Mutex
	Config     LocalStorageConfig
	LocalFiles map[string]P2PFile
}

func NewP2PStorage() LocalStorage {
	return &P2PStorage{
		Config:     LocalStorageConfig{ChunkSize: LocalStorageSize},
		LocalFiles: make(map[string]P2PFile),
	}
}

type FileHashTmp int

func (f FileHashTmp) String() string {
	return fmt.Sprintf("%d", f)
}

func (s *P2PStorage) Dump() {
	//s.Lock()
	//for _, file := range s.LocalFiles {
	//	for _, chunk := range file {
	//		log.Println(chunk)
	//	}
	//}
	//s.Unlock()
}

// AddFile Add a file to local storage. Return the hashed file when successfull
func (s *P2PStorage) AddFile(fileData []byte) (FileID, error) {
	hash := NewHashFromFile(fileData)
	//hash := tmp
	key := hash.String()
	s.Lock()
	if _, ok := s.LocalFiles[key]; ok {
		s.Unlock()
		return nil, errors.New("Trying to add existing file")
	}
	s.LocalFiles[key] = NewFile(hash, Complete, fileData, s.Config.ChunkSize)
	//file := FileDataToChunks(fileData, s.Config.ChunkSize)
	//s.LocalFiles[key] = P2PFile{state: Complete, Data: fileData, Chunks: make(map[ChunkID]Chunk, len(file))}
	//for i, chunk := range file {
	//	//log.Printf("Adding Chunk whith ID: %v\nFile: %v\nBytes: %v\n", chunk.Id, hash, chunk.B)
	//	s.LocalFiles[key].Chunks[chunk.ID()] = file[i]
	//}
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
		v, ok1 := hash.(FileHash)
		if !ok1 {
			return errors.New("FileID is not a file Hash")
		}
		s.LocalFiles[key] = NewFile(v, Updated, []byte{}, s.Config.ChunkSize)
	}
	file := s.LocalFiles[key]
	file.AddChunks(chunks)
	s.LocalFiles[key] = file
	//for i, chunk := range chunks {
	//	//log.Printf("Addind Chunk nb: %v | Data: %s\n", chunk.Id, string(chunk.B))
	//	// TODO: Throw error if Chunk already in storage ?
	//	s.LocalFiles[key].Chunks[chunk.Id] = chunks[i]
	//	//log.Printf("Added Chunk: %v \n", *s.LocalFiles[key][chunk.Id])
	//}
	//for key, chunk := range chunks {
	//	fmt.Printf("in storage: %v with value:%s\n", key, string(chunk.B))
	//}
	s.Unlock()
	//s.LocalFiles[key] = FileDataToChunks(fileData, s.Config.ChunkSize)
	return nil
}

// GetChunkIDsInStorage Search for requested chunk with file hash
func (s *P2PStorage) GetRequestedChunks(hash FileID, ids []ChunkID) ([]Chunk, error) {
	s.Lock()
	file, ok := s.LocalFiles[hash.String()]
	if !ok {
		s.Unlock()
		//return nil, fmt.Errorf("Requested file is not in storage: {%x}", hash)
		return nil, fmt.Errorf("Requested file is not in storage: {%s}", hash.String())
	}
	chunks := file.GetRequestedChunks(ids)
	s.Unlock()
	return chunks, nil
}

func (s *P2PStorage) GetChunkIDsInStorage(hash FileID) ([]ChunkID, error) {
	s.Lock()
	file, ok := s.LocalFiles[hash.String()]
	if !ok {
		s.Unlock()
		//return nil, fmt.Errorf("Requested file is not in storage: {%x}", hash)
		return nil, fmt.Errorf("Requested file is not in storage: {%s}", hash.String())
	}
	chunks := file.GetChunksIDs()
	s.Unlock()
	return chunks, nil
}

//// GetRequestedChunks Search for requested chunk with file hash
//func (s *P2PStorage) GetRequestedChunks(hash FileID, start, end ChunkID) ([]Chunk, error) {
//	s.Lock()
//	file, ok := s.LocalFiles[hash.String()]
//	if !ok {
//		s.Unlock()
//		//return nil, fmt.Errorf("Requested file is not in storage: {%x}", hash)
//		return nil, fmt.Errorf("Requested file is not in storage: {%s}", hash.String())
//	}
//	if start > end {
//		s.Unlock()
//		return nil, fmt.Errorf("Requested Chuncks out of range: [ start: %v, end: %v ]", start, end)
//	}
//	data := make([]Chunk, end-start+1)
//	for i := 0; start <= end; i, start = i+1, start+1 {
//		log.Printf("Searching for block with id: %v", start)
//		chunk, ok := file[start]
//		if !ok {
//			s.Unlock()
//			return nil, fmt.Errorf("Requested Chuncks not in local storage: [ start: %v, end: %v ]", start, end)
//		}
//		data[i] = chunk
//		copy(data[i].B, chunk.B)
//	}
//	s.Unlock()
//	return data, nil
//}

var FILENOTFOUND = errors.New("P2PFile not in storage")

// GetChunkIDsInStorage Search for requested chunk with file hash
func (s *P2PStorage) GetFileData(hash FileID) ([]byte, error) {
	s.Lock()
	file, ok := s.LocalFiles[hash.String()]
	if !ok {
		s.Unlock()
		return nil, FILENOTFOUND
	}
	data := file.GetData()
	s.Unlock()
	return data, nil
}

//// GetDataFromLocalChunks Search for requested chunks with file hash and aggregate them in bytes
//func (s *P2PStorage) GetDataFromLocalChunks(hash FileHash, ids []ChunkID) ([]byte, error) {
//	var dataLen int
//	chunks, err := s.GetRequestedChunks(hash, start, end)
//	if err != nil {
//		return nil, err
//	}
//	for _, chunk := range chunks {
//		dataLen += len(chunk.B)
//	}
//	data := make([]byte, dataLen)
//	dataLen = 0
//	for _, chunk := range chunks {
//		copy(data[dataLen:], chunk.B)
//		dataLen += len(chunk.B)
//	}
//	return data, nil
//}
