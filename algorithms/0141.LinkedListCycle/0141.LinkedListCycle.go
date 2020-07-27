// Source : https://leetcode.com/problems/linked-list-cycle/
// Author : Zhonghuan Gao
// Date   : 2020-07-11

/********************************************************************************** 
*
Given a linked list, determine if it has a cycle in it.
To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
*
***********************************************************************************/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

/**
题意：
给一个链表，判断其是否有环

解法一：
1.使用两个指针，快指针和慢指针
2.慢指针一次往后遍历一个节点；快指针一次往后遍历两个节点
3.如果链表有环，那么快指针和慢指针必定在某点相遇
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var slow, fast *ListNode
	slow, fast = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}