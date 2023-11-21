package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	fmt.Println(hasPathSum(root, 22))

}
