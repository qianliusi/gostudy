package sort

import (
	"sort"
	"testing"
)

func TestGolangSort1(t *testing.T) {
	a := genRandomSliceLimit(100000000, 1000000000)
	sort.Ints(a)
	//fmt.Println(a)
}
