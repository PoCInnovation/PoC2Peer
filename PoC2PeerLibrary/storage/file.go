package storage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

type FileHash []byte

func (h FileHash) String() string {
	return string(h)
}

func (h FileHash) Eq(lhs FileHash) bool {
	return bytes.Compare(h, lhs) == 0
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

type FileState int

const (
	Complete  FileState = 0
	Updated   FileState = 1
	Unchanged FileState = 2
)

type P2PFile struct {
	hash   FileHash
	state  FileState
	Data   []byte
	Chunks map[ChunkID]Chunk
}

func NewFile(hash FileHash, state FileState, fileData []byte, chunkSize int) P2PFile {
	chunks := FileDataToChunks(fileData, chunkSize)
	newFile := P2PFile{hash: hash, state: state, Data: fileData, Chunks: make(map[ChunkID]Chunk, len(chunks))}
	for i, chunk := range chunks {
		log.Printf("Adding Chunk whith ID: %v\nFile: %v\nBytes: %v\n", chunk.Id, hash, chunk.B)
		newFile.Chunks[chunk.ID()] = chunks[i]
	}
	return newFile
}

func (f P2PFile) GetRequestedChunks(ids []ChunkID) []Chunk {
	n := 0
	for _, id := range ids {
		if _, ok := f.Chunks[id]; !ok {
			log.Printf("search for requested chunk {%v} but not found in storage\n", id)
		} else {
			n += 1
		}
	}
	chunks := make([]Chunk, n)
	n = 0
	for _, id := range ids {
		if val, ok := f.Chunks[id]; ok {
			chunks[n] = val
			n += 1
		}
	}
	return chunks
}

func (f *P2PFile) AddChunks(chunks []Chunk) error {
	for i, chunk := range chunks {
		// TODO: keep ??
		if _, ok := f.Chunks[chunk.Id]; ok {
			log.Println("Adding Chunk but was already there")
		}
		f.Chunks[chunk.Id] = chunks[i]
		f.state = Updated
	}
	return nil
}

func (f P2PFile) Complete() bool {
	return f.state == Complete
}

func (f P2PFile) GetChunksIDs() []ChunkID {
	ids := make([]ChunkID, len(f.Chunks))
	i := 0
	for id, _ := range f.Chunks {
		ids[i] = id
	}
	return ids
}

func (f *P2PFile) GetData() []byte {
	if f.Complete() || f.state == Unchanged {
		return f.Data
	}
	f.UpdateData()
	return f.Data
}

func (f *P2PFile) UpdateData() {
	dataLen := 0
	for i := 0; ; i += 1 {
		chunk, ok1 := f.Chunks[ChunkID(i)]
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
		chunk, ok1 := f.Chunks[ChunkID(i)]
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
	f.Data = data
	if f.hash.Eq(NewHashFromFile(data)) {
		f.state = Complete
	}
}

func FileDataToChunks(fileData []byte, chunkSize int) (chunks []Chunk) {
	for i := 0; ; i += 1 {
		if len(fileData) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(fileData) < chunkSize {
			chunkSize = len(fileData)
		}

		chunks = append(chunks, NewChunkFromData(ChunkID(i), chunkSize, fileData[0:chunkSize]))
		fileData = fileData[chunkSize:]
	}
	return
}

func NewHashFromFile(data []byte) FileHash {
	hash := sha256.Sum256(data)
	return hash[:]
}
