package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	next := pre.Next
	for next != nil {
		if next.Val == val {
			pre.Next = next.Next
			break
		}
		pre = pre.Next
		next = pre.Next
	}
	return dummy.Next
}
