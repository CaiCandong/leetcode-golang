package problem0114

import "kit"

type TreeNode = kit.TreeNode

//二叉树根据中序遍历使用右指针域展开为链表
// res 指向叶子节点,即链表的最后一个元素
// root 指向链表的头节点,第一个元素
// 获取到了链表的头部和尾部指针,从而进行链表合并操作
//func flatten(root*TreeNode){
//	if root==nil{
//		return
//	}
//	helper(root)
//}
//func helper(root *TreeNode)*TreeNode{
//	switch  {
//	//情况1:左右指针域都为空,即当前节点为叶子节点
//	//直接返回
//	case root.Left==nil&&root.Right==nil:
//		return root
//	//情况2:左节点为空
//	//直接进入下一层
//	case root.Left==nil:
//		return helper(root.Right)
//	//情况3:右节点为空
//	//将左节点转移到右节点上,进入下一层
//	case root.Right==nil:
//		root.Right=root.Left
//		root.Left=nil
//		return helper(root.Right)
//	//情况4:左右节点都不为空
//	//step1:处理右节点,处理完毕是右链表
//	//step2:处理左节点,处理完毕是左链表
//	//step3:将右节点连接到左节点后
//	//step4:将
//	default:
//		res:=helper(root.Right)
//		helper(root.Left).Right=root.Right
//		root.Right=root.Left
//		root.Left=nil
//		return res
//	}
//}
/*
	解法2:非递归
*/
func flatten(root *TreeNode) {
	for root != nil {
		switch {
		case root.Right == nil && root.Left != nil:
			root.Right = root.Left
			root.Left = nil
		case root.Left != nil && root.Right != nil:
			last := root.Left
			for last.Right != nil {
				last = last.Right
			}
			last.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}

var pre *TreeNode

func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	pre = nil
	recur(root)
}

func recur(root *TreeNode) {
	if root == nil {
		return
	}

	recur(root.Right)
	recur(root.Left)
	root.Right = pre
	root.Left = nil
	pre = root
}
