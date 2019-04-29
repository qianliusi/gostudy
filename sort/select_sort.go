package sort

func selectSort(d []int) {
	length := len(d)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if d[minIndex] > d[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			d[i], d[minIndex] = d[minIndex], d[i]
		}
	}
}
