package prioque

import "container/heap"

type prio byte

const (
	Min prio = iota + 1
	Max
)

type Item[T any] struct {
	Value    T
	Priority int
	index    int
}

type PriorityQueue[T any] struct {
	pq *priorityQueue[T]
}

type priorityQueue[T any] struct {
	que  []*Item[T]
	prio prio
}

func New[T any](prio prio, items ...*Item[T]) *PriorityQueue[T] {
	q := &PriorityQueue[T]{
		pq: &priorityQueue[T]{
			que:  make([]*Item[T], len(items)),
			prio: prio,
		},
	}

	copy(q.pq.que, items)
	heap.Init(q.pq)

	return q
}

func (q *PriorityQueue[T]) Len() int {
	return len(q.pq.que)
}

func (q *PriorityQueue[T]) Push(item *Item[T]) {
	heap.Push(q.pq, item)
}

func (q *PriorityQueue[T]) Pop() *Item[T] {
	if item, ok := heap.Pop(q.pq).(*Item[T]); ok {
		return item
	} else {
		return nil
	}
}

func (q *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(q.pq, item.index)
}

func (q *PriorityQueue[T]) Remove(index int) *Item[T] {
	item := heap.Remove(q.pq, index)
	if item, ok := item.(*Item[T]); ok {
		return item
	} else {
		return nil
	}
}

func (q *priorityQueue[T]) Len() int {
	return len(q.que)
}

func (q *priorityQueue[T]) Less(i int, j int) bool {
	if q.prio == Min {
		return q.que[i].Priority < q.que[j].Priority
	} else {
		return q.que[i].Priority > q.que[j].Priority
	}
}

func (q *priorityQueue[T]) Swap(i int, j int) {
	q.que[i], q.que[j] = q.que[j], q.que[i]
	q.que[i].index = i
	q.que[j].index = j
}

func (q *priorityQueue[T]) Push(x any) {
	if x, ok := x.(*Item[T]); ok {
		q.que = append(q.que, x)
		x.index = len(q.que)
	}
}

func (q *priorityQueue[T]) Pop() any {
	n := len(q.que)
	x := q.que[n-1]
	x.index = -1
	q.que = q.que[:n-1]
	return x
}
