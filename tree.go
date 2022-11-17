package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的深度
func treeMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeMaxDepth(root.Left) + 1
	right := treeMaxDepth(root.Right) + 1
	return int(math.Max(float64(left), float64(right)))
}

// TODO 判断是否平衡二叉树
func isBalanced(root *TreeNode) bool {
	return false
}

// n个骰子
func twoSum(n int) []float64 {
	return nil
}
