// Source : https://leetcode.com/problems/maximal-square/
// Author : Zhonghuan Gao
// Date   : 2020-02-21

/**********************************************************************************
*
Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

Example 1:
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4

Example 2:
Input: matrix = [["0","1"],["1","0"]]
Output: 1

Example 3:
Input: matrix = [["0"]]
Output: 0

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] is '0' or '1'.

*
***********************************************************************************/

package main

/**
解法一：动态规划
思路：
	定义 dp[i][j] 表示点i,j所处正方形的最大面积
状态转移方程：
	if matrix[i][j] == 1 && dp[i-1][j-1] >= 1 && dp[i-1][j] >= 1 && dp[i][j-1] >= 1
		// 取三者最小值
		dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), min(dp[i-1][j-1], dp[i][j-1])) + 1
	else if matrix[i][j] == 1
		dp[i][j] = 1
	else
		dp[i][j] = 0
时间复杂度：O(m*n)
空间复杂度：O(m*n)
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if int(matrix[i][0]) == 49 {
			dp[i][0] = 1
		}
	}
	for j := 0; j < n; j++ {
		if int(matrix[0][j]) == 49 {
			dp[0][j] = 1
		}
	}

	maxEdge := 0
	if int(matrix[0][0]) == 49 {
		maxEdge = 1
	}
	if m > 1 && int(matrix[1][0]) == 49 {
		maxEdge = 1
	}
	if n > 1 && int(matrix[0][1]) == 49 {
		maxEdge = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if int(matrix[i][j]) == 49  {
				if dp[i-1][j-1] >= 1 && dp[i-1][j] >= 1 && dp[i][j-1] >= 1 {
					dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), min(dp[i-1][j-1], dp[i][j-1])) + 1
				} else {
					dp[i][j] = 1
				}
				maxEdge = max(maxEdge, dp[i][j])
			}
		}
	}

	return maxEdge * maxEdge
}