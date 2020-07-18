// Source : https://leetcode.com/problems/remove-linked-list-elements/
// Author : Zhonghuan Gao
// Date   : 2020-07-18

/**********************************************************************************
*
Remove all elements from a linked list of integers that have value val.

Example:
Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
*
***********************************************************************************/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

type ListNode struct {
	Val int
	Next *ListNode
}

/**
解法一：借用新链表
创建一个新链表，指针 newHead 指向新链表的头节点，指针 newCur 指向新链表的当前节点。
遍历原链表，如果节点的值不等于目标值 val，则将节点插入新的链表。最后将新链表的尾节点
指向空
*/
func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	newCur := newHead

	cur := head
	for cur != nil {
		if cur.Val != val {
			newCur.Next = cur
			newCur = newCur.Next
		}
		cur = cur.Next
	}
	newCur.Next = nil

	return newHead.Next
}

/**
解法二：操作原链表
遍历链表，若当前节点的下一个节点的值等于目标值 val，则删除下一个节点，
继续往后遍历，直到链表遍历完成。若第一个节点的值等于目标值 val，则返回
head.Next，否则直接返回 head
*/
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	cur := head

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	if head.Val == val {
		return head.Next
	}
	return head
}

/**
解法三：使用指针的指针操作原链表
 */
func removeElements3(head *ListNode, val int) *ListNode {
	for p := &head; *p != nil; {
		entry := *p
		if entry.Val == val {
			*p = entry.Next
		} else {
			p = &(*p).Next
		}
	}
	return head
}