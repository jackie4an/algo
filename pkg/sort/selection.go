package sort

import (
	"cmp"
	"fmt"
	"strings"
)

type PrioQue[T cmp.Ordered] struct {
	que []T
	n   int
}

func NewPrioQue[T cmp.Ordered](items ...T) *PrioQue[T] {
	pq := &PrioQue[T]{
		que: make([]T, len(items)+1),
		n:   0,
	}

	for _, item := range items {
		pq.Insert(item)
	}

	return pq
}

func (pq PrioQue[T]) String() string {
	var sb strings.Builder

	sb.WriteString("[ ")
	for i := 1; i < len(pq.que); i++ {
		sb.WriteString(fmt.Sprintf("%v", pq.que[i]))
		if i != len(pq.que)-1 {
			sb.WriteByte(',')
		}
		sb.WriteByte(' ')
	}
	sb.WriteString("]")

	return sb.String()
}

func (pq *PrioQue[T]) Parent(i int) int {
	if i <= 1 {
		return -1
	}

	return i / 2
}

func (pq *PrioQue[T]) LeftChild(i int) int {
	if i == 0 {
		return -1
	}

	return 2 * i
}

func (pq *PrioQue[T]) RightChild(i int) int {
	if i == 0 {
		return -1
	}

	return 2*i + 1
}

func (pq *PrioQue[T]) Insert(item T) {
	if pq.n >= len(pq.que) {
		fmt.Println("Warning: priority queue overflow!")
	} else {
		pq.n++
		pq.que[pq.n] = item
		pq.bubbleUp(pq.n)
	}
}

func (pq *PrioQue[T]) ExtractMin() T {
	var m T

	if pq.n <= 0 {
		fmt.Println("Warning: empty priority queue")
	} else {
		m = pq.que[1]
		pq.que[1] = pq.que[pq.n]
		pq.n--
		pq.bubbleDown(1)
	}

	return m
}

func (pq *PrioQue[T]) bubbleDown(p int) {
	c := pq.LeftChild(p)
	minIndex := p

	for i := 0; i <= 1; i++ {
		if c+i <= pq.n &&
			cmp.Compare(pq.que[minIndex], pq.que[c+i]) == 1 {
			minIndex = c + i
		}
	}

	if minIndex != p {
		pq.swap(p, minIndex)
		pq.bubbleDown(minIndex)
	}
}

func (pq *PrioQue[T]) bubbleUp(p int) {
	parent := pq.Parent(p)
	if parent == -1 {
		return // root
	}

	if cmp.Compare(pq.que[parent], pq.que[p]) == 1 {
		pq.swap(p, parent)
		pq.bubbleUp(parent)
	}
}

func (pq *PrioQue[T]) swap(i, j int) {
	pq.que[i], pq.que[j] = pq.que[j], pq.que[i]
}

func Heapsort[T cmp.Ordered](items []T) {
	q := NewPrioQue(items...)

	for i := range items {
		items[i] = q.ExtractMin()
	}
}
