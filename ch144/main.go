// 144. 二叉树的前序遍历
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
// 示例 1：

// 输入：root = [1,null,2,3]
// 输出：[1,2,3]
// 示例 2：

// 输入：root = []
// 输出：[]
// 示例 3：

// 输入：root = [1]
// 输出：[1]
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode, arr []int) []int {
	if root == nil {
		return []int{}
	}
	res := arr
	res = append(arr, root.Val)
	fmt.Println(res)
	preorderTraversal(root.Left, res)
	preorderTraversal(root.Right, res)

	return res
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
	preorderTraversal(node, []int{})
}
