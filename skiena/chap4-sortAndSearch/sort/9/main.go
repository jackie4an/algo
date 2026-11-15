package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	s := []int{-12, 12, 0, -3, 23, 2, 12, 94, 13, 5, -2}
	tcs := []struct {
		t int
		k int
	}{
		{
			t: 92,
			k: 2,
		},
	}

	slices.Sort(s)
	for i, tc := range tcs {
		fmt.Printf("test case %d: ", i+1)
		res := solve(s, tc.t, tc.k)
		if res != nil {
			fmt.Printf("target: %d, k: %d, result: %v\n", tc.t, tc.k, res)
		} else {
			fmt.Printf("can't find %d element(s) with target sum = %d\n", tc.k, tc.t)
		}
	}
}

/*
sort s: O(nlogn)

k = 1, then find s[i] == t

	O(logn)

k = 2, then find s[i] + s[j] == t:

	for i = [0, n):
		take s[i], find t = t - s[i] with binary search in O(logn)
	O(nlogn)

k = 3, then find s[i] + s[j] + s[k] == t:

	for i = [0, n):
		take s[i], find t = t - s[i]:
			for j = [i+1, n):
				f(k-1) = f(2): O(nlogn)
	O(n^2logn)

...
*/

func solve(nums []int, t, k int) []int {
	if k == 0 {
		return nil
	}

	if k == 1 {
		if search(nums, t) {
			return []int{t}
		} else {
			return nil
		}
	}

	res := make([]int, 0, k)

	for i := range nums {
		if next := solve(nums[i+1:], t-nums[i], k-1); next != nil {
			res = append(res, nums[i])
			res = append(res, next...)

			return res
		}
	}

	return nil
}

func target2Sum(nums []int, target int) []int {
	for i := range nums {
		if search(nums, target-nums[i]) {
			return []int{nums[i], target - nums[i]}
		}
	}

	return nil
}

func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		switch cmp.Compare(nums[mid], target) {
		case -1:
			low = mid + 1
		case 0:
			return true
		case 1:
			high = mid - 1
		}
	}

	return false
}
