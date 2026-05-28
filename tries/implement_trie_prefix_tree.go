package tries

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings.

There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the PrefixTree class:

PrefixTree() Initializes the prefix tree object.
void insert(String word) Inserts the string word into the prefix tree.
boolean search(String word) Returns true if the string word is in the prefix tree (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.

Example 1:

Input:
["PrefixTree", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]

Output:
[null, null, true, false, true, null, true]

Explanation:
PrefixTree prefixTree = new PrefixTree();
prefixTree.insert("apple");
prefixTree.search("apple");   // return true
prefixTree.search("app");     // return false
prefixTree.startsWith("app"); // return true
prefixTree.insert("app");
prefixTree.search("app");     // return true

Constraints:

1 <= word.length, prefix.length <= 2000
word and prefix consist only of lowercase English letters.
At most 3 * 10^4 calls in total will be made to insert, search, and startsWith.
*/

type Node struct {
	char       rune
	end        bool
	neighbours map[rune]*Node
}

type PrefixTree struct {
	head *Node
}

func Constructor() PrefixTree {
	return PrefixTree{head: &Node{neighbours: map[rune]*Node{}}}
}

func (this *PrefixTree) Insert(word string) {
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

func (this *PrefixTree) Search(word string) bool {
	node := this.find(word)
	if node == nil {
		return false
	}
	return node.end
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this.find(prefix)
	return node != nil
}

func (this *PrefixTree) find(word string) *Node {
	curr := this.head
	for _, char := range word {
		node, ok := curr.neighbours[char]
		if !ok {
			return nil
		}

		curr = node
	}

	return curr
}
