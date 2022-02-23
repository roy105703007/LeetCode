/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := 0
	r := 0
	fmt.Println(root.Left)
	if root.Left != nil {
		l = maxDepth(root.Left)
	} else {
		l = 0
	}
	if root.Right != nil {
		r = maxDepth(root.Right)
	} else {
		r = 0
	}
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}