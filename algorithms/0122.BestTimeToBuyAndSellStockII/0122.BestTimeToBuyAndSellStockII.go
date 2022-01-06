// Source : https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
// Author : Zhonghuan Gao
// Date   : 2020-02-27

/**********************************************************************************
You are given an array prices for which the ith element is the price of a given stock on day i.
Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.

Example 2:
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.

Example 3:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e., max profit = 0.

Constraints:
1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104
*
***********************************************************************************/

package main

/**
解法一：动态规划
思路：定义dp[i] 为第 i 天的「累计最大收益」，那么有下面两种状态
- 目前不持有任何股票，对应的「累计最大收益」记为 dp[i][0]
- 目前持有一支股票，对应的「累计最大收益」记为 dp[i][1]

状态转移方程:
	dp[i][0] = max(dp[i−1][0], dp[i−1][1] + prices[i])
	dp[i][1] = max(dp[i−1][1], dp[i−1][0] - prices[i])
时间复杂度: O(n)
时间复杂度: O(n)
Ref: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode-s/
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

/**
解法一的空间优化
*/
func maxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}

/**
解法二：贪心
*/
func maxProfit3(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		maxProfit += max(0, prices[i]-prices[i-1])
	}
	return maxProfit
}
