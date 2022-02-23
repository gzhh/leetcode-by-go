// Source : https://leetcode.com/problems/swap-nodes-in-pairs/
// Author : Zhonghuan Gao
// Date   : 2022-02-23

/**********************************************************************************
*
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Example 1:
Input: head = [1,2,3,4]
Output: [2,1,4,3]

Example 2:
Input: head = []
Output: []

Example 3:
Input: head = [1]
Output: [1]

Constraints:
The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100
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

func swapPairs(head *ListNode) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		tmpNext := cur.Next
		tmpNextNextNext := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmpNext
		cur.Next.Next.Next = tmpNextNextNext
		cur = cur.Next.Next
	}
	return dummyHead.Next
}
