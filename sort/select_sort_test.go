package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort1(t *testing.T) {
	a := genRandomSlice(1000000)
	selectSort(a)
	fmt.Println(a)
}

func BenchmarkSelectSort1(b *testing.B) {
	b.StopTimer() // 暂停计时器
	a := genRandomSlice(1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sa := make([]int, len(a))
		copy(sa, a)
		selectSort(sa)
	}
}
