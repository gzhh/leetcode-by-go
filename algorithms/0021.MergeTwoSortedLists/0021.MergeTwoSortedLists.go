// Source : https://leetcode.com/problems/merge-two-sorted-lists/
// Author : Zhonghuan Gao
// Date   : 2020-07-17

/**********************************************************************************
**
Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
**
**********************************************************************************/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
解法一：
创建新的空节点，head指向空节点，且指针p为当前节点。判断l1当前节点的val是否小于l2当前节点的val，
若小于，将p指向l1，l1的当前节点往后移一位；否则，将p指向l2，l2的当前节点往后移一位。
循环直到两链表其中任一为空，跳出循环。若l1或l2的剩余节点不为空，则将p指向l1或l2。
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return head.Next
}

/**
解法二：递归
分析此题的特点满足：函数输出可以作为输入，进而更新输入，直到两个链表都为空。
*/
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}
