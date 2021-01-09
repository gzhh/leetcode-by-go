// Source : https://leetcode.com/problems/unique-paths/
// Author : Zhonghuan Gao
// Date   : 2020-01-09

/**********************************************************************************
*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?


Example 1:
Input: m = 3, n = 7
Output: 28

Example 2:
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down

Example 3:
Input: m = 7, n = 3
Output: 28

Example 4:
Input: m = 3, n = 3
Output: 6

Constraints:
1 <= m, n <= 100
It's guaranteed that the answer will be less than or equal to 2 * 109.
*
***********************************************************************************/

package main

/**
解法一：动态规划
状态转移方程 dp[i][j] = sum(dp[i-1][j], dp[i][j-1])
时间复杂度：O(m*n)
空间复杂度：O(m*n)
*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][1] = 1
	}
	for i := 1; i <= n; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}
