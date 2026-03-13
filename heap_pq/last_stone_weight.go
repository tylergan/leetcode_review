package heappq

import "container/heap"

/*
You are given an array of integers stones where stones[i] represents the weight of the ith stone.

We want to run a simulation on the stones as follows:

At each step we choose the two heaviest stones, with weight x and y and smash them togethers
If x == y, both stones are destroyed
If x < y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
Continue the simulation until there is no more than one stone remaining.

Return the weight of the last remaining stone or return 0 if none remain.

Example 1:

Input: stones = [2,3,6,2,4]

Output: 1
Explanation:
We smash 6 and 4 and are left with a 2, so the array becomes [2,3,2,2].
We smash 3 and 2 and are left with a 1, so the array becomes [1,2,2].
We smash 2 and 2, so the array becomes [1].

Example 2:

Input: stones = [1,2]

Output: 1
Constraints:

1 <= stones.length <= 20
1 <= stones[i] <= 100
*/

func lastStoneWeight(stones []int) int {
	h := &IntHeapMax{}
	heap.Init(h)
	for _, stone := range stones {
		heap.Push(h, stone)
	}

	for h.Len() > 1 {
		stone1, stone2 := (heap.Pop(h)).(int), (heap.Pop(h)).(int)
		if stone1 == stone2 {
			continue
		}
		smashedRock := max(stone1, stone2) - min(stone1, stone2)
		heap.Push(h, smashedRock)
	}

	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}
