// Source : https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// Author : Zhonghuan Gao
// Date   : 2020-07-16

/**********************************************************************************
**
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:
Input: 1->1->2
Output: 1->2

Example 2:
Input: 1->1->2->3->3
Output: 1->2->3
**
**********************************************************************************/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
解法一：
循环链表，当下一个节点的值等于当前节点的值，删除下一个节点，继续循环，指向当前节点的指针不动；
当下一个节点的值不等于当前节点的值，将指向当前节点的指针指向下一个节点。
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		var tmp *ListNode
		if cur.Val == cur.Next.Val {
			tmp = cur.Next.Next
			cur.Next.Next = nil
			cur.Next = tmp
		} else {
			cur = cur.Next
		}
	}
	return head
}

/**
解法二：解法一的优化
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
