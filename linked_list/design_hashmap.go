package linked_list

import "math/rand"

/*
Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:

MyHashMap() initializes the object with an empty map.
void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.

Example 1:

Input: ["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]

Output: [null, null, null, 1, -1, null, 1, null, -1]
Explanation:
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // The map is now [[1,1]]
myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
myHashMap.get(1); // return 1, The map is now [[1,1], [2,2]]
myHashMap.get(3); // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
myHashMap.get(2); // return 1, The map is now [[1,1], [2,1]]
myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
myHashMap.get(2); // return -1 (i.e., not found), The map is now [[1,1]]

Constraints:

0 <= key, value <= 1,000,000
At most 10,000 calls will be made to put, get, and remove.

Implementation note: the key-only counterpart Design HashSet lives in the same
package, so this file uses mapSkipNode (the bucket node carries a value too) and
NewMyHashMap (the package already exports Constructor for LRUCache and only one
package-level Constructor is allowed).
*/

type mapSkipNode struct {
	key  int
	val  int
	next []*mapSkipNode
}

type MyHashMap struct {
	heads      []*mapSkipNode
	maxSkipLvl int
	size       int
	rng        *rand.Rand
}

func NewMyHashMap() MyHashMap {
	maxSize := 10000
	heads := make([]*mapSkipNode, maxSize)

	maxSkipLvl := 20
	for i := range maxSize {
		heads[i] = &mapSkipNode{next: make([]*mapSkipNode, maxSkipLvl)}
	}

	return MyHashMap{
		heads:      heads,
		maxSkipLvl: maxSkipLvl,
		size:       maxSize,
		rng:        rand.New(rand.NewSource(1)),
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.size
}

func (this *MyHashMap) getPreds(key int) []*mapSkipNode {
	preds := make([]*mapSkipNode, this.maxSkipLvl)

	curr := this.heads[this.hash(key)]
	for lvl := this.maxSkipLvl - 1; lvl >= 0; lvl-- {
		for curr.next[lvl] != nil && curr.next[lvl].key < key {
			curr = curr.next[lvl]
		}
		preds[lvl] = curr
	}

	return preds
}

func (this *MyHashMap) Put(key int, value int) {
	preds := this.getPreds(key)

	if preds[0].next[0] != nil && preds[0].next[0].key == key { // already exists, just override the value
		preds[0].next[0].val = value
		return
	}

	rLvl := this.randLvl()
	node := &mapSkipNode{key: key, val: value, next: make([]*mapSkipNode, rLvl)}
	for lvl := 0; lvl < rLvl; lvl++ {
		node.next[lvl] = preds[lvl].next[lvl]
		preds[lvl].next[lvl] = node
	}
}

func (this *MyHashMap) randLvl() int {
	lvl := 1
	for lvl < this.maxSkipLvl && this.rng.Intn(4) == 0 {
		lvl++
	}
	return lvl
}

func (this *MyHashMap) Get(key int) int {
	curr := this.heads[this.hash(key)]
	for lvl := this.maxSkipLvl - 1; lvl >= 0; lvl-- {
		for curr.next[lvl] != nil && curr.next[lvl].key < key {
			curr = curr.next[lvl]
		}
	}

	res := curr.next[0]
	if res == nil || res.key != key {
		return -1
	}
	return res.val
}

func (this *MyHashMap) Remove(key int) {
	preds := this.getPreds(key)

	remNode := preds[0].next[0]
	if remNode == nil || remNode.key != key { // D.N.E
		return
	}

	for lvl := 0; lvl < len(remNode.next); lvl++ {
		preds[lvl].next[lvl] = remNode.next[lvl]
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := NewMyHashMap();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
