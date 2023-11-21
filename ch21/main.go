package main

import (
	"fmt"
)

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 示例 1：

// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 示例 2：

// 输入：l1 = [], l2 = []
// 输出：[]
// 示例 3：

// 输入：l1 = [], l2 = [0]
// 输出：[0]

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := []int{-6, -5, 6, 6, 7}
	l2 := []int{0}

	list1 := &ListNode{
		l1[0], nil,
	}

	list2 := &ListNode{
		l2[0], nil,
	}

	for i := 1; i < len(l1); i++ {
		temp := list1
		for {
			if temp.Next == nil {
				break
			}
			temp = temp.Next
		}
		temp.Next = &ListNode{
			l1[i],
			nil,
		}
	}
	for i := 1; i < len(l2); i++ {
		temp := list2
		for {
			if temp.Next == nil {
				break
			}
			temp = temp.Next
		}
		temp.Next = &ListNode{
			l2[i],
			nil,
		}
	}
	// fmt.Println(list1)
	// fmt.Println(list2)

	// mergeTwoLists(list1, list2)
	lists := mergeTwoLists(&ListNode{}, &ListNode{})
	fmt.Println(lists)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := ListNode{}
	//链头
	head := &res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			head.Next = list2
			break
		}

		if list2 == nil {
			head.Next = list1
			break
		}

		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}

		head = head.Next
	}

	return res.Next
}
