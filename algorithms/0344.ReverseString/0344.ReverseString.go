// Source : https://leetcode.com/problems/reverse-string/
// Author : Zhonghuan Gao
// Date   : 2020-10-14

/**********************************************************************************
**
Write a function that reverses a string. The input string is given as an array of characters char[].
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
You may assume all the characters consist of printable ascii characters.

Example 1:
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
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

// two point solution 1
func reverseString1(s []byte)  {
	len := len(s)
	for i := 0; i < len / 2; i++ {
		s[i], s[len-i-1] = s[len-i-1], s[i]
	}
}

// two point solution 2
func reverseString2(s []byte)  {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}