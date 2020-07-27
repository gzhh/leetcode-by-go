// Source : https://leetcode.com/problems/reverse-linked-list/
// Author : Zhonghuan Gao
// Date   : 2020-07-12

/**********************************************************************************
*
Given a singly linked list, determine if it is a palindrome.

Example 1:
Input: 1->2
Output: false

Example 2:
Input: 1->2->2->1
Output: true

Follow up:
Could you do it in O(n) time and O(1) space?
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
解法一：
使用快慢指针找到链表的中点，然后按顺序将链表后半段节点的值存入数组s，
然后从后向前遍历数组s，并从前向后遍历原链表，若链表当前节点的值不等于
数组当前的值，则不是回文链表，返回false；否则是回文链表，最后返回true。
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var s []int
	for slow != nil {
		s = append(s, slow.Val)
		slow = slow.Next
	}

	len := len(s)
	cur := head
	for i := len - 1; i >= 0; i-- {
		if cur.Val != s[i] {
			return false
		}
		cur = cur.Next
	}
	return true
}

/**
解法二：
将链表转化成数组，利用数组判断是否是回文
 */
func isPalindrome2(head *ListNode) bool {
	cur := head
	var s []int
	for cur != nil {
		s = append(s, cur.Val)
		cur = cur.Next
	}

	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
