package zfunc

func Eval(s string) []int {
	runes := []rune(s)
	z := make([]int, len(runes))
	l, r := 1, 1

	for i := 1; i < len(runes); i++ {
		if i > r {
			j := 0
			for j < len(runes) && r < len(runes) &&
				runes[j] == runes[r] {
				j++
				r++
			}
			if r-i > 0 {
				l = i
				r--
				z[i] = r - l + 1
			}
		} else {
			a := r - l + 1
			b := r - i + 1
			k := a - b
			if z[k] < b {
				z[i] = z[k]
			} else {
				j := b + 1
				q := r + 1
				for j < len(runes) && q < len(runes) &&
					runes[j] == runes[q] {
					j++
					q++
				}
				if q-r > 1 {
					l = i
					r = q - 1
					z[i] = q - i
				}
			}
		}
	}

	return z
}
