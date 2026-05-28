package tries

import "testing"

func TestWordDictionaryExampleFlow(t *testing.T) {
	wordDictionary := NewWordDictionary()

	wordDictionary.AddWord("day")
	wordDictionary.AddWord("bay")
	wordDictionary.AddWord("may")

	tests := []struct {
		word string
		want bool
	}{
		{word: "say", want: false},
		{word: "day", want: true},
		{word: ".ay", want: true},
		{word: "b..", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := wordDictionary.Search(tt.word)
			if got != tt.want {
				t.Fatalf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestWordDictionaryWildcardPositions(t *testing.T) {
	wordDictionary := NewWordDictionary()

	for _, word := range []string{"bad", "dad", "mad", "pad"} {
		wordDictionary.AddWord(word)
	}

	tests := []struct {
		word string
		want bool
	}{
		{word: ".ad", want: true},
		{word: "b.d", want: true},
		{word: "ba.", want: true},
		{word: "...", want: true},
		{word: "..d", want: true},
		{word: "b..", want: true},
		{word: ".a.", want: true},
		{word: ".z.", want: false},
		{word: "b.z", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := wordDictionary.Search(tt.word)
			if got != tt.want {
				t.Fatalf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestWordDictionarySearchIsLengthSensitive(t *testing.T) {
	wordDictionary := NewWordDictionary()

	wordDictionary.AddWord("b")
	wordDictionary.AddWord("bad")

	tests := []struct {
		word string
		want bool
	}{
		{word: "b", want: true},
		{word: "b.", want: false},
		{word: "b..", want: true},
		{word: ".", want: true},
		{word: "..", want: false},
		{word: "...", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := wordDictionary.Search(tt.word)
			if got != tt.want {
				t.Fatalf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestWordDictionarySharedPrefixes(t *testing.T) {
	wordDictionary := NewWordDictionary()

	wordDictionary.AddWord("at")
	wordDictionary.AddWord("atom")
	wordDictionary.AddWord("attic")

	tests := []struct {
		word string
		want bool
	}{
		{word: "at", want: true},
		{word: "a.", want: true},
		{word: "a..", want: false},
		{word: "ato.", want: true},
		{word: "att..", want: true},
		{word: "atti.", want: true},
		{word: "attic", want: true},
		{word: "atti", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := wordDictionary.Search(tt.word)
			if got != tt.want {
				t.Fatalf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
