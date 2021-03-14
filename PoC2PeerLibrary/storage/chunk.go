package storage

// ChunkID identifies a chunk of content
type ChunkID uint32

// Chunk represents a chunk of content
type Chunk struct {
	Id ChunkID
	B  []byte
}

func (c *Chunk) ID() ChunkID {
	return c.Id
}

func newChunk(id ChunkID, size int) Chunk {
	var c Chunk
	c.B = make([]byte, size)
	c.Id = id
	return c
}

func NewChunkFromData(id ChunkID, size int, data []byte) Chunk {
	var c Chunk
	c.B = make([]byte, size)
	c.Id = id
	copy(c.B, data)
	return c
}
