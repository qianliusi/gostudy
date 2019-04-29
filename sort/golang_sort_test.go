package sort

import (
	"sort"
	"testing"
)

func TestGolangSort1(t *testing.T) {
	a := genRandomSliceLimit(10000000, 100000000)
	sort.Ints(a)
	//fmt.Println(a)
}
