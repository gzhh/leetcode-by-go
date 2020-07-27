// Source : https://leetcode.com/problems/reverse-linked-list/
// Author : Zhonghuan Gao
// Date   : 2020-07-12

/**********************************************************************************
*
Reverse a singly linked list.

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

Follow up:
A linked list can be reversed either iteratively or recursively. Could you implement both?
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
解法一：递归
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

/**
解法二：迭代
 */
func reverseList2(head *ListNode) *ListNode {
	var prev, cur, tmp *ListNode
	cur = head
	for cur != nil {
		tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}
