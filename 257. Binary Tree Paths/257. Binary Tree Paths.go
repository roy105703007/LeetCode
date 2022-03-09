/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	dfsTree(root, strconv.Itoa((*root).Val), &res)
	return res

}

func dfsTree(root *TreeNode, str string, res *[]string) {
	if (*root).Left == nil && (*root).Right == nil {
		*res = append(*res, str)
		return
	}
	if (*root).Left != nil {
		dfsTree((*root).Left, str+"->"+strconv.Itoa((*root).Left.Val), res)
	}
	if (*root).Right != nil {
		dfsTree((*root).Right, str+"->"+strconv.Itoa((*root).Right.Val), res)
	}
}