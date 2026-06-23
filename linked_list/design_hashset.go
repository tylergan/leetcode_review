package linked_list

import "math/rand"

/*
Design a HashSet without using any built-in hash table libraries.

Implement MyHashSet class:

void add(key) Inserts the value key into the HashSet.
bool contains(key) Returns whether the value key exists in the HashSet or not.
void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.

Example 1:

Input: ["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]

Output: [null, null, null, true, false, null, true, null, false]
Explanation:
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1); // set = [1]
myHashSet.add(2); // set = [1, 2]
myHashSet.contains(1); // return True
myHashSet.contains(3); // return False, (not found)
myHashSet.add(2); // set = [1, 2]
myHashSet.contains(2); // return True
myHashSet.remove(2); // set = [1]
myHashSet.contains(2); // return False, (already removed)

Constraints:

0 <= key <= 1,000,000
At most 10,000 calls will be made to add, remove, and contains.

Implementation note: each bucket is its own skip list, so collisions resolve
in O(log m) instead of O(m) where m is the bucket's chain length. Named
NewMyHashSet (not Constructor) because the linked_list package already exposes
Constructor for LRUCache; Go only allows one package-level Constructor.
*/

type skipNode struct {
	val  int
	next []*skipNode
}

type MyHashSet struct {
	heads      []*skipNode
	rng        *rand.Rand
	maxSkipLvl int
	size       int
}

func NewMyHashSet() MyHashSet {
	maxSize := 10000
	heads := make([]*skipNode, 10000)

	maxSkipLvl := 20
	for i := range heads {
		heads[i] = &skipNode{next: make([]*skipNode, maxSkipLvl)}
	}

	return MyHashSet{
		heads:      heads,
		rng:        rand.New(rand.NewSource(1)),
		maxSkipLvl: maxSkipLvl,
		size:       maxSize,
	}
}

func (this *MyHashSet) hash(key int) int {
	return key % this.size
}

func (this *MyHashSet) getPreds(key int) []*skipNode {
	preds := make([]*skipNode, this.maxSkipLvl)

	curr := this.heads[this.hash(key)]
	for lvl := this.maxSkipLvl - 1; lvl >= 0; lvl-- {
		for curr.next[lvl] != nil && curr.next[lvl].val < key {
			curr = curr.next[lvl]
		}
		preds[lvl] = curr
	}

	return preds
}

func (this *MyHashSet) Add(key int) {
	preds := this.getPreds(key)
	if preds[0].next[0] != nil && preds[0].next[0].val == key { // been added before
		return
	}

	rLvl := this.randomLevel()
	node := &skipNode{val: key, next: make([]*skipNode, rLvl)}
	for lvl := 0; lvl < rLvl; lvl++ {
		node.next[lvl] = preds[lvl].next[lvl]
		preds[lvl].next[lvl] = node
	}
}

func (this *MyHashSet) randomLevel() int {
	lvl := 1
	for lvl < this.maxSkipLvl && this.rng.Intn(4) == 0 {
		lvl++
	}
	return lvl
}

func (this *MyHashSet) Remove(key int) {
	preds := this.getPreds(key)

	remNode := preds[0].next[0]
	if remNode == nil || remNode.val != key { // D.N.E
		return
	}

	for lvl := 0; lvl < len(remNode.next); lvl++ {
		if preds[lvl].next[lvl] != remNode {
			break // should not occur
		}
		preds[lvl].next[lvl] = remNode.next[lvl]
	}
}

func (this *MyHashSet) Contains(key int) bool {
	curr := this.heads[this.hash(key)]
	for lvl := this.maxSkipLvl - 1; lvl >= 0; lvl-- {
		for curr.next[lvl] != nil && curr.next[lvl].val < key {
			curr = curr.next[lvl]
		}
	}

	return curr.next[0] != nil && curr.next[0].val == key
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := NewMyHashSet();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
