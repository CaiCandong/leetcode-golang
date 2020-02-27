package problem0111

import (
	"kit"
)

type TreeNode = kit.TreeNode

/*
	解法1:广度优先搜索,自建结构体存储节点所在的深度
	使用切片来完成队列功能
*/

func minDepth1(root *TreeNode) int { //广度优先遍历
	type myNode struct {
		Node  *TreeNode
		leavl int
	}
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	queue := make([]*myNode, 1)
	queue[0] = &myNode{root, 1}
	//终止条件:队列为空
	for len(queue) != 0 {
		//出队列
		node := queue[0]
		queue = queue[1:]

		if node.Node.Left == nil && node.Node.Right == nil {
			return node.leavl
		}
		if node.Node.Left != nil {
			queue = append(queue, &myNode{node.Node.Left, node.leavl + 1})
		}
		if node.Node.Right != nil {
			queue = append(queue, &myNode{node.Node.Right, node.leavl + 1})
		}
	}
	return 0
}

/*
	解法二:使用递归实现
	自下而上记录深度
*/
//func minDepth(root *TreeNode) int {
//	switch {
//	case root == nil:
//		return 0
//	case root.Left == nil:
//		return minDepth(root.Right) + 1
//	case root.Right == nil:
//		return minDepth(root.Left) + 1
//	default:
//		return min(minDepth(root.Left), minDepth(root.Right)) + 1
//	}
//}
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
/*
	解法三:回溯法,深度优先搜索+剪枝
	切片完成栈
*/
func minDepth(root *TreeNode) int {
	type myNode struct {
		Node  *TreeNode
		leavl int
	}
	if root == nil {
		return 0
	}
	stack := make([]*myNode, 1)
	stack[0] = &myNode{root, 1}
	min := 1 << 63
	for len(stack) != 0 {
		//出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//进栈
		if node.Node.Left != nil && node.leavl < min {
			stack = append(stack, &myNode{node.Node.Left, node.leavl + 1})
		}
		if node.Node.Right != nil && node.leavl < min {
			stack = append(stack, &myNode{node.Node.Right, node.leavl + 1})
		}
		if node.Node.Left == nil && node.Node.Right == nil {
			min = node.leavl
		}
	}
	return min
}
