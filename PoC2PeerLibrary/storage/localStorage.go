package storage

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

type LocalStorage interface {
	AddFile(fileData []byte) (FileID, error)
	AddReceivedFileChunks(hash FileID, chunks []Chunk) error
	GetRequestedChunks(hash FileID, ids []ChunkID) ([]Chunk, error)
	GetChunkIDsInStorage(hash FileID) ([]ChunkID, error)
	GetFileData(hash FileID) ([]byte, error)
	DeleteData(hash FileID) error
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
	s.LocalFiles[key] = NewFile(hash, FSComplete, fileData, s.Config.ChunkSize)
	//file := FileDataToChunks(fileData, s.Config.ChunkSize)
	//s.LocalFiles[key] = P2PFile{state: FSComplete, Data: fileData, Chunks: make(map[ChunkID]Chunk, len(file))}
	//for i, chunk := range file {
	//	//log.Printf("Adding Chunk whith ID: %v\nFile: %v\nBytes: %v\n", chunk.Id, hash, chunk.B)
	//	s.LocalFiles[key].Chunks[chunk.ID()] = file[i]
	//}
	//s.LocalFiles[key] = FileDataToChunks(fileData, s.Config.ChunkSize)
	s.Unlock()
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
		s.LocalFiles[key] = NewFile(v, FSUpdated, []byte{}, s.Config.ChunkSize)
	}
	file := s.LocalFiles[key]
	file.AddChunks(chunks)
	file.UpdateData()
	s.LocalFiles[key] = file
	s.Unlock()
	return nil
}

// GetChunkIDsInStorage Search for requested chunk with file Hash
func (s *P2PStorage) GetRequestedChunks(hash FileID, ids []ChunkID) ([]Chunk, error) {
	s.Lock()
	file, ok := s.LocalFiles[hash.String()]
	if !ok {
		s.Unlock()
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
		return nil, fmt.Errorf("Requested file is not in storage: {%s}", hash.String())
	}
	chunks := file.GetChunksIDs()
	s.Unlock()
	sort.Sort(ChunkIDs(chunks))
	return chunks, nil
}

// TODO: modify Sort
type ChunkIDs []ChunkID

func (a ChunkIDs) Len() int           { return len(a) }
func (a ChunkIDs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ChunkIDs) Less(i, j int) bool { return a[i] < a[j] }

var FILENOTFOUND = errors.New("P2PFile not in storage")

// GetChunkIDsInStorage Search for requested chunk with file Hash
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

func (s *P2PStorage) DeleteData(hash FileID) error {
	s.Lock()
	file, ok := s.LocalFiles[hash.String()]
	if !ok {
		s.Unlock()
		return FILENOTFOUND
	}
	file.DeleteData()
	s.Unlock()
	return nil
}
