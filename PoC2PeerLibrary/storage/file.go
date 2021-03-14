package storage

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
)

// FileHash:
type FileHash []byte

// String: Return string representation from FileHash
func (h FileHash) String() string {
	return string(h)
	//return fmt.Sprintf("%x", h)
}

// Eq: Compare two FileHash
func (h FileHash) Eq(lhs FileHash) bool {
	return bytes.Compare(h, lhs) == 0
}

// TODO
func (h FileHash) Decode() FileHash {
	//return []byte(h)
	//he, err := hex.DecodeString(h.String())
	//if err != nil {
	//	log.Printf("decoding filehash failed")
	//	return []byte("")
	//}
	//fmt.Println(h)
	//return FileHash(he)
	return []byte(fmt.Sprintf("%x", h))
}

// FileState: State of a file, representing if it's complete, updated or unchanged
type FileState int

// TODO
const (
	FSComplete  FileState = 0
	FSUpdated   FileState = 1
	FSUnchanged FileState = 2
)

// P2PFile: Represent a P2PFile with all basic file infos.
//  - Hash: unique ID for the file based on file content
//  - state: FileState representing if the file is complete, updated or unchanged.
//  - Data: File content formed with chunks' data. Used to avoid excess data merging when requesting data.
//  - Chunks: All stored Chunks accesible by ChunkID.
type P2PFile struct {
	Hash   FileHash
	State  FileState
	Data   []byte
	Chunks map[ChunkID]Chunk
}

// TODO
func NewFile(hash FileHash, state FileState, fileData []byte, chunkSize int) P2PFile {
	chunks := FileDataToChunks(fileData, chunkSize)
	newFile := P2PFile{
		Hash:   hash,
		State:  state,
		Data:   fileData,
		Chunks: make(map[ChunkID]Chunk, len(chunks)),
	}
	for i, chunk := range chunks {
		//log.Printf("Adding Chunk whith ID: %v\nFile: %v\nBytes: %v\n", chunk.Id, hash, chunk.B)
		newFile.Chunks[chunk.ID()] = chunks[i]
	}
	if len(chunks) > 0 {
		log.Printf("Adding File %x whith chunks from %v to %v", newFile.Hash, chunks[0], chunks[len(chunks)-1])
	} else {
		log.Printf("Adding File %x whith no chunks", newFile.Hash)
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

// TODO
func (f *P2PFile) AddChunks(chunks []Chunk) error {
	for i, chunk := range chunks {
		// TODO: keep ??
		if _, ok := f.Chunks[chunk.Id]; ok {
			log.Println("Adding Chunk but was already there")
		}
		f.Chunks[chunk.Id] = chunks[i]
		f.State = FSUpdated
	}
	return nil
}

// Complete: Check P2PFile state, return true if file is completed.
func (f P2PFile) Complete() bool {
	return f.State == FSComplete
}

// GetChunksIDs: Return all ChunkID's in a P2PFile.
func (f P2PFile) GetChunksIDs() []ChunkID {
	ids := make([]ChunkID, len(f.Chunks))
	i := 0
	for id, _ := range f.Chunks {
		ids[i] = id
		i += 1
	}
	return ids
}

// GetData: Return stored data in a P2PFile.
// if the file is completed or unchanged since last call to GetData, directly return stored data.
// otherwise (FileState = FSUpdated), call UpdateData to reform data from received Chunk.
func (f *P2PFile) GetData() []byte {
	if f.Complete() || f.State == FSUnchanged {
		return f.Data
	}
	f.UpdateData()
	return f.Data
}

// TODO:
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
	if f.Hash.Eq(NewHashFromFile(data)) {
		f.State = FSComplete
	}
}

//DeleteData: Delete file Data to free ressources
func (f *P2PFile) DeleteData() {
	f.Data = []byte{}
}

// FileDataToChunks: Transform array of bytes (content of a file) into Chunk's array with 'chunkSize' size.
func FileDataToChunks(fileData []byte, chunkSize int) (chunks []Chunk) {
	for i := 0; ; i += 1 {
		// Break if we are at the end of Data
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

// NewHashFromFile: Create a New Hashfrom file's content with SHA256
func NewHashFromFile(data []byte) FileHash {
	hash := sha256.Sum256(data)
	return hash[:]
}
