/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (70.39%)
 * Likes:    412
 * Dislikes: 0
 * Total Accepted:    112.2K
 * Total Submissions: 159.4K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * 输出: [1,3,2]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	//迭代算法
	stack:=make([]*TreeNode,0)
	res:=make([]int,0)
	current:=root
	for current!=nil||len(stack)!=0{
		for current!=nil{
			stack=append(stack,current)
			current=current.Left
		}
		current=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		res=append(res,current.Val)
		current=current.Right
	}
	return res
}
// @lc code=end

