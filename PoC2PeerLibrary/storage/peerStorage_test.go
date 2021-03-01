package storage

import (
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	inputs := [][]ChunkID{
		[]ChunkID{0, 1, 2, 3, 2},
		[]ChunkID{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]ChunkID{0, 1, 2, 3, 2, 0, 1, 2, 3, 2, 0, 1, 2, 3, 2, 0, 1, 2, 3, 2, 0, 1, 2, 3, 2, 0, 1, 2, 3, 2, 0, 1, 2, 3, 2},
		[]ChunkID{0, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	for _, slice := range inputs {
		result := removeDuplicates(slice)
		seen := make(map[ChunkID]struct{})
		for _, id := range result {
			if _, ok := seen[id]; ok {
				t.Errorf("Id already seen: %v", id)
			} else {
				seen[id] = struct{}{}
			}
		}
		for _, id := range result {
			found := false
			for _, id1 := range result {
				if id == id1 {
					t.Logf("Id found : %v", id)
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Id Not found in final slice: %v", id)
			}
		}
	}
}
