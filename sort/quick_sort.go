package sort

func quickSort(d []int) {
	l := 0
	r := len(d) - 1
	if l < r {
		pi := partition(d[l : r+1])
		quickSort(d[l:pi])
		quickSort(d[pi+1 : r+1])
	}
}

func partition(d []int) int {
	p := 0
	index := p + 1
	for i := index; i < len(d); i++ {
		if d[i] < d[p] {
			swap(d, i, index)
			index += 1
		}
	}
	swap(d, p, index-1)
	return index - 1
}
