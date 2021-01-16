// Source : https://leetcode.com/problems/triangle/
// Author : Zhonghuan Gao
// Date   : 2020-01-16

/**********************************************************************************
*
Given a triangle array, return the minimum path sum from top to bottom.
For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

Example 1:
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).

Example 2:
Input: triangle = [[-10]]
Output: -10

Constraints:
1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104

Follow up: Could you do this using only O(n) extra space, where n is the total number of rows in the triangle?
*
***********************************************************************************/

package main

import "math"

/**
解法一：动态规划
dp[i][j]为三角形从顶到点triangle[i][j]的最短路径
状态转移方程：dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
时间复杂度：O(n*n)
空间复杂度：O(n*n)
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			dp[i][i] = triangle[i][i]
		} else {
			dp[i][0] += dp[i-1][0] + triangle[i][0]
			dp[i][i] += dp[i-1][i-1] + triangle[i][i]
		}
	}
	minimum := min(dp[1][0], dp[1][1])
	if n == 2 {
		return minimum
	}
	for i := 2; i < n; i++ {
		tmpMin := dp[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			tmpMin = min(dp[i][j], tmpMin)
		}
		minimum = min(tmpMin, dp[i][i])
	}
	return minimum
}

/**
解法二：动态规划+空间优化
状态转移方程：dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
时间复杂度：O(n*n)
空间复杂度：O(n)
Ref: https://leetcode-cn.com/problems/triangle/solution/san-jiao-xing-zui-xiao-lu-jing-he-by-leetcode-solu/
 */
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	f := make([]int, n)
	f[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i] = f[i - 1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			f[j] = min(f[j - 1], f[j]) + triangle[i][j]
		}
		f[0] += triangle[i][0]
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, f[i])
	}
	return ans
}