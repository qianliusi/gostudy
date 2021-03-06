package sort

import (
	"testing"
)

func TestQuickSortOpt1(t *testing.T) {
	//a := genRandomSliceLimit(10, 20)
	a := genRandomSliceLimit(100000000, 100000000)
	quickSortOpt(a)
	//fmt.Println(a)
}

func BenchmarkQuickSortOpt1(b *testing.B) {
	b.StopTimer() // 暂停计时器
	a := genRandomSlice(1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sa := make([]int, len(a))
		copy(sa, a)
		quickSort(sa)
	}
}
