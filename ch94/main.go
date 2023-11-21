// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

//

// 示例 1：

// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
// 示例 2：

// 输入：root = []
// 输出：[]
// 示例 3：

// 输入：root = [1]
// 输出：[1]
//

// 提示：

// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//

// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

func main() {
	head := &TreeNode{
		Val:   5,
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}
	head.Left = &TreeNode{
		Val:   2,
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}
	head.Left.Left = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	head.Left.Right = &TreeNode{
		Val:   4,
		Left:  &TreeNode{},
		Right: nil,
	}
	head.Left.Right.Left = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	head.Right = &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(inorderTraversal(head))
}
