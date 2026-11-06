package divideandconquer

import "math"

func MaxSubArray(array []int) ([]int, int) {
	l, r, sum := maxSubArray(array, 0, len(array)-1)

	return array[l : r+1], sum
}

func maxSubArray(array []int, l, r int) (nl, nr, sum int) {
	if l == r {
		return l, r, array[l]
	}

	m := (l + r) / 2
	ll, lr, ls := maxSubArray(array, l, m)
	rl, rr, rs := maxSubArray(array, m+1, r)
	cl, cr, cs := crossSum(array, l, m, r)

	if ls >= rs && ls >= cs {
		return ll, lr, ls
	} else if rs >= ls && rs >= cs {
		return rl, rr, rs
	} else if cs >= ls && cs >= rs {
		return cl, cr, cs
	}

	return
}

func crossSum(array []int, l, m, r int) (nl, nr, sum int) {
	leftSum := math.MinInt
	for i := m; i >= l; i-- {
		sum += array[i]
		if sum > leftSum {
			leftSum = sum
			nl = i
		}
	}

	sum = 0
	rightSum := math.MinInt
	for i := m + 1; i <= r; i++ {
		sum += array[i]
		if sum > rightSum {
			rightSum = sum
			nr = i
		}
	}

	return nl, nr, leftSum + rightSum
}
