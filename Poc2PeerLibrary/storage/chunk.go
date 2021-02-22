package storage

// ChunkID identifies a chunk of content
type ChunkID uint32

// Chunk represents a chunk of content
type Chunk struct {
	ID ChunkID
	B  []byte
}

func newChunk(id ChunkID, size int) *Chunk {
	var c Chunk
	c.B = make([]byte, size)
	c.ID = id
	return &c
}

func newChunkFromData(id ChunkID, size int, data []byte) *Chunk {
	chunk := newChunk(id, size)
	copy(chunk.B, data)
	return chunk
}

func FileDataToChunks(fileData []byte, chunkSize int) []*Chunk {
	var chunks []*Chunk
	for i := 0; ; i += 1 {
		if len(fileData) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(fileData) < chunkSize {
			chunkSize = len(fileData)
		}

		chunks = append(chunks, newChunkFromData(ChunkID(i), chunkSize, fileData[0:chunkSize]))
		fileData = fileData[chunkSize:]
	}

	return chunks
}
