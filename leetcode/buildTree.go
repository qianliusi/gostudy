package leetcode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func indexOf(a []int, x int) int {
	for k, v := range a {
		if v == x {
			return k
		}
	}
	return -1
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := preorder[0]
	node := &TreeNode{root, nil, nil}
	if len(preorder) == 1 {
		return node
	}
	idx := indexOf(inorder, root)
	node.Left = BuildTree(preorder[1:idx+1], inorder[:idx])
	node.Right = BuildTree(preorder[idx+1:], inorder[idx+1:])
	return node
}
