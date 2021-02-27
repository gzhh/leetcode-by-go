// Source : https://leetcode.com/problems/perfect-squares/
// Author : Zhonghuan Gao
// Date   : 2020-02-21

/**********************************************************************************
*
Given an integer n, return the least number of perfect square numbers that sum to n.
A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.

Example 1:
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.

Example 2:
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.

Constraints:
1 <= n <= 104
*
***********************************************************************************/

package main

/**
解法一：动态规划（背包问题）
思路：
	参考 322.CoinChange: https://leetcode.com/problems/coin-change/
时间复杂度：O(m*n)
空间复杂度：O(m*n)
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	coins := make([]int, n)
	j := 0
	for i := 1; i <= n; i++ {
		if i * i <= n {
			coins[j] = i * i
			j++
		}
	}

	dp := make([]int, n + 1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i - coins[j]] + 1)
			}
		}
	}

	return dp[n]
}

/**
解法一优化
Ref: https://leetcode-cn.com/problems/perfect-squares/solution/279-wan-quan-ping-fang-shu-wan-quan-bei-htyas/
 */
func numSquares2(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	for i := 0; i <= n; i++ {
		for j := 1; j * j <= i; j++ {
			dp[i] = min(dp[i], dp[i - j * j] + 1)
		}
	}

	return dp[n]
}