package leetcode

import (
	"fmt"
	"sort"
	"strconv"
)

func singleNumber(nums []int) int {
	s := 0
	for _, v := range nums {
		s ^= v
	}
	return s
}

func majorityElement(nums []int) int {
	c := 0
	m := 0
	for k, v := range nums {
		if v == nums[m] {
			c++
		} else {
			c--
			if c == 0 {
				m = k + 1
			}
		}
	}
	return nums[m]
}

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); {
		for j := 0; j < len(matrix[i]); {
			if target == matrix[i][j] {
				return true
			} else if target > matrix[i][j] {
				j++
			} else {
				i++
				break
			}
		}
	}
	return false
}

func twoSum(nums []int, target int) []int {
	var result []int
	itemMap := make(map[int]int, len(nums))
	for i, e := range nums {
		itemMap[e] = i
	}
	for i, e := range nums {
		if j, ok := itemMap[target-e]; ok && i != j {
			result = append(append(result, i), j)
			return result
		}
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (c *ListNode) String() string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf("%v -> %v", c.Val, c.Next)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	flag := false
	first := true
	var head *ListNode
	var last *ListNode
	for l1 != nil && l2 != nil {
		if first {
			first = false
			last = &ListNode{}
			head = last
			last.Val = l1.Val + l2.Val
			if last.Val > 9 {
				last.Val = last.Val % 10
				flag = true
			}
		} else {
			p := &ListNode{}
			last.Next = p
			p.Val = l1.Val + l2.Val
			if flag {
				p.Val++
			}
			if p.Val > 9 {
				p.Val = p.Val % 10
				flag = true
			} else {
				flag = false
			}
			last = p
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil || l2 != nil {
		var l *ListNode
		if l1 != nil {
			l = l1
		} else {
			l = l2
		}
		for l != nil {
			p := &ListNode{}
			last.Next = p
			p.Val = l.Val
			if flag {
				p.Val++
			}
			if p.Val > 9 {
				p.Val = p.Val % 10
				flag = true
			} else {
				flag = false
			}
			last = p
			l = l.Next
		}
	}
	if flag {
		p := &ListNode{1, nil}
		last.Next = p
	}
	return head
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	m := make(map[rune]int)
	left := 0
	maxLen := 0
	for i, c := range s {
		if lastPosition, ok := m[c]; ok && lastPosition >= left {
			left = lastPosition + 1
		} else {
			currentLen := i - left + 1
			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
		m[c] = i
	}
	return maxLen
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}

func longestPalindrome(s string) string {
	odd := longestPalindromeOdd(s)
	even := longestPalindromeEven(s)
	if len(odd) > len(even) {
		return odd
	}
	return even
}

func longestPalindromeOdd(s string) string {
	if s == "" {
		return ""
	}
	slen := len(s)
	maxlen := 0
	index := 0
	for i := range s {
		for length := 1; i-length >= 0 && (i+length) < slen; length++ {
			if s[i-length] == s[i+length] {
				if length > maxlen {
					maxlen = length
					index = i
				}
			} else {
				break
			}
		}
	}
	return s[index-maxlen : index+maxlen+1]
}

func longestPalindromeEven(s string) string {
	if s == "" {
		return ""
	}
	slen := len(s)
	if slen == 1 {
		return s
	}
	maxlen := 0
	index := -1
	oneIndex := -1

	for i := range s {
		if i+1 < slen && s[i] == s[i+1] {
			oneIndex = i
			for length := 1; i-length >= 0 && (i+1+length) < slen; length++ {
				if s[i-length] == s[i+1+length] {
					if length > maxlen {
						maxlen = length
						index = i
					}
				} else {
					break
				}
			}
		}
	}
	if index == -1 {
		index = oneIndex
	}
	if index == -1 {
		return ""
	}
	return s[index-maxlen : index+1+maxlen+1]
}

func isMatch2(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	first := s != "" && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first && isMatch(s[1:], p))
	} else {
		return first && isMatch(s[1:], p[1:])
	}
}
func isMatch(s string, p string) bool {
	memo := make(map[string]bool, 100)
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		key := strconv.Itoa(i) + "_" + strconv.Itoa(j)
		if v, ok := memo[key]; ok {
			return v
		}
		if j == len(p) {
			return i == len(s)
		}
		first := i < len(s) && (p[j] == s[i] || p[j] == '.')
		ans := false
		if j <= len(p)-2 && p[j+1] == '*' {
			ans = dp(i, j+2) || (first && dp(i+1, j))
		} else {
			ans = first && dp(i+1, j+1)
		}
		memo[key] = ans
		return ans
	}
	return dp(0, 0)
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0
	for l < r {
		width := r - l
		area := 0
		if height[l] < height[r] {
			area = width * height[l]
			if area > max {
				max = area
			}
			l++
		} else {
			area = width * height[r]
			if area > max {
				max = area
			}
			r--
		}
	}
	return max
}

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}
	var result [][]int
	for i, a := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := length - 1
		for k > j {
			b := nums[j]
			c := nums[k]
			sum := a + b + c
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				item := []int{a, b, c}
				result = append(result, item)
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			}
		}
	}
	return result
}

func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	s := m[digits[0]]
	for i := 1; i < len(digits); i++ {
		temp := make([]string, 0)
		for j := 0; j < len(s); j++ {
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, s[j]+m[digits[i]][k])
			}
		}
		s = temp
	}
	return s
}

func letterCombinations(digits string) []string {
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var ans []string
	var dfs func(com, next string)
	dfs = func(com, next string) {
		d := next[:1]
		l := m[d]
		for i := 0; i < len(l); i++ {
			le := l[i : i+1]
			cl := com + le
			if len(next[1:]) == 0 {
				ans = append(ans, cl)
			} else {
				dfs(cl, next[1:])
			}
		}
	}
	if len(digits) != 0 {
		dfs("", digits)
	}
	return ans
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	first := dummy
	second := dummy
	for i := 0; first != nil; i++ {
		first = first.Next
		if i > n {
			second = second.Next
		}
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			if stack[len(stack)-1] == m[s[i]] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	prev := prehead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return prehead.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length < 2 {
		return lists[0]
	}
	m := length / 2
	return mergeTwoLists(mergeKLists(lists[:m]), mergeKLists(lists[m:]))
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	var mergeListNode *ListNode
	for i := 0; i < len(lists); i++ {
		mergeListNode = mergeTwoLists(mergeListNode, lists[i])
	}
	return mergeListNode
}

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(left, right int, out string)
	dfs = func(left, right int, out string) {
		if left > right {
			return
		}
		if left == 0 && right == 0 {
			ans = append(ans, out)
		} else {
			if left > 0 {
				dfs(left-1, right, out+"(")
			}
			if right > 0 {
				dfs(left, right-1, out+")")
			}
		}
	}
	dfs(n, n, "")
	return ans
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func nextPermutation(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	firstDescend := -1
	big := -1
	for i := length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			firstDescend = i
			break
		}
	}
	if firstDescend > -1 {
		for i := firstDescend + 1; i < length; i++ {
			if nums[firstDescend] < nums[i] {
				big = i
			} else {
				break
			}
		}
		if big > -1 {
			nums[firstDescend], nums[big] = nums[big], nums[firstDescend]
		}
	}
	reverse(nums[firstDescend+1:])
}

func longestValidParentheses(s string) int {
	ans := 0
	var stack = &stack{}
	stack.push(-1)
	for i, v := range s {
		if v == '(' {
			stack.push(i)
		} else {
			stack.pop()
			if stack.empty() {
				stack.push(i)
			} else {
				maxLen := i - stack.peek().(int)
				if maxLen > ans {
					ans = maxLen
				}
			}
		}

	}
	return ans
}

type stack struct {
	elements []interface{}
}

func (s *stack) push(e interface{}) {
	s.elements = append(s.elements, e)
}
func (s *stack) pop() interface{} {
	if s.empty() {
		return nil
	}
	p := s.elements[len(s.elements)-1]
	s.elements = s.elements[0 : len(s.elements)-1]
	return p
}
func (s *stack) peek() interface{} {
	if s.empty() {
		return nil
	}
	return s.elements[len(s.elements)-1]
}
func (s *stack) empty() bool {
	return len(s.elements) == 0
}

func search(nums []int, target int) int {
	var rightVal = func(x, start int) int {
		if x < start {
			return x + 0x3f3f3f3f - start
		}
		return x
	}
	length := len(nums)
	start := 0
	end := length - 1
	for start <= end {
		middle := (start + end) >> 1
		if target == nums[middle] {
			return middle
		}
		if rightVal(target, nums[0]) > rightVal(nums[middle], nums[0]) {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}

func combinationSum(candidates []int, target int) [][]int {
	var comb [][]int
	for i := range candidates {
		if candidates[i] == target {
			comb = append(comb, []int{candidates[i]})
		} else if candidates[i] < target {
			rtn := combinationSum(candidates[i:], target-candidates[i])
			for j := range rtn {
				if len(rtn[j]) == 0 {
					continue
				}
				comb = append(comb, append([]int{candidates[i]}, rtn[j]...))
			}
		}
	}
	return comb
}
