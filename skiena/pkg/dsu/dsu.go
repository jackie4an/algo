package dsu

type DSU[T comparable] struct {
	p    map[T]T
	size map[T]int
}

func New[T comparable](items []T) *DSU[T] {
	d := DSU[T]{
		p:    make(map[T]T, len(items)),
		size: make(map[T]int, len(items)),
	}

	for _, item := range items {
		d.p[item] = item
		d.size[item]++
	}

	return &d
}

func (d *DSU[T]) Union(a, b T) bool {
	ra, rb := d.Find(a), d.Find(b)

	if ra == rb {
		return false
	}

	if d.size[ra] < d.size[rb] {
		d.merge(ra, rb)
	} else {
		d.merge(rb, ra)
	}

	return true
}

func (d *DSU[T]) Find(x T) T {
	if x != d.p[x] {
		d.p[x] = d.Find(d.p[x])
	}

	return d.p[x]
}

func (d *DSU[T]) merge(r1, r2 T) {
	d.size[r2] = d.size[r2] + d.size[r1]
	d.p[r1] = r2
}
