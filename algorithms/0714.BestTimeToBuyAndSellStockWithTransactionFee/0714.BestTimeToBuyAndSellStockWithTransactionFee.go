// Source : https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// Author : Zhonghuan Gao
// Date   : 2020-02-27

/**********************************************************************************
Your are given an array of integers prices, for which the i-th element is the price of a given stock on day i; and a non-negative integer fee representing a transaction fee.
You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction. You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)
Return the maximum profit you can make.

Example 1:
Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
Output: 8
Explanation: The maximum profit can be achieved by:
- Buying at prices[0] = 1
- Selling at prices[3] = 8
- Buying at prices[4] = 4
- Selling at prices[5] = 9
- The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.

Note:
0 < prices.length <= 50000.
0 < prices[i] < 50000.
0 <= fee < 50000.
*
***********************************************************************************/

package main

/**
解法一：动态规划
思路：
	参考 122.Best Time to Buy and Sell Stock II (任意次交易)
	https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
状态转移方程:
	// i:0-n j:0-k
	dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i] - fee)
	dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
时间复杂度: O(k*n)
时间复杂度: O(k)
Ref: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-han-sh-rzlz/
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

/**
解法二：贪心
Ref: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-han-sh-rzlz/
*/
func maxProfit2(prices []int, fee int) int {
	n := len(prices)
	maxProfit := 0
	buy := prices[0] + fee
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			maxProfit += prices[i] - buy
			buy = prices[i]
		}
	}
	return maxProfit
}
