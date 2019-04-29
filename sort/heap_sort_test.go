package sort

import (
	"testing"
)

func TestHeapSort1(t *testing.T) {
	a := genRandomSliceLimit(10000000, 100000000)
	heapSort(a)
	//fmt.Println(a)
}

func BenchmarkHeapSort1(b *testing.B) {
	b.StopTimer() // 暂停计时器
	a := genRandomSlice(1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sa := make([]int, len(a))
		copy(sa, a)
		heapSort(sa)
	}
}
