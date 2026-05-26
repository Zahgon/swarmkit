package scheduler

type nodeMaxHeap struct {
	nodes    []NodeInfo
	lessFunc func(*NodeInfo, *NodeInfo) bool
	length   int
}

func (h nodeMaxHeap) Len() int { _ = "STUB: not implemented"; return 0 }

func (h nodeMaxHeap) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (h nodeMaxHeap) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// reversed to make a max-heap
	return false
}

func (h *nodeMaxHeap) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (h *nodeMaxHeap) Pop() interface{} {
	_ = "STUB: not implemented"

	// return value is never used
	return nil
}
