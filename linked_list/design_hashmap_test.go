package linked_list

import (
	"math/rand"
	"testing"
)

func TestMyHashMapExampleSequence(t *testing.T) {
	m := NewMyHashMap()

	m.Put(1, 1)
	m.Put(2, 2)
	if got := m.Get(1); got != 1 {
		t.Fatalf("Get(1) = %d, want 1", got)
	}
	if got := m.Get(3); got != -1 {
		t.Fatalf("Get(3) = %d, want -1", got)
	}
	m.Put(2, 1) // update existing key
	if got := m.Get(2); got != 1 {
		t.Fatalf("Get(2) = %d, want 1 after update", got)
	}
	m.Remove(2)
	if got := m.Get(2); got != -1 {
		t.Fatalf("Get(2) = %d, want -1 after remove", got)
	}
}

func TestMyHashMapUpdateDoesNotDuplicate(t *testing.T) {
	m := NewMyHashMap()

	m.Put(10, 100)
	m.Put(10, 200)
	m.Put(10, 300)
	if got := m.Get(10); got != 300 {
		t.Fatalf("Get(10) = %d, want 300 after repeated updates", got)
	}
	m.Remove(10) // a single remove must fully clear the key, proving no dupes
	if got := m.Get(10); got != -1 {
		t.Fatalf("Get(10) = %d, want -1 after remove", got)
	}
}

func TestMyHashMapRemoveMissingIsNoOp(t *testing.T) {
	m := NewMyHashMap()

	m.Put(42, 7)
	m.Remove(99) // never inserted
	if got := m.Get(42); got != 7 {
		t.Fatalf("Get(42) = %d, want 7 after removing an absent key", got)
	}
}

// value 0 is a legal stored value; only a missing key yields -1.
func TestMyHashMapZeroValueIsDistinctFromMissing(t *testing.T) {
	m := NewMyHashMap()

	m.Put(5, 0)
	if got := m.Get(5); got != 0 {
		t.Fatalf("Get(5) = %d, want 0 (stored value, not missing)", got)
	}
	if got := m.Get(6); got != -1 {
		t.Fatalf("Get(6) = %d, want -1 (missing)", got)
	}
}

// Keys k and k+10000 share a bucket (size = 10000), exercising the per-bucket
// skip list rather than just distinct slots.
func TestMyHashMapBucketCollisions(t *testing.T) {
	m := NewMyHashMap()

	pairs := map[int]int{5: 50, 10005: 100, 20005: 200, 30005: 300}
	for k, v := range pairs {
		m.Put(k, v)
	}
	for k, v := range pairs {
		if got := m.Get(k); got != v {
			t.Fatalf("Get(%d) = %d, want %d (colliding bucket)", k, got, v)
		}
	}

	m.Put(10005, 999) // update a middle chain element
	if got := m.Get(10005); got != 999 {
		t.Fatalf("Get(10005) = %d, want 999 after update", got)
	}

	m.Remove(20005)
	if got := m.Get(20005); got != -1 {
		t.Fatalf("Get(20005) = %d, want -1 after remove", got)
	}
	for k, want := range map[int]int{5: 50, 10005: 999, 30005: 300} {
		if got := m.Get(k); got != want {
			t.Fatalf("Get(%d) = %d, want %d after removing a bucket sibling", k, got, want)
		}
	}
}

func TestMyHashMapBoundaryKeys(t *testing.T) {
	m := NewMyHashMap()

	m.Put(0, 11)
	m.Put(1000000, 22)
	if got := m.Get(0); got != 11 {
		t.Fatalf("Get(0) = %d, want 11", got)
	}
	if got := m.Get(1000000); got != 22 {
		t.Fatalf("Get(1000000) = %d, want 22", got)
	}
}

// Cross-check the full op stream against Go's built-in map.
func TestMyHashMapMatchesReferenceMap(t *testing.T) {
	m := NewMyHashMap()
	ref := map[int]int{}

	rng := rand.New(rand.NewSource(7))
	const ops = 20000
	const keySpace = 3000 // small space => frequent collisions and updates

	get := func(k int) int {
		if v, ok := ref[k]; ok {
			return v
		}
		return -1
	}

	for i := 0; i < ops; i++ {
		key := rng.Intn(keySpace)
		switch rng.Intn(3) {
		case 0:
			val := rng.Intn(1_000_001)
			m.Put(key, val)
			ref[key] = val
		case 1:
			m.Remove(key)
			delete(ref, key)
		default:
			if got, want := m.Get(key), get(key); got != want {
				t.Fatalf("op %d: Get(%d) = %d, want %d", i, key, got, want)
			}
		}
	}

	for key := 0; key < keySpace; key++ {
		if got, want := m.Get(key), get(key); got != want {
			t.Fatalf("final: Get(%d) = %d, want %d", key, got, want)
		}
	}
}
