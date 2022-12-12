package lib

type QueueItem interface {
	Cost() int
	SetIndex(int)
}

type PQ []QueueItem

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].Cost() < pq[j].Cost()
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].SetIndex(i)
	pq[j].SetIndex(j)
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(QueueItem)
	item.SetIndex(n)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.SetIndex(-1)
	*pq = old[0 : n-1]
	return item
}
