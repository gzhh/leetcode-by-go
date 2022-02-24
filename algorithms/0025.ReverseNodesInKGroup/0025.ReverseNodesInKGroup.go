// Source : https://leetcode.com/problems/reverse-nodes-in-k-group/
// Author : Zhonghuan Gao
// Date   : 2022-02-24

/**********************************************************************************
**
Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
You may not alter the values in the list's nodes, only nodes themselves may be changed.

Example 1:
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]

Example 2:
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]

Constraints:
The number of nodes in the list is n.
1 <= k <= n <= 5000
0 <= Node.val <= 1000

Follow-up: Can you solve the problem in O(1) extra memory space?
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
解法一：递归
*/
func reverse(head, tail *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != tail {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	newHead := reverse(head, cur)
	head.Next = reverseKGroup(cur, k)
	return newHead
}

/**
解法二：迭代
参考：https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/wei-cha-fa-huo-zhan-de-si-lu-by-davis-we-xrcc/
*/
func reverseKGroup2(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	pre, tail := dummyHead, dummyHead

	for {
		count := k
		for count > 0 && tail != nil {
			tail = tail.Next
			count--
		}
		if tail == nil {
			break
		}

		newHead := pre.Next
		for pre.Next != tail {
			cur := pre.Next
			pre.Next = cur.Next
			cur.Next = tail.Next
			tail.Next = cur
		}
		pre, tail = newHead, newHead
	}
	return dummyHead.Next
}
