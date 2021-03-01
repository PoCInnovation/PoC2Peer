package storage

// PeerID identifies a peer
type PeerID interface {
	String() string
}

type PeerStorage interface {
	AddPeerFileChunks(peer PeerID, hash FileID, chunkIDS []ChunkID) error
}

//type P2PRemoteStorage map[PeerID]P2PFileStorage
//
//type P2PFileStorage map[FileID][]ChunkID
//
//func (s P2PRemoteStorage) AddPeerFileChunks(peer PeerID, hash FileID, chunkIDS []ChunkID) error {
//	peerStorage, ok := s[peer]
//	if !ok {
//		peerStorage = make(P2PFileStorage)
//	}
//	fileChunks, ok := peerStorage[hash]
//	if !ok {
//		peerStorage[hash] = chunkIDS
//	} else {
//		peerStorage[hash] = removeDuplicates(append(fileChunks, chunkIDS...))
//	}
//	s[peer] = peerStorage
//	return nil
//}

type P2PRemoteStorage map[string]P2PFileStorage

type P2PFileStorage map[PeerID][]ChunkID

func NewP2PRemoteStorage() *P2PRemoteStorage {
	ret := make(P2PRemoteStorage)
	return &ret
}

func (s P2PRemoteStorage) AddPeerFileChunks(peer PeerID, hash FileID, chunkIDS []ChunkID) error {
	peerStorage, ok := s[hash.String()]
	if !ok {
		peerStorage = make(P2PFileStorage)
	}
	fileChunks, ok := peerStorage[peer]
	if !ok {
		peerStorage[peer] = chunkIDS
	} else {
		peerStorage[peer] = removeDuplicates(append(fileChunks, chunkIDS...))
	}
	s[hash.String()] = peerStorage
	return nil
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
