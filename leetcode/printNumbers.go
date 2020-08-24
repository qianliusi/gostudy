package leetcode

import "math"

func printNumbers(n int) []int {
	m := int(math.Pow10(n))
	res := make([]int, 0)
	for i := 1; i < m; i++ {
		res = append(res, i)
	}
	return res
}
