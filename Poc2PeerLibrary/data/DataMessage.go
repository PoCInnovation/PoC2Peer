package data

type DataMsg struct {
	Start ChunkID
	End   ChunkID
	Data  []byte
}
