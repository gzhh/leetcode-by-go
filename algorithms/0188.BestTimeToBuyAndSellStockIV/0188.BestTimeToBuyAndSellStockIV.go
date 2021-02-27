// Source : https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// Author : Zhonghuan Gao
// Date   : 2020-02-27

/**********************************************************************************
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
Design an algorithm to find the maximum profit. You may complete at most k transactions.
Notice that you may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Example 1:
Input: k = 2, prices = [2,4,1]
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.

Example 2:
Input: k = 2, prices = [3,2,6,5,0,3]
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4. Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.

Constraints:
0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000
*
***********************************************************************************/

package main

import "math"

/**
解法一：动态规划
思路：
	参考 123.Best Time to Buy and Sell Stock III (最多进行两次交易)
	https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
状态转移方程:
	// i:0-n j:0-k
	buy[j] = max(buy[j], sell[j] - prices[i])
	sell[j] = max(sell[j], buy[j-1] + prices[i])
时间复杂度: O(k*n)
时间复杂度: O(k)
Ref: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-iv-by-8xtkp/
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	k = min(k, n / 2)

	buy := make([]int, k+1)
	sell := make([]int, k+1)
	buy[0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt64 / 2
		sell[i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], sell[0] - prices[i])
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j] - prices[i])
			sell[j] = max(sell[j], buy[j-1] + prices[i])
		}
	}
	return max(sell...)
}