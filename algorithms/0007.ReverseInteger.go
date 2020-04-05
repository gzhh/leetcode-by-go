// Source : https://leetcode.com/problems/reverse-integer/
// Author : Zhonghuan Gao
// Date   : 2020-04-05

/********************************************************************************** 
*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
Input: 123
Output: 321

Example 2:
Input: -123
Output: -321

Example 3:
Input: 120
Output: 21

Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*
***********************************************************************************/

/**
 * 解法一：
 */
func reverse(x int) int {
	const INT_MAX = 1 << 31 - 1
	const INT_MIN = -(1 << 31)
	res := 0
	for x != 0 {
		res = res * 10 + x % 10
		x /= 10
		if (res < INT_MIN || res > INT_MAX) {
			return 0
		}
	}
	return res
}
