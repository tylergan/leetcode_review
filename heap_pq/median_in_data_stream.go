package heappq

import "container/heap"

/*
The median is the middle value in a sorted list of integers. For lists of even length, there is no middle value, so the median is the mean of the two middle values.

For example:

For arr = [1,2,3], the median is 2.
For arr = [1,2], the median is (1 + 2) / 2 = 1.5
Implement the MedianFinder class:

MedianFinder() initializes the MedianFinder object.
void addNum(int num) adds the integer num from the data stream to the data structure.
double findMedian() returns the median of all elements so far.
Example 1:

Input:
["MedianFinder", "addNum", "1", "findMedian", "addNum", "3" "findMedian", "addNum", "2", "findMedian"]

Output:
[null, null, 1.0, null, 2.0, null, 2.0]

Explanation:
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.findMedian(); // return 1.0
medianFinder.addNum(3);    // arr = [1, 3]
medianFinder.findMedian(); // return 2.0
medianFinder.addNum(2);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
Constraints:

-100,000 <= num <= 100,000
findMedian will only be called after adding at least one integer to the data structure.
*/

type MedianFinder struct {
	leftSide  *IntHeapMax // this contains the largest element of the smaller half at the root
	rightSide *IntHeapMin // this contains the smallest element of the larger half at the root
}

func NewMedianFinder() MedianFinder {
	leftSide, rightSide := &IntHeapMax{}, &IntHeapMin{}
	heap.Init(leftSide)
	heap.Init(rightSide)
	return MedianFinder{
		leftSide:  leftSide,
		rightSide: rightSide,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.rightSide.Len() > 0 {
		if num > (*this.rightSide)[0] { // push largest elements to right side
			heap.Push(this.rightSide, num)
		} else { // push smallest elements to left side
			heap.Push(this.leftSide, num)
		}
	} else {
		heap.Push(this.leftSide, num)
	}

	// if the sides differ by more than one, than we need to rebalance so that
	// we can get the median in constant time
	if this.leftSide.Len() > this.rightSide.Len()+1 { // too many elements in left side, push one to right side
		heap.Push(this.rightSide, heap.Pop(this.leftSide))
	} else if this.rightSide.Len() > this.leftSide.Len()+1 { // too many elements in right side, push one to left side
		heap.Push(this.leftSide, heap.Pop(this.rightSide))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// odd case
	if this.leftSide.Len() > this.rightSide.Len() {
		return float64((*this.leftSide)[0])
	}
	if this.rightSide.Len() > this.leftSide.Len() {
		return float64((*this.rightSide)[0])
	}

	// even case
	return float64((*this.leftSide)[0]+(*this.rightSide)[0]) / 2.0
}
