// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：

// L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为：

// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	arr := []*ListNode{}
	node := head

	for node != nil {
		arr = append(arr, node)
		node = node.Next
	}

	l, h := 0, len(arr)-1
	for l < h {
		arr[l].Next = arr[h]
		l++
		if l == h {
			break
		}
		arr[h].Next = arr[l]
		h--
	}
	arr[l].Next = nil
}
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	reorderList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
