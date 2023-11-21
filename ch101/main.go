// 给你一个二叉树的根节点 root ， 检查它是否轴对称。

//

// 示例 1：

// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
// 示例 2：

// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
//

// 提示：

// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//

// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	fmt.Println(root)
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
