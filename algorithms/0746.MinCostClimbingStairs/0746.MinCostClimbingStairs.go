// Source : https://leetcode.com/problems/climbing-stairs/
// Author : Zhonghuan Gao
// Date   : 2020-01-01

/**********************************************************************************
*
On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).
Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

Example 1:
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.

Example 2:
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].

Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].
*
***********************************************************************************/

package main

/**
解法一：动态规划
思路：当前的最低花费为min(到达前两级阶梯的最低花费+爬前两级阶梯的花费, 到达前一级阶梯的最低花费+爬前一级阶梯的花费)
状态转移方程 dp[i] = max(prices[i] - minBuyPrice, maxProfit)
时间复杂度 O(n)
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] , dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-2] + cost[i-2], dp[i-1] + cost[i-1])
	}
	return dp[n]
}