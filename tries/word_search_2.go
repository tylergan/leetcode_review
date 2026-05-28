package tries

/*
Given a 2-D grid of characters board and a list of strings words, return all words that are present in the grid.

For a word to be present it must be possible to form the word with a path in the board with horizontally or vertically neighboring cells. The same cell may not be used more than once in a word.

Example 1:

Input:
board = [
  ["a","b","c","d"],
  ["s","a","a","t"],
  ["a","c","k","e"],
  ["a","c","d","n"]
],
words = ["bat","cat","back","backend","stack"]

Output: ["cat","back","backend"]
Example 2:

Input:
board = [
  ["x","o"],
  ["x","o"]
],
words = ["xoxo"]

Output: []
Constraints:

1 <= board.length, board[i].length <= 12
board[i] consists only of lowercase English letter.
1 <= words.length <= 30,000
1 <= words[i].length <= 10
words[i] consists only of lowercase English letters.
All strings within words are distinct.
*/

func newTrie(words []string) *Node {
	head := &Node{neighbours: map[rune]*Node{}}

	for _, word := range words {
		insert(head, word)
	}

	return head
}

func insert(head *Node, word string) {
	curr := head
	for i := range len(word) {
		char := rune(word[i])

		node, ok := curr.neighbours[char]
		if !ok {
			node = &Node{
				char:       char,
				neighbours: map[rune]*Node{},
			}

			curr.neighbours[char] = node
		}

		curr = node
	}

	curr.word = word
}

func findWords(board [][]byte, words []string) []string {
	head := newTrie(words)

	var res []string
	m, n := len(board), len(board[0])

	var dfs func(*Node, int, int)
	dfs = func(node *Node, i int, j int) {
		char := board[i][j]
		child, ok := node.neighbours[rune(char)]
		if !ok {
			return
		}

		if child.word != "" {
			res = append(res, child.word)
			child.word = "" // Avoid duplicates by marking the word as found.
		}

		board[i][j] = byte('.')

		dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
		for _, dir := range dirs {
			rowIdx, colIdx := i+dir[0], j+dir[1]
			if rowIdx < 0 || rowIdx >= m || colIdx < 0 || colIdx >= n {
				continue
			}

			dfs(child, rowIdx, colIdx)
		}

		board[i][j] = char
	}

	for i := range m {
		for j := range n {
			dfs(head, i, j)
		}
	}

	return res
}
