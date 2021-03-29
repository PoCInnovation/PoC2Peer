package storage

import "errors"

// PeerID identifies a peer
type PeerID interface {
	String() string
}

type PeerStorage interface {
	UpdateFileInfos(peer PeerID, hash FileID, fileSize int, chunkIDS []ChunkID) error
	PeersHasFileChunks(hash FileID) (map[PeerID][]ChunkID, error)
	FileSize(hash FileID) (int, error)
}

type P2PRemoteStorage map[string]P2PFileStorage

type P2PFileStorage struct {
	Haves    map[PeerID][]ChunkID
	FileSize int
}

func NewP2PRemoteStorage() *P2PRemoteStorage {
	ret := make(P2PRemoteStorage)
	return &ret
}

func (s P2PRemoteStorage) UpdateFileInfos(peer PeerID, hash FileID, fileSize int, chunkIDS []ChunkID) error {
	peerStorage, ok := s[hash.String()]
	if !ok {
		peerStorage = P2PFileStorage{
			Haves:    make(map[PeerID][]ChunkID),
			FileSize: -1,
		}
	}
	fileChunks, ok := peerStorage.Haves[peer]
	if !ok {
		peerStorage.Haves[peer] = chunkIDS
	} else {
		peerStorage.Haves[peer] = removeDuplicates(append(fileChunks, chunkIDS...))
	}
	if fileSize != -1 && peerStorage.FileSize == -1 {
		peerStorage.FileSize = fileSize
	}
	s[hash.String()] = peerStorage
	return nil
}

func (s P2PRemoteStorage) PeersHasFileChunks(hash FileID) (map[PeerID][]ChunkID, error) {
	file, ok := s[hash.String()]
	if !ok {
		return nil, errors.New("File not remote Peer Storage")
	}
	// Create the target map
	targetMap := make(map[PeerID][]ChunkID, len(file.Haves))

	// Copy from the original map to the target map
	for key, value := range file.Haves {
		newValue := make([]ChunkID, len(value))
		copy(newValue, value)
		targetMap[key] = newValue
	}
	return targetMap, nil
}

func (s P2PRemoteStorage) FileSize(hash FileID) (int, error) {
	file, ok := s[hash.String()]
	if !ok {
		return -1, errors.New("File not remote Peer Storage")
	}
	return file.FileSize, nil
}

func removeDuplicates(s []ChunkID) []ChunkID {
	// Use of empty struct to optimize memory instead of boolean
	// https://dave.cheney.net/2014/03/25/the-empty-struct
	seen := make(map[ChunkID]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
