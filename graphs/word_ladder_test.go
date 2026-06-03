package graphs

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{
			name:      "example 1 - shortest transformation exists",
			beginWord: "cat",
			endWord:   "sag",
			wordList:  []string{"bat", "bag", "sag", "dag", "dot"},
			want:      4,
		},
		{
			name:      "example 2 - end word not in list",
			beginWord: "cat",
			endWord:   "sag",
			wordList:  []string{"bat", "bag", "sat", "dag", "dot"},
			want:      0,
		},
		{
			name:      "classic example",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			name:      "end exists but no valid path",
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"hot", "dog"},
			want:      0,
		},
		{
			name:      "single character direct transformation",
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			want:      2,
		},
		{
			name:      "begin equals end",
			beginWord: "same",
			endWord:   "same",
			wordList:  []string{"came", "lame"},
			want:      1,
		},
		{
			name:      "chooses shorter bridge among multiple routes",
			beginWord: "aaa",
			endWord:   "bbb",
			wordList:  []string{"aab", "abb", "bbb", "aba", "baa", "bab"},
			want:      4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if got != tt.want {
				t.Fatalf("ladderLength(%q, %q, %v) = %d, want %d", tt.beginWord, tt.endWord, tt.wordList, got, tt.want)
			}
		})
	}
}
