package leetcode

func Fib(n int) int {
	if n < 2 {
		return n
	}
	n0 := 0
	n1 := 1
	temp := 0
	for i := 2; i <= n; i++ {
		temp = n1
		n1 = (n1 + n0) % (1e9 + 7)
		n0 = temp
	}
	return n1
}
