package stack

import (
	"math"
	"reflect"
	"testing"
)

type minStackOp struct {
	name string
	arg  int
}

func runMinStackOps(t *testing.T, ops []minStackOp) []int {
	t.Helper()

	ms := Constructor()
	var got []int

	for _, op := range ops {
		switch op.name {
		case "push":
			ms.Push(op.arg)
		case "pop":
			ms.Pop()
		case "top":
			got = append(got, ms.Top())
		case "getMin":
			got = append(got, ms.GetMin())
		default:
			t.Fatalf("unknown op %q", op.name)
		}
	}

	return got
}

func TestMinStackExample(t *testing.T) {
	ops := []minStackOp{
		{name: "push", arg: 1},
		{name: "push", arg: 2},
		{name: "push", arg: 0},
		{name: "getMin"},
		{name: "pop"},
		{name: "top"},
		{name: "getMin"},
	}
	want := []int{0, 2, 1}

	got := runMinStackOps(t, ops)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("example outputs = %v, want %v", got, want)
	}
}

func TestMinStackNegativeValues(t *testing.T) {
	ops := []minStackOp{
		{name: "push", arg: -2},
		{name: "push", arg: 0},
		{name: "push", arg: -3},
		{name: "getMin"},
		{name: "pop"},
		{name: "top"},
		{name: "getMin"},
	}
	want := []int{-3, 0, -2}

	got := runMinStackOps(t, ops)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("negative outputs = %v, want %v", got, want)
	}
}

func TestMinStackWithDuplicates(t *testing.T) {
	ops := []minStackOp{
		{name: "push", arg: 3},
		{name: "push", arg: 5},
		{name: "getMin"},
		{name: "push", arg: 2},
		{name: "push", arg: 2},
		{name: "getMin"},
		{name: "pop"},
		{name: "getMin"},
		{name: "pop"},
		{name: "getMin"},
		{name: "top"},
	}
	want := []int{3, 2, 2, 3, 5}

	got := runMinStackOps(t, ops)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("duplicate outputs = %v, want %v", got, want)
	}
}

func TestMinStackMinInt(t *testing.T) {
	ops := []minStackOp{
		{name: "push", arg: math.MaxInt32},
		{name: "push", arg: math.MinInt32},
		{name: "getMin"},
		{name: "pop"},
		{name: "getMin"},
	}
	want := []int{math.MinInt32, math.MaxInt32}

	got := runMinStackOps(t, ops)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("minint outputs = %v, want %v", got, want)
	}
}
