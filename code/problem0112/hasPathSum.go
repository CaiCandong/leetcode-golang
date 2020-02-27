package problem0112

import "kit"

type TreeNode = kit.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSum3(root, sum)
}

//深度优先遍历
//非递归
//包裹链表
// 对节点增加数据项
func hasPathSum1(root *TreeNode, sum int) bool {
	type myNode struct {
		Node *TreeNode
		sum  int
	}
	if root == nil {
		return false
	}
	stack := make([]myNode, 0)
	stack = append(stack, myNode{root, root.Val})

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Node.Left != nil {
			stack = append(stack, myNode{node.Node.Left, node.sum + node.Node.Left.Val})
		}
		if node.Node.Right != nil {
			stack = append(stack, myNode{node.Node.Right, node.sum + node.Node.Right.Val})
		}
		if node.Node.Left == nil && node.Node.Right == nil && node.sum == sum {
			return true
		}
	}
	return false
}

// 解法2:同解法1
// 使用双栈
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	stack_node := make([]*TreeNode, 0)
	stack_node = append(stack_node, root)
	stack_sum := make([]int, 0)
	stack_sum = append(stack_sum, root.Val)
	for len(stack_node) != 0 {
		cur_node := stack_node[len(stack_node)-1]
		stack_node = stack_node[:len(stack_node)-1]
		cur_sum := stack_sum[len(stack_sum)-1]
		stack_sum = stack_sum[:len(stack_sum)-1]

		if cur_node.Left != nil {
			stack_node = append(stack_node, cur_node.Left)
			stack_sum = append(stack_sum, cur_sum+cur_node.Left.Val)
		}
		if cur_node.Right != nil {
			stack_node = append(stack_node, cur_node.Right)
			stack_sum = append(stack_sum, cur_sum+cur_node.Right.Val)
		}
		if cur_node.Left == nil && cur_node.Right == nil && cur_sum == sum {
			return true
		}
	}
	return false
}

//解法3:递归

func hasPathSum3(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return sum == root.Val
	}
	return hasPathSum3(root.Left, sum-root.Val) || hasPathSum3(root.Right, sum-root.Val)
}
