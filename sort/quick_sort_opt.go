package sort

func quickSortOpt(d []int) {
	r := len(d)
	if r < 10 {
		insertionSort(d)
		return
	}
	if r > 1 {
		pi := part(d[:r])
		quickSort(d[:pi])
		quickSort(d[pi+1 : r])
	}

}

func insertionSort(arr []int) {
	for i := range arr {
		pre := i - 1
		current := arr[i]
		for pre >= 0 && arr[pre] > current {
			arr[pre+1] = arr[pre]
			pre -= 1
		}
		arr[pre+1] = current
	}
}

func part(d []int) int {
	selectMidPivot(d)
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

func selectMidPivot(d []int) {
	length := len(d)
	mid := length >> 1
	high := length - 1
	//使用三数取中法选择枢轴
	if d[mid] > d[high] {
		swap(d, mid, high)
	}
	if d[0] > d[high] {
		swap(d, 0, high)
	}
	if d[mid] > d[0] {
		swap(d, mid, 0)
	}
}
