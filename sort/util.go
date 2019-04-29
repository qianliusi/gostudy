package sort

import (
	"math/rand"
	"time"
)

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func genRandomSlice(len int) []int {
	var ret = make([]int, len, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		ret[i] = rand.Intn(10000)
	}
	return ret
}

func genRandomSliceLimit(len, max int) []int {
	var ret = make([]int, len, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		ret[i] = rand.Intn(max)
	}
	return ret
}

func genRandomArray() [10]int {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10000)
	}
	return arr
}
