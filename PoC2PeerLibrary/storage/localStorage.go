package storage

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type LocalStorage interface {
	AddFile(fileData []byte) (FileID, error)
	AddFileChunks(hash FileID, chunks []Chunk) error
	GetRequestedChunks(hash FileID, ids []ChunkID) ([]Chunk, error)
	FilesList() []FileHash
	FileInfos(hash FileID) (ids []ChunkID, fileSize int, err error)
	FileData(hash FileID) (data []byte, err error)
	DeleteFileData(hash FileID) error
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

// AddFile Add a file to local storage. Return the hashed file when successfull
func (s *P2PStorage) AddFile(fileData []byte) (FileID, error) {
	hash := NewHashFromFile(fileData)
	key := hash.String()
	s.Lock()
	if _, ok := s.LocalFiles[key]; ok {
		s.Unlock()
		return nil, errors.New("Trying to add existing file")
	}
	s.LocalFiles[key] = NewFile(hash, FSComplete, fileData, s.Config.ChunkSize)
	s.Unlock()
	return hash, nil
}

// AddFile Add a file to local storage. Return the hashed file when successfull
func (s *P2PStorage) AddFileChunks(hash FileID, chunks []Chunk) error {
	if len(chunks) < 1 {
		// TODO: Return Err ?
		log.Printf("AddFileChunks received an empty chunks List for File %x, nothing to do ...")
		return nil
	}
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
	log.Printf("Adding File %x whith chunks from %v to %v", hash, chunks[0].Id, chunks[len(chunks)-1].Id)
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

func (s *P2PStorage) FilesList() []FileHash {
	s.Lock()
	lst := make([]FileHash, len(s.LocalFiles))
	i := 0
	for _, file := range s.LocalFiles {
		lst[i] = file.Hash
		i += 1
	}
	s.Unlock()
	return lst
}

func (s *P2PStorage) FileInfos(hash FileID) ([]ChunkID, int, error) {
	s.Lock()
	file, ok := s.LocalFiles[hash.String()]
	if !ok {
		s.Unlock()
		return nil, -1, fmt.Errorf("Requested file is not in storage: {%s}", hash.String())
	}
	sz := file.Size()
	chunks := file.GetChunksIDs()
	s.Unlock()
	return chunks, sz, nil
}

var FILENOTFOUND = errors.New("P2PFile not in storage")

// GetChunkIDsInStorage Search for requested chunk with file Hash
func (s *P2PStorage) FileData(hash FileID) ([]byte, error) {
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

func (s *P2PStorage) DeleteFileData(hash FileID) error {
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
