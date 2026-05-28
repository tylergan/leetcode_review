package tries

/*
Design a data structure that supports adding new words and searching for existing words.

Implement the WordDictionary class:

void addWord(word) Adds word to the data structure.
bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.

Example 1:

Input:
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["day"],["bay"],["may"],["say"],["day"],[".ay"],["b.."]]

Output:
[null, null, null, null, false, true, true, true]

Explanation:
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("day");
wordDictionary.addWord("bay");
wordDictionary.addWord("may");
wordDictionary.search("say"); // return false
wordDictionary.search("day"); // return true
wordDictionary.search(".ay"); // return true
wordDictionary.search("b.."); // return true

Constraints:

1 <= word.length <= 20
word in addWord consists of lowercase English letters.
word in search consist of '.' or lowercase English letters.
There will be at most 2 dots in word for search queries.
At most 10,000 calls will be made to addWord and search.
*/

type WordDictionary struct {
	head *Node
}

func NewWordDictionary() WordDictionary {
	return WordDictionary{head: &Node{neighbours: map[rune]*Node{}}}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.head
	for _, char := range word {
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

	curr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(this.head, word, 0)
}

func (this *WordDictionary) dfs(node *Node, word string, i int) bool {
	if node == nil {
		return false
	}

	if i == len(word) {
		return node.end
	}

	if word[i] == byte('.') {
		for _, neighbour := range node.neighbours {
			if found := this.dfs(neighbour, word, i+1); found {
				return true
			}
		}
		return false
	}
	return this.dfs(node.neighbours[rune(word[i])], word, i+1)
}
