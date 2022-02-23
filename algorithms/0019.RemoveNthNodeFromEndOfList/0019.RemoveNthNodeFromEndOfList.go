// Source : https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Author : Zhonghuan Gao
// Date   : 2022-02-23

/**********************************************************************************
*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]

Constraints:
The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

Follow up: Could you do this in one pass?
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
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	// 快指针先走
	fast := dummyHead
	for n > 0 {
		fast = fast.Next
		n--
	}

	// 快慢指针同时走
	slow := dummyHead
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 删除结点
	slow.Next = slow.Next.Next

	return dummyHead.Next
}
