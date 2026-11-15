package main

import "fmt"

type color int

const (
	red color = iota + 1
	white
	blue
)

func (c color) String() string {
	switch c {
	case red:
		return "red"
	case white:
		return "white"
	case blue:
		return "blue"
	default:
		return "unknown"
	}
}

func main() {
	items := []color{
		blue,
		white,
		blue,
		red,
		white,
	}

	fmt.Println("before:", items)

	for i := 1; i < len(items); i++ {
		j := i - 1
		ci := items[i]
		cj := items[j]
		for j >= 0 && ci < cj {
			swap(items, i, j)
			i = j
			j--
			cj = items[j]
		}
	}

	fmt.Println("after:", items)
}

func swap(arr []color, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
