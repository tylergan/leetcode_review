package linked_list

import (
	"math/rand"
	"testing"
)

func TestMyHashSetExampleSequence(t *testing.T) {
	set := NewMyHashSet()

	set.Add(1)
	set.Add(2)
	if !set.Contains(1) {
		t.Fatal("Contains(1) = false, want true")
	}
	if set.Contains(3) {
		t.Fatal("Contains(3) = true, want false")
	}
	set.Add(2) // duplicate add is a no-op
	if !set.Contains(2) {
		t.Fatal("Contains(2) = false, want true")
	}
	set.Remove(2)
	if set.Contains(2) {
		t.Fatal("Contains(2) = true, want false after remove")
	}
}

func TestMyHashSetRemoveMissingIsNoOp(t *testing.T) {
	set := NewMyHashSet()

	set.Add(42)
	set.Remove(7) // never added; must not affect existing keys
	if !set.Contains(42) {
		t.Fatal("Contains(42) = false, want true after removing an absent key")
	}
	set.Remove(42)
	set.Remove(42) // removing twice must stay a no-op
	if set.Contains(42) {
		t.Fatal("Contains(42) = true, want false")
	}
}

func TestMyHashSetReAddAfterRemove(t *testing.T) {
	set := NewMyHashSet()

	set.Add(5)
	set.Remove(5)
	set.Add(5)
	if !set.Contains(5) {
		t.Fatal("Contains(5) = false, want true after re-adding")
	}
}

// Keys that share a bucket exercise the per-bucket skip list, which is the
// whole point of this implementation. With size = 10000, k and k+10000 collide.
func TestMyHashSetBucketCollisions(t *testing.T) {
	set := NewMyHashSet()

	colliding := []int{5, 10005, 20005, 30005}
	for _, k := range colliding {
		set.Add(k)
	}
	for _, k := range colliding {
		if !set.Contains(k) {
			t.Fatalf("Contains(%d) = false, want true (same bucket as %v)", k, colliding)
		}
	}

	if set.Contains(40005) { // same bucket, never added
		t.Fatal("Contains(40005) = true, want false")
	}

	set.Remove(20005) // remove a middle element of the chain
	if set.Contains(20005) {
		t.Fatal("Contains(20005) = true, want false after remove")
	}
	for _, k := range []int{5, 10005, 30005} {
		if !set.Contains(k) {
			t.Fatalf("Contains(%d) = false after removing a bucket sibling, want true", k)
		}
	}
}

func TestMyHashSetBoundaryKeys(t *testing.T) {
	set := NewMyHashSet()

	for _, k := range []int{0, 1000000} {
		set.Add(k)
		if !set.Contains(k) {
			t.Fatalf("Contains(%d) = false, want true", k)
		}
	}
}

// Cross-check every operation against Go's built-in map to catch any drift in
// the skip-list bucket logic across a long randomized op stream.
func TestMyHashSetMatchesReferenceMap(t *testing.T) {
	set := NewMyHashSet()
	ref := map[int]bool{}

	rng := rand.New(rand.NewSource(99))
	const ops = 20000
	const keySpace = 3000 // small space => frequent collisions and re-adds

	for i := 0; i < ops; i++ {
		key := rng.Intn(keySpace)
		switch rng.Intn(3) {
		case 0:
			set.Add(key)
			ref[key] = true
		case 1:
			set.Remove(key)
			delete(ref, key)
		default:
			if got, want := set.Contains(key), ref[key]; got != want {
				t.Fatalf("op %d: Contains(%d) = %v, want %v", i, key, got, want)
			}
		}
	}

	for key := 0; key < keySpace; key++ {
		if got, want := set.Contains(key), ref[key]; got != want {
			t.Fatalf("final: Contains(%d) = %v, want %v", key, got, want)
		}
	}
}
