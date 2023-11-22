// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

// 不允许修改 链表。

// 示例 1：

// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。
// 示例 2：

// 输入：head = [1,2], pos = 0
// 输出：返回索引为 0 的链表节点
// 解释：链表中有一个环，其尾部连接到第一个节点。
// 示例 3：

// 输入：head = [1], pos = -1
// 输出：返回 null
// 解释：链表中没有环。

// 提示：

// 链表中节点的数目范围在范围 [0, 104] 内
// -105 <= Node.val <= 105
// pos 的值为 -1 或者链表中的一个有效索引

// 进阶：你是否可以使用 O(1) 空间解决此题？

// 快慢指针
// 思路与算法

// 我们使用两个指针，fast\textit{fast}fast 与 slow\textit{slow}slow。它们起始都位于链表的头部。随后，slow 指针每次向后移动一个位置，而 fast 指针向后移动两个位置。如果链表中存在环，则 fast 指针最终将再次与 slow 指针在环中相遇。

// 如下图所示，设链表中环外部分的长度为 a。slow 指针进入环后，又走了 bbb 的距离与 fast 相遇。此时，fast 指针已经走完了环的 n 圈，因此它走过的总距离为 a+n(b+c)+b=a+(n+1)b+nca+n(b+c)+b=a+(n+1)b+nca+n(b+c)+b=a+(n+1)b+nc。

// 根据题意，任意时刻，ast 指针走过的距离都为 slow 指针的 2 倍。因此，我们有

// a+(n+1)b+nc=2(a+b)  ⟹  a=c+(n−1)(b+c)a+(n+1)b+nc=2(a+b)  a=c+(n-1)(b+c)
// a+(n+1)b+nc=2(a+b)⟹a=c+(n−1)(b+c)
// 有了 a=c+(n−1)(b+c)a=c+(n-1)(b+c)a=c+(n−1)(b+c) 的等量关系，我们会发现：从相遇点到入环点的距离加上 n−1n-1n−1 圈的环长，恰好等于从链表头部到入环点的距离。

// 这个圈数n好像就是1，不会是其他的数，因此可以得出a == c
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func main() {
	head := new(ListNode)
	node1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	head = &ListNode{
		Val:  3,
		Next: node1,
	}
	node2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	node1.Next = node2

	node3 := &ListNode{
		Val:  -4,
		Next: nil,
	}
	node2.Next = node3
	node3.Next = node1

	fmt.Println(*head)
	fmt.Println(detectCycle(head))
}
