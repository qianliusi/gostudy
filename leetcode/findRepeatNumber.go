package leetcode

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func findRepeatNumber(nums []int) int {
	for k, v := range nums {
		if v != k {
			if nums[v] == v {
				return v
			}
			swap(nums, k, v)
		}
	}
	return 0
}
