package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
【递归+移动指针】
对头节点等于移除值进行递归移除；
定义指向head的新节点node，对next的值与val比较，匹配到进行节点移除
*/
func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return head
	}

	// 链表头节点值等于val，则递归移除头节点
	if head.Val == val {
		node := head.Next
		head = nil
		return removeElements(node, val)
	}

	// 定义新节点指向head
	node := head
	for node != nil && node.Next != nil {
		if node.Next.Val == val {
			// next 节点值等于val，则移除 next 节点。
			node.Next = node.Next.Next
		} else {
			//移动节点指针
			node = node.Next
		}
	}
	return head
}
