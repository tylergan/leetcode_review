package arrays

import (
	"reflect"
	"testing"
)

func TestEncodeDecodeRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		{
			name: "example 1: normal strings",
			strs: []string{"neet", "code", "love", "you"},
		},
		{
			name: "example 2: with special characters",
			strs: []string{"we", "say", ":", "yes"},
		},
		{
			name: "empty list",
			strs: []string{},
		},
		{
			name: "single empty string",
			strs: []string{""},
		},
		{
			name: "multiple empty strings",
			strs: []string{"", "", ""},
		},
		{
			name: "single string",
			strs: []string{"hello"},
		},
		{
			name: "strings with spaces",
			strs: []string{"hello world", "foo bar", "test"},
		},
		{
			name: "strings with numbers",
			strs: []string{"test123", "456", "abc789def"},
		},
		{
			name: "strings with special characters",
			strs: []string{"hello#world", "foo@bar", "test$123"},
		},
		{
			name: "strings with delimiters that might be problematic",
			strs: []string{"a,b,c", "x:y:z", "p|q|r"},
		},
		{
			name: "strings with newlines",
			strs: []string{"line1\nline2", "test\n\ntest"},
		},
		{
			name: "mixed empty and non-empty",
			strs: []string{"hello", "", "world", "", ""},
		},
		{
			name: "very long string",
			strs: []string{"a very long string that contains many characters to test encoding and decoding of longer content"},
		},
		{
			name: "unicode characters",
			strs: []string{"hello", "世界", "🌍", "café"},
		},
		{
			name: "strings with quotes",
			strs: []string{"\"quoted\"", "'single'", "mix\"ed'quotes"},
		},
		{
			name: "strings with backslashes",
			strs: []string{"path\\to\\file", "escape\\n\\t"},
		},
		{
			name: "strings that look like length prefixes",
			strs: []string{"4#test", "10#", "#"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solution{}
			encoded := s.Encode(tt.strs)
			decoded := s.Decode(encoded)

			if !reflect.DeepEqual(decoded, tt.strs) {
				t.Errorf("Round trip failed.\nOriginal: %v\nEncoded: %q\nDecoded: %v", tt.strs, encoded, decoded)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		{
			name: "basic encode",
			strs: []string{"hello", "world"},
		},
		{
			name: "empty list encode",
			strs: []string{},
		},
		{
			name: "single string encode",
			strs: []string{"test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solution{}
			encoded := s.Encode(tt.strs)

			// The encoded string should be non-nil
			if encoded == "" && len(tt.strs) > 0 {
				// Only fail if we expect some output
				for _, str := range tt.strs {
					if str != "" {
						t.Errorf("Encode() returned empty string for non-empty input")
						break
					}
				}
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		validate func(t *testing.T, original, decoded []string)
	}{
		{
			name:  "decode preserves order",
			input: []string{"first", "second", "third"},
			validate: func(t *testing.T, original, decoded []string) {
				if len(decoded) != len(original) {
					t.Errorf("Length mismatch: got %d, want %d", len(decoded), len(original))
					return
				}
				for i := range original {
					if decoded[i] != original[i] {
						t.Errorf("Order not preserved at index %d: got %q, want %q", i, decoded[i], original[i])
					}
				}
			},
		},
		{
			name:  "decode preserves exact content",
			input: []string{"exact", "content", "preserved"},
			validate: func(t *testing.T, original, decoded []string) {
				for i := range original {
					if decoded[i] != original[i] {
						t.Errorf("Content mismatch at index %d: got %q, want %q", i, decoded[i], original[i])
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solution{}
			encoded := s.Encode(tt.input)
			decoded := s.Decode(encoded)
			tt.validate(t, tt.input, decoded)
		})
	}
}
