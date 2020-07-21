// Source : https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
// Author : Zhonghuan Gao
// Date   : 2020-07-21

/**********************************************************************************
*
Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.
Return the decimal value of the number in the linked list.

Example 1:
Input: head = [1,0,1]
Output: 5
Explanation: (101) in base 2 = (5) in base 10

Example 2:
Input: head = [0]
Output: 0

Example 3:
Input: head = [1]
Output: 1

Example 4:
Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
Output: 18880

Example 5:
Input: head = [0,0]
Output: 0

Constraints:
The Linked List is not empty.
Number of nodes will not exceed 30.
Each node's value is either 0 or 1.
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
先将链表反转，然后遍历遍历链表
num = cur.Val * 2^i
 */
func getDecimalValue(head *ListNode) int {
	num := 0
	// TODO
	return num
}

/**
解法二：位操作
 */
func getDecimalValue2(head *ListNode) int {
	num := 0
	for head != nil {
		num <<= 1
		num += head.Val
		head = head.Next
	}
	return num
}
