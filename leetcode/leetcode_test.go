package leetcode

import "testing"

func TestSingleNumber1(t *testing.T) {
	a := []int{4, 1, 2, 1, 2}
	n := singleNumber(a)
	if n != 4 {
		t.Fail()
	}
}
