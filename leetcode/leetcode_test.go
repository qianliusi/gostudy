package leetcode

import (
	"fmt"
	"testing"
)

func TestSingleNumber1(t *testing.T) {
	a := []int{4, 1, 2, 1, 2}
	n := singleNumber(a)
	if n != 4 {
		t.Fail()
	}
}

func TestMajorityElement1(t *testing.T) {
	a := []int{2, 2, 1, 1, 1, 2, 2}
	n := majorityElement(a)
	if n != 2 {
		t.Fail()
	}
}

func TestSearchMatrix1(t *testing.T) {
	a := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	n := searchMatrix(a, 20)
	if n != false {
		t.Fail()
	}
}
func TestTwoSum(t *testing.T) {
	a := []int{3, 2, 4}
	fmt.Println(twoSum(a, 6))
}
func TestAddTwoNumbers(t *testing.T) {
	first := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	end := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(addTwoNumbers(first, end))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("bb"))
}

func TestIsMatch(t *testing.T) {
	fmt.Println(isMatch("aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c"))

}
func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}

func TestLetterCombinations(t *testing.T) {
	letterCombinations("2345678923456789")
}

func TestRemoveNthFromEnd(t *testing.T) {
	first := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(removeNthFromEnd(first, 2))
}

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}
func TestMergeKLists(t *testing.T) {
	fmt.Println(mergeKLists(nil))
}

func TestNextPermutation(t *testing.T) {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
func TestLongestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses("()(())"))
}
func TestSearch(t *testing.T) {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
