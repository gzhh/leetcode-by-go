// Source : https://leetcode.com/problems/climbing-stairs/
// Author : Zhonghuan Gao
// Date   : 2020-01-01

/**********************************************************************************
*
Say you have an array for which the ith element is the price of a given stock on day i.
If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
Note that you cannot sell a stock before you buy one.

Example 1:
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.

Example 2:
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*
***********************************************************************************/

package main

/**
解法一：暴力法
 */
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	maxProfit := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if prices[i] - prices[j] > maxProfit {
				maxProfit = prices[i] - prices[j]
			}
		}
	}
	return maxProfit
}

/**
解法二：动态规划
状态转移方程 dp[i] = max(prices[i] - minBuyPrice, maxProfit)
时间复杂度 O(n)
*/
func maxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	maxProfit := 0
	minBuyPrice := prices[0]
	for i := 1; i < n; i++ {
		if prices[i] < minBuyPrice {
			minBuyPrice = prices[i]
		} else if prices[i] - minBuyPrice > maxProfit {
			maxProfit = prices[i] - minBuyPrice
		}
	}
	return maxProfit
}

