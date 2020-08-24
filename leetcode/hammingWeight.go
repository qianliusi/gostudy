package leetcode

func hammingWeight(num uint32) int {
	var res uint32 = 0
	for num > 0 {
		res += num & 1
		num >>= 1
	}
	return int(res)
}
