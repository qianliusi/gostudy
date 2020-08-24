package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	result := make([]int, 0)
	for visit := head; visit != nil; visit = visit.Next {
		result = append(result, visit.Val)
	}
	length := len(result)
	for i := 0; i < length-i-1; i++ {
		result[i], result[length-i-1] = result[length-i-1], result[i]
	}
	return result
}
