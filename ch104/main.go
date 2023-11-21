package main

import "fmt"

//知识点：递归 & 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {
	head := &TreeNode{
		Val:   1,
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}
	head.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: &TreeNode{},
	}
	head.Right = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: &TreeNode{},
	}
	head.Left.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	head.Right.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(maxDepth(head))
}
