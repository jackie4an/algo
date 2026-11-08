package pq

type prio byte

const (
	Min prio = iota + 1
	Max
)

type Item interface {
	Priority() int
}

type PriorityQueue struct {
	que  []Item
	prio prio
}

func New(prio prio) *PriorityQueue {
	return &PriorityQueue{prio: prio}
}

func (q *PriorityQueue) Len() int {
	return len(q.que)
}

func (q *PriorityQueue) Less(i int, j int) bool {
	if q.prio == Min {
		return q.que[i].Priority() < q.que[j].Priority()
	} else {
		return q.que[i].Priority() > q.que[j].Priority()
	}
}

func (q *PriorityQueue) Swap(i int, j int) {
	q.que[i], q.que[j] = q.que[j], q.que[i]
}

func (q *PriorityQueue) Push(x any) {
	if x, ok := x.(Item); ok {
		q.que = append(q.que, x)
	}
}

func (q *PriorityQueue) Pop() any {
	x := q.que[len(q.que)-1]
	q.que = q.que[:len(q.que)-1]
	return x
}
