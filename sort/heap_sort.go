package sort

import "fmt"

func heapSort(d []int) {
	//p(d, "heapSort start")
	buildMaxHeap(d)
	length := len(d)
	for i := length - 1; i > 0; i-- {
		d[0], d[i] = d[i], d[0]
		heapAdjust(d, 0, i)
	}
}

func buildMaxHeap(d []int) {
	length := len(d)
	if length == 1 {
		return
	}
	first := length/2 - 1
	for i := first; i >= 0; i-- {
		heapAdjust(d, i, length)
	}
}
func heapAdjust(d []int, index int, end int) {
	for i := index; 2*i+2 < end; {
		left := 2*i + 1
		right := 2*i + 2
		maxInx := findMax(d, i, left, right)
		if maxInx == i {
			return
		}
		if maxInx == left {
			d[i], d[left] = d[left], d[i]
			//p(d, "line 60")
			i = left
		} else {
			d[i], d[right] = d[right], d[i]
			//p(d, "line 64")
			i = right
		}
	}

}
func p(d []int, s string) {
	fmt.Println(d, s)
}

func findMax(d []int, idx1 int, idx2 ...int) int {
	l := len(d)
	maxIdx := idx1
	for _, v := range idx2 {
		if v < l && d[v] > d[maxIdx] {
			maxIdx = v
		}
	}
	return maxIdx
}

/*func heapSort(arr []int) []int {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}*/
