package binary_search

import "testing"

func TestTimeMapExample(t *testing.T) {
	tm := Constructor()

	tm.Set("alice", "happy", 1)
	if got := tm.Get("alice", 1); got != "happy" {
		t.Fatalf("get(alice, 1) = %q, want %q", got, "happy")
	}
	if got := tm.Get("alice", 2); got != "happy" {
		t.Fatalf("get(alice, 2) = %q, want %q", got, "happy")
	}

	tm.Set("alice", "sad", 3)
	if got := tm.Get("alice", 3); got != "sad" {
		t.Fatalf("get(alice, 3) = %q, want %q", got, "sad")
	}
}

func TestTimeMapUnknownKey(t *testing.T) {
	tm := Constructor()
	if got := tm.Get("missing", 10); got != "" {
		t.Fatalf("get(missing, 10) = %q, want empty string", got)
	}
}

func TestTimeMapExactTimestamps(t *testing.T) {
	tm := Constructor()
	tm.Set("k", "v1", 1)
	tm.Set("k", "v2", 5)
	tm.Set("k", "v3", 10)

	if got := tm.Get("k", 1); got != "v1" {
		t.Fatalf("get(k, 1) = %q, want %q", got, "v1")
	}
	if got := tm.Get("k", 5); got != "v2" {
		t.Fatalf("get(k, 5) = %q, want %q", got, "v2")
	}
	if got := tm.Get("k", 10); got != "v3" {
		t.Fatalf("get(k, 10) = %q, want %q", got, "v3")
	}
}

func TestTimeMapBetweenTimestamps(t *testing.T) {
	tm := Constructor()
	tm.Set("k", "v1", 2)
	tm.Set("k", "v2", 6)
	tm.Set("k", "v3", 9)

	if got := tm.Get("k", 1); got != "" {
		t.Fatalf("get(k, 1) = %q, want empty string", got)
	}
	if got := tm.Get("k", 4); got != "v1" {
		t.Fatalf("get(k, 4) = %q, want %q", got, "v1")
	}
	if got := tm.Get("k", 8); got != "v2" {
		t.Fatalf("get(k, 8) = %q, want %q", got, "v2")
	}
	if got := tm.Get("k", 100); got != "v3" {
		t.Fatalf("get(k, 100) = %q, want %q", got, "v3")
	}
}

func TestTimeMapMultipleKeys(t *testing.T) {
	tm := Constructor()
	tm.Set("a", "x", 1)
	tm.Set("b", "y", 2)
	tm.Set("a", "z", 3)

	if got := tm.Get("a", 2); got != "x" {
		t.Fatalf("get(a, 2) = %q, want %q", got, "x")
	}
	if got := tm.Get("b", 2); got != "y" {
		t.Fatalf("get(b, 2) = %q, want %q", got, "y")
	}
	if got := tm.Get("a", 3); got != "z" {
		t.Fatalf("get(a, 3) = %q, want %q", got, "z")
	}
}

func TestTimeMapSingleEntry(t *testing.T) {
	tm := Constructor()
	tm.Set("k", "v", 100)

	if got := tm.Get("k", 99); got != "" {
		t.Fatalf("get(k, 99) = %q, want empty string", got)
	}
	if got := tm.Get("k", 100); got != "v" {
		t.Fatalf("get(k, 100) = %q, want %q", got, "v")
	}
	if got := tm.Get("k", 101); got != "v" {
		t.Fatalf("get(k, 101) = %q, want %q", got, "v")
	}
}
