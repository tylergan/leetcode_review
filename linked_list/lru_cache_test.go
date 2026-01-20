package linked_list

import "testing"

type lruOp struct {
	name  string
	key   int
	value int
}

func runLRUOps(t *testing.T, capacity int, ops []lruOp) []int {
	t.Helper()

	cache := Constructor(capacity)
	var got []int

	for _, op := range ops {
		switch op.name {
		case "put":
			cache.Put(op.key, op.value)
		case "get":
			got = append(got, cache.Get(op.key))
		default:
			t.Fatalf("unknown op %q", op.name)
		}
	}

	return got
}

func TestLRUCacheExample(t *testing.T) {
	ops := []lruOp{
		{name: "put", key: 1, value: 10},
		{name: "get", key: 1},
		{name: "put", key: 2, value: 20},
		{name: "put", key: 3, value: 30},
		{name: "get", key: 2},
		{name: "get", key: 1},
	}
	want := []int{10, 20, -1}

	got := runLRUOps(t, 2, ops)
	if len(got) != len(want) {
		t.Fatalf("example outputs len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("example outputs[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestLRUCacheCapacityOne(t *testing.T) {
	ops := []lruOp{
		{name: "put", key: 1, value: 1},
		{name: "get", key: 1},
		{name: "put", key: 2, value: 2},
		{name: "get", key: 1},
		{name: "get", key: 2},
	}
	want := []int{1, -1, 2}

	got := runLRUOps(t, 1, ops)
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("capacity=1 outputs[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestLRUCacheGetUpdatesRecency(t *testing.T) {
	ops := []lruOp{
		{name: "put", key: 1, value: 1},
		{name: "put", key: 2, value: 2},
		{name: "get", key: 1},
		{name: "put", key: 3, value: 3},
		{name: "get", key: 2},
		{name: "get", key: 1},
		{name: "get", key: 3},
	}
	want := []int{1, -1, 1, 3}

	got := runLRUOps(t, 2, ops)
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("get-recency outputs[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestLRUCachePutUpdatesValue(t *testing.T) {
	ops := []lruOp{
		{name: "put", key: 1, value: 1},
		{name: "put", key: 2, value: 2},
		{name: "put", key: 1, value: 10},
		{name: "put", key: 3, value: 3},
		{name: "get", key: 1},
		{name: "get", key: 2},
		{name: "get", key: 3},
	}
	want := []int{10, -1, 3}

	got := runLRUOps(t, 2, ops)
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("put-update outputs[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}
