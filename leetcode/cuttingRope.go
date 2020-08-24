package leetcode

func CuttingRope1(n int) int {
	if n == 2 {
		return 1
	}
	res := 0
	for i := 2; i < n; i++ {
		res = Max(res, CuttingRope1(n-i)*i, (n-i)*i)
	}
	return res
}

func CuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	res := make([]int, n+1)
	res[2] = 1
	for i := 3; i <= n; i++ {
		for j := 2; j < i; j++ {
			res[i] = Max(res[i], res[i-j]*j, (i-j)*j)
		}
	}
	return res[n]
}

func Max(x int, y ...int) int {
	for _, v := range y {
		if x < v {
			x = v
		}
	}
	return x
}
