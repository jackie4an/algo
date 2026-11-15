package divideandconquer

import (
	"cmp"
	"math"
	"slices"
)

type Point struct {
	X, Y float64
}

type Pair struct {
	Point1   Point
	Point2   Point
	Distance float64
}

func minPair(p1, p2 Pair) Pair {
	if p1.Distance <= p2.Distance {
		return p1
	} else {
		return p2
	}
}

func maxPair(p1, p2 Pair) Pair {
	if p1.Distance >= p2.Distance {
		return p1
	} else {
		return p2
	}
}

func ClosestPair(points []Point) Pair {
	py := make([]Point, len(points))
	copy(py, points)

	slices.SortFunc(points, cmpX)
	slices.SortFunc(py, cmpY)

	return closestPairRec(points, py)
}

func closestPairRec(px, py []Point) Pair {
	if len(px) <= 3 {
		return closestPairBruteForce(px)
	}

	mid := len(px) / 2
	pl := closestPairRec(px[:mid], py[:mid])
	pr := closestPairRec(px[mid:], py[mid:])

	recPair := minPair(pl, pr)
	stripPair := stripClosest(py, px[mid], recPair.Distance)

	return minPair(recPair, stripPair)
}

func closestPairBruteForce(points []Point) Pair {
	var p Pair

	p.Distance = math.MaxFloat64

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			d := distance(points[i], points[j])
			if d < p.Distance {
				p.Distance = d
				p.Point1 = points[i]
				p.Point2 = points[j]
			}
		}
	}

	return p
}

func stripClosest(py []Point, pmid Point, d float64) Pair {
	var (
		strip []Point
		p     Pair
	)

	for i := range py {
		if math.Abs(py[i].X-pmid.X) < d {
			strip = append(strip, py[i])
		}
	}

	p.Distance = d

	for i := range strip {
		for j := i + 1; j < len(strip) && (strip[j].Y-strip[i].Y) < p.Distance; j++ {
			dist := distance(strip[i], strip[j])
			if dist < p.Distance {
				p.Distance = dist
				p.Point1 = strip[i]
				p.Point2 = strip[j]
			}
		}
	}

	return p
}

func distance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	return math.Sqrt(dx*dx + dy*dy)
}

func cmpX(p1, p2 Point) int {
	c := cmp.Compare(p1.X, p2.X)
	if c == 0 {
		c = cmp.Compare(p1.Y, p2.Y)
	}

	return c
}

func cmpY(p1, p2 Point) int {
	c := cmp.Compare(p1.Y, p2.Y)
	if c == 0 {
		c = cmp.Compare(p1.X, p2.X)
	}

	return c
}
