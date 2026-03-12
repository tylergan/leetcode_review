package binary_tree

import (
	"reflect"
	"testing"
)

func TestSerialiseAndDeserialiseBtree(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
	}{
		{
			name:  "empty tree",
			input: nil,
		},
		{
			name:  "single node",
			input: []*int{intPtr(1)},
		},
		{
			name:  "example 1",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), nil, nil, intPtr(4), intPtr(5)},
		},
		{
			name:  "left skewed",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3)},
		},
		{
			name:  "right skewed",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3)},
		},
		{
			name:  "negative values",
			input: []*int{intPtr(0), intPtr(-1), intPtr(-2)},
		},
		{
			name:  "complete tree",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)},
		},
		{
			name:  "sparse tree",
			input: []*int{intPtr(1), nil, intPtr(2), intPtr(3), nil, nil, intPtr(4)},
		},
	}

	codec := Constructor()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			serialized := codec.serialize(root)
			deserialized := codec.deserialize(serialized)
			got := serializeTree(deserialized)
			if !reflect.DeepEqual(got, tt.input) {
				t.Fatalf("round-trip failed: got %v, want %v", got, tt.input)
			}
		})
	}
}
