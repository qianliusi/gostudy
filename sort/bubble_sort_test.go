package sort

import (
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	a := genRandomArray()
	bubbleSort(a[:])
	if a[1] != 2 {
		t.Fail()
	}
}

func BenchmarkBubbleSort1(b *testing.B) {
	b.StopTimer() // 暂停计时器
	a := genRandomSlice(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sa := make([]int, len(a))
		copy(sa, a)
		bubbleSort(sa)
	}
}
