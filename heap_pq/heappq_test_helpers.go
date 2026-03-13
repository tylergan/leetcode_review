package heappq

type IntHeapMin []int

func (h IntHeapMin) Len() int { return len(h) }

func (h IntHeapMin) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeapMin) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeapMin) Push(val any) { *h = append(*h, val.(int)) }

func (h *IntHeapMin) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type IntHeapMax []int

func (h IntHeapMax) Len() int { return len(h) }

func (h IntHeapMax) Less(i, j int) bool { return h[i] > h[j] }

func (h IntHeapMax) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeapMax) Push(x any) { *h = append(*h, x.(int)) }

func (h *IntHeapMax) Pop() any {
	n := len(*h)
	old := *h
	x := old[n-1]
	*h = old[:n-1]
	return x
}
