package problem0094

import (
	"kit"
)

/*
	二叉树的中序遍历
*/
type TreeNode = kit.TreeNode

//输入二叉树的根节点
//输出二叉树的中序遍历数组 左根右
/*
	解法1:递归
*/
func InorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 10)
	helper(root, res)
	return res
}
func helper(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = helper(root.Left, res)
	res = append(res, root.Val)
	res = helper(root.Right, res)
	return res
}
