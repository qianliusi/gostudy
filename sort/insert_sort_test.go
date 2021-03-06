package sort

import (
	"testing"
)

func TestInsertSort1(t *testing.T) {
	a := genRandomSlice(100000)
	insertionSort(a)
	//fmt.Println(a)
}

func BenchmarkInsertSort1(b *testing.B) {
	b.StopTimer() // 暂停计时器
	a := genRandomSlice(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sa := make([]int, len(a))
		copy(sa, a)
		insertSort(sa)
	}
}
