package sort

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	m := length / 2
	return merge(mergeSort(arr[:m]), mergeSort(arr[m:]))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}
