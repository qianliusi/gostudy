package sort

import (
	"testing"
)

func TestMergeSort1(t *testing.T) {
	a := genRandomSliceLimit(100000000, 100000000)
	mergeSort(a)
	//fmt.Println(a)
}

func BenchmarkMergeSort1(b *testing.B) {
	b.StopTimer() // 暂停计时器
	a := genRandomSlice(1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sa := make([]int, len(a))
		copy(sa, a)
		mergeSort(sa)
	}
}
