package leetcode

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	t := myPow(x, n/2)
	res := t * t
	if n&1 == 1 {
		res *= x
	}
	return res
}
