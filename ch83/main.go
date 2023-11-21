// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。

// 示例 1：

// 输入：head = [1,1,2]
// 输出：[1,2]
// 示例 2：

// 输入：head = [1,1,2,3,3]
// 输出：[1,2,3]

// 提示：

// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序 排列
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	temp := head

	for {
		if temp == nil || temp.Next == nil {
			break
		}
		if temp.Val == temp.Next.Val {
			//跳过下一个node直接指向下下个node
			temp.Next = temp.Next.Next
		} else {
			//跳向下一个node
			temp = temp.Next
		}
	}
	return head
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{},
	}
	head.Next = &ListNode{
		Val:  1,
		Next: &ListNode{},
	}
	head.Next.Next = &ListNode{
		Val:  2,
		Next: &ListNode{},
	}
	head.Next.Next.Next = &ListNode{
		Val:  3,
		Next: &ListNode{},
	}
	head.Next.Next.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	// for head != nil {
	// 	fmt.Println(head.val)
	// 	head = head.Next
	// }
	res := deleteDuplicates(head)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
