// Source : https://leetcode.com/problems/unique-paths-ii/
// Author : Zhonghuan Gao
// Date   : 2020-01-09

/**********************************************************************************
*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
Now consider if some obstacles are added to the grids. How many unique paths would there be?
An obstacle and space is marked as 1 and 0 respectively in the grid.

Example 1:
Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
Output: 2
Explanation: There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right

Example 2:
Input: obstacleGrid = [[0,1],[0,0]]
Output: 1

Constraints:
m == obstacleGrid.length
n == obstacleGrid[i].length
1 <= m, n <= 100
obstacleGrid[i][j] is 0 or 1.
*
***********************************************************************************/

package main

/**
解法一：动态规划
状态转移方程 dp[i][j] = sum(dp[i-1][j], dp[i][j-1])
时间复杂度：O(m*n)
空间复杂度：O(m*n)
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	dp := make([][]int, m)
	flagm := 0
	for i := range dp {
		dp[i] = make([]int, n)
		if flagm > 0 {
			continue
		} else if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			flagm = 1
		}
	}
	flagn := 0
	for i := 1; i < n; i++ {
		if flagn > 0 {
			continue
		} else if obstacleGrid[0][i] == 0 {
			dp[0][i] = 1
		} else {
			flagn = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}


/**
解法二：动态规划
解法一的优化，不使用额外存储空间
时间复杂度：O(m*n)
空间复杂度：O(m)
Ref: https://leetcode-cn.com/problems/unique-paths-ii/solution/bu-tong-lu-jing-ii-by-leetcode-solution-2/
 */
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j - 1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}
