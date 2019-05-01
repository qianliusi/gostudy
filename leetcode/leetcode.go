package leetcode

func singleNumber(nums []int) int {
	s := 0
	for _, v := range nums {
		s ^= v
	}
	return s
}

func majorityElement(nums []int) int {
	c := 0
	m := 0
	for k, v := range nums {
		if v == nums[m] {
			c++
		} else {
			c--
			if c == 0 {
				m = k + 1
			}
		}
	}
	return nums[m]
}
