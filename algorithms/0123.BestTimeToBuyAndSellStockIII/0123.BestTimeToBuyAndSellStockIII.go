// Source : https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
// Author : Zhonghuan Gao
// Date   : 2020-02-27

/**********************************************************************************
Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete at most two transactions.
Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:
Input: prices = [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.

Example 2:
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.

Example 3:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.

Example 4:
Input: prices = [1]
Output: 0

Constraints:
1 <= prices.length <= 105
0 <= prices[i] <= 105
*
***********************************************************************************/

package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
解法一：动态规划
思路：由于我们最多可以完成两笔交易，因此在任意一天结束之后，我们会处于以下五个状态中的一种：
	- 未进行过任何操作；
	- 只进行过一次买操作；
	- 进行了一次买操作和一次卖操作，即完成了一笔交易；
	- 在完成了一笔交易的前提下，进行了第二次买操作；
	- 完成了全部两笔交易。
状态转移方程:
	dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
	dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
	dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
时间复杂度: O(n)
时间复杂度: O(n*5)
Ref: https://programmercarl.com/0123.%E4%B9%B0%E5%8D%96%E8%82%A1%E7%A5%A8%E7%9A%84%E6%9C%80%E4%BD%B3%E6%97%B6%E6%9C%BAIII.html
*/
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][5]int, n)

	dp[0][1], dp[0][2], dp[0][3], dp[0][4] = -prices[0], 0, -prices[0], 0
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[n-1][4]
}

/**
解法二：动态规划（解法一的空间优化）
思路：由于我们最多可以完成两笔交易，因此在任意一天结束之后，我们会处于以下五个状态中的一种：
	- 未进行过任何操作；
	- 只进行过一次买操作；
	- 进行了一次买操作和一次卖操作，即完成了一笔交易；
	- 在完成了一笔交易的前提下，进行了第二次买操作；
	- 完成了全部两笔交易。
状态转移方程:
	buy1 = max(buy1, -prices[i])
	sell1 = max(sell1, buy1 + prices[i])
	buy2 = max(buy2, sell1 - prices[i])
	sell2 = max(sell2, buy2 + prices[i])
时间复杂度: O(n)
时间复杂度: O(1)
Ref: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-iii-by-wrnt/
*/
func maxProfit1(prices []int) int {
	buy1, sell1, buy2, sell2 := -prices[0], 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}
