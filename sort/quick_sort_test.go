package sort

import (
	"testing"
)

func TestQuickSort1(t *testing.T) {
	//a := genRandomSliceLimit(10, 20)
	a := genRandomSliceLimit(100000000, 1000000000)
	quickSort(a)
	//fmt.Println(a)
}

func BenchmarkQuickSort1(b *testing.B) {
	b.StopTimer() // 暂停计时器
	a := genRandomSlice(1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sa := make([]int, len(a))
		copy(sa, a)
		quickSort(sa)
	}
}
