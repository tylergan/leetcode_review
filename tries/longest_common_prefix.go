package tries

import "strings"

/*
You are given an array of strings strs. Return the longest common prefix of all the strings.

If there is no longest common prefix, return an empty string "".

Example 1:

Input: strs = ["bat","bag","bank","band"]

Output: "ba"

Example 2:

Input: strs = ["dance","dag","danger","damage"]

Output: "da"

Example 3:

Input: strs = ["neet","feet"]

Output: ""

Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] is made up of lowercase English letters if it is non-empty.
*/

type TrieNode struct {
	char       rune
	neighbours map[rune]*TrieNode
}

// insertPrefix inserts str into the trie and returns the prefix of str that
// already existed in the trie. Because a brand-new node always has empty
// neighbours, every char after the first miss is new too, so the returned
// prefix is the contiguous prefix str shares with the words inserted so far.
func insertPrefix(str string, trie *TrieNode) string {
	var prefix strings.Builder

	for _, c := range str {
		nextNode := &TrieNode{
			char:       c,
			neighbours: map[rune]*TrieNode{},
		}

		if _, ok := trie.neighbours[c]; !ok {
			trie.neighbours[c] = nextNode
		} else {
			prefix.WriteString(string(c))
		}

		trie = trie.neighbours[c]
	}

	return prefix.String()
}

func longestCommonPrefix(strs []string) string {
	trie := &TrieNode{
		neighbours: map[rune]*TrieNode{},
	}

	insertPrefix(strs[0], trie) // first insertion will return an empty prefix

	validPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if prefix := insertPrefix(strs[i], trie); len(prefix) < len(validPrefix) { // a valid prefix is just the shortest prefix found from the entire array
			validPrefix = prefix
		}
	}

	return validPrefix
}
