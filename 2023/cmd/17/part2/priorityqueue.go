package main

type priorityQueue []*cell

func (h priorityQueue) Len() int {
	return len(h)
}

func (h priorityQueue) Less(i, j int) bool {
	return h[i].fLoss < h[j].fLoss
}

func (h priorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIndex = i
	h[j].heapIndex = j
}

func (h *priorityQueue) Push(x any) {
	n := len(*h)
	item := x.(*cell)
	item.heapIndex = n
	*h = append(*h, item)
}

func (h *priorityQueue) Pop() any {
	old := *h
	n := len(old)
	cell := old[n-1]
	old[n-1] = nil
	cell.heapIndex = -1
	*h = old[0 : n-1]
	return cell
}
