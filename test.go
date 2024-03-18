package main

import "fmt"

// 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 根据指定的数组 (切片) 生成对应的链表并返回头节点
func GenerateListNodesByArray(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for _, val := range nums {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	return dummy.Next
}

// 打印指定的链表
func Dump(head *ListNode) {
	cur := head
	// 直接使用 cur 遍历链表
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func Reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		curNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = curNext
	}
	return pre
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	myList := GenerateListNodesByArray(nums)

	myList = Reverse(myList)

	Dump(myList)
}
