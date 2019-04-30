package sort

func quickSort(d []int) {
	r := len(d)
	if r > 1 {
		pi := partition(d[:r])
		quickSort(d[:pi])
		quickSort(d[pi+1 : r])
	}
}

func partition(d []int) int {
	p := 0
	index := p + 1
	for i := index; i < len(d); i++ {
		if d[i] < d[p] {
			swap(d, i, index)
			index++
		}
	}
	swap(d, p, index-1)
	return index - 1
}
