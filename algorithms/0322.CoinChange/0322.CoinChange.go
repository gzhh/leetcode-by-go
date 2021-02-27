// Source : https://leetcode.com/problems/coin-change/
// Author : Zhonghuan Gao
// Date   : 2020-02-21

/**********************************************************************************
*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
You may assume that you have an infinite number of each kind of coin.

Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1

Example 3:
Input: coins = [1], amount = 0
Output: 0

Example 4:
Input: coins = [1], amount = 1
Output: 1

Example 5:
Input: coins = [1], amount = 2
Output: 2

Constraints:
1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104
*
***********************************************************************************/

package main

/**
解法一：动态规划（背包问题）
思路：
	定义 m=amount,n为币值的种数，dp[i] 表示钱i所用的最少钱币个数，coins[j] 代表的是第 j 枚硬币的面值，即我们枚举最后一枚硬币面额是 coins[j]
状态转移方程：
	dp[i] = min(dp[i], dp[i - coins[j]] + 1) {i:0~m+1 j:0~n}
时间复杂度：O(m*n)
空间复杂度：O(m*n)
Ref: https://leetcode-cn.com/problems/coin-change/solution/322-ling-qian-dui-huan-by-leetcode-solution/
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i - coins[j]] + 1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}