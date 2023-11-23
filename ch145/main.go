// 145. 二叉树的后序遍历
// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	arr := []int{}
	var raversal func(*TreeNode)
	raversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		raversal(node.Right)
		raversal(node.Left)
		arr = append(arr, node.Val)
	}
	raversal(root)
	return arr
}

func main() {
	node := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(postorderTraversal(node))
}
