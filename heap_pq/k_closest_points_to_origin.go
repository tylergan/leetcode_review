package heappq

import (
	"container/heap"
)

/*You are given an 2-D array points where points[i] = [xi, yi] represents the coordinates of a point on an X-Y axis plane. You are also given an integer k.

Return the k closest points to the origin (0, 0).

The distance between two points is defined as the Euclidean distance (sqrt((x1 - x2)^2 + (y1 - y2)^2)).

You may return the answer in any order.

Example 1:



Input: points = [[0,2],[2,2]], k = 1

Output: [[0,2]]
Explanation : The distance between (0, 2) and the origin (0, 0) is 2. The distance between (2, 2) and the origin is sqrt(2^2 + 2^2) = 2.82842. So the closest point to the origin is (0, 2).

Example 2:

Input: points = [[0,2],[2,0],[2,2]], k = 2

Output: [[0,2],[2,0]]
Explanation: The output [2,0],[0,2] would also be accepted.

Constraints:

1 <= k <= points.length <= 1000
-100 <= points[i][0], points[i][1] <= 100
*/

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// no need sqrt since it won't change ordering between dist1 and dist2
	// and reduces the computational complexity
	dist1 := pq[i][0]*pq[i][0] + pq[i][1]*pq[i][1]
	dist2 := pq[j][0]*pq[j][0] + pq[j][1]*pq[j][1]
	return dist1 > dist2
}

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.([]int)) }

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	old := *pq
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, point := range points {
		heap.Push(pq, point)
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	return *pq
}
