// Source : https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// Author : Zhonghuan Gao
// Date   : 2020-02-27

/**********************************************************************************
*
Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

Example:
Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
*
***********************************************************************************/

package main

/**
解法一：动态规划
思路：定义dp[i] 为第 i 天的「累计最大收益」，那么有下面三种状态
- 目前持有一支股票，对应的「累计最大收益」记为 dp[i][0]
- 目前不持有任何股票，并且处于冷冻期中，对应的「累计最大收益」记为 dp[i][1]
- 目前不持有任何股票，并且不处于冷冻期中，对应的「累计最大收益」记为 dp[i][2]

状态转移方程:
	dpi][0] = max(dpi−1][0], dpi−1][2]−prices[i])
	dpi][1] = dpi−1][0] + prices[i]
	dpi][2] = max(dpi−1][1], dpi−1][2])
时间复杂度: O(n)
时间复杂度: O(n)
Ref: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/solution/zui-jia-mai-mai-gu-piao-shi-ji-han-leng-dong-qi-4/
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}