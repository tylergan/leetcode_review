package linked_list

/*
Implement the Least Recently Used (LRU) cache class LRUCache. The class should support the following operations

LRUCache(int capacity) Initialize the LRU cache of size capacity.
int get(int key) Return the value corresponding to the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the introduction of the new pair causes the cache to exceed its capacity, remove the least recently used key.
A key is considered used if a get or a put operation is called on it.

Ensure that get and put each run in O(1) average time complexity.

Example 1:

Input:
["LRUCache", [2], "put", [1, 10],  "get", [1], "put", [2, 20], "put", [3, 30], "get", [2], "get", [1]]

Output:
[null, null, 10, null, null, 20, -1]

Explanation:
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 10);  // cache: {1=10}
lRUCache.get(1);      // return 10
lRUCache.put(2, 20);  // cache: {1=10, 2=20}
lRUCache.put(3, 30);  // cache: {2=20, 3=30}, key=1 was evicted
lRUCache.get(2);      // returns 20
lRUCache.get(1);      // return -1 (not found)
Constraints:

1 <= capacity <= 100
0 <= key <= 1000
0 <= value <= 1000
*/

type lruNode struct {
	Key  int
	Val  int
	Next *lruNode
	Prev *lruNode
}

type LRUCache struct {
	head     *lruNode
	tail     *lruNode
	mapping  map[int]*lruNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	head, tail := &lruNode{}, &lruNode{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		head:     head,
		tail:     tail,
		mapping:  map[int]*lruNode{},
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.mapping[key]; ok {
		l.remove(node)
		l.add(node)
		return node.Val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	// now newly accessed, remove it from DLL
	if node, ok := l.mapping[key]; ok {
		l.remove(node)
	}

	// add most recent to dummy tail and override mapping
	node := &lruNode{Key: key, Val: value}
	l.add(node)
	l.mapping[key] = node

	// remove lru if we go beyond capacity
	if len(l.mapping) > l.capacity {
		lru := l.head.Next
		l.remove(lru)
		delete(l.mapping, lru.Key)
	}
}

func (l *LRUCache) remove(node *lruNode) {
	nodePrev, nodeNext := node.Prev, node.Next
	nodePrev.Next = nodeNext
	nodeNext.Prev = nodePrev
}

func (l *LRUCache) add(newMostRecent *lruNode) {
	mostRecent := l.tail.Prev

	// connect last mostRecent to the newMostRecent
	mostRecent.Next = newMostRecent
	newMostRecent.Prev = mostRecent

	// connect newMostRecent to the dummy tail
	l.tail.Prev = newMostRecent
	newMostRecent.Next = l.tail
}
