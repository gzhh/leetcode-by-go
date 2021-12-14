// Source : https://leetcode.com/problems/integer-break/
// Author : Zhonghuan Gao
// Date   : 2021-12-14

/**********************************************************************************
**
Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.
Return the maximum product you can get.

Example 1:
Input: n = 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.

Example 2:
Input: n = 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.

Constraints:
2 <= n <= 58
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[n]
}
