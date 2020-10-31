// Source : https://leetcode.com/problems/number-of-islands/
// Author : Zhonghuan Gao
// Date   : 2020-10-31

/**********************************************************************************
*
Given an m x n 2d grid map of '1's (land) and '0's (water), return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3


Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*
***********************************************************************************/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main


/**
解法一：递归
遍历二维数组，如果一个位置为 1，则以其为起始节点开始进行深度优先搜索，在深度优先搜索的过程中，
每个搜索到的 1 都会被重新标记为 0。最终岛屿的数量就是我们进行深度优先搜索的次数。
 */
var m, n int

func dfs(grid [][]byte, r, c int) {
	grid[r][c] = '0'
	if r - 1 >= 0 && grid[r-1][c] == '1' {
		dfs(grid, r - 1, c)
	}
	if r + 1 < m && grid[r+1][c] == '1' {
		dfs(grid, r + 1, c)
	}
	if c - 1 >= 0 && grid[r][c-1] == '1' {
		dfs(grid, r, c - 1)
	}
	if c + 1 < n && grid[r][c+1] == '1' {
		dfs(grid, r, c + 1)
	}
}

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	m, n = len(grid), len(grid[0])

	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(grid, i, j)
			}
		}
	}
	return cnt
}