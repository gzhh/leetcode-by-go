// Source : https://leetcode.com/problems/spiral-matrix/
// Author : Zhonghuan Gao
// Date   : 2022-02-22

/**********************************************************************************
*
Given an m x n matrix, return all elements of the matrix in spiral order.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*
***********************************************************************************/

package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, m*n)

	top, bottom := 0, m-1
	left, right := 0, n-1
	cur, end := 0, m*n-1
	for cur <= end {
		for i := left; cur <= end && i <= right; i++ {
			ans[cur] = matrix[top][i]
			cur++
		}
		top++

		for i := top; cur <= end && i <= bottom; i++ {
			ans[cur] = matrix[i][right]
			cur++
		}
		right--

		for i := right; cur <= end && i >= left; i-- {
			ans[cur] = matrix[bottom][i]
			cur++
		}
		bottom--

		for i := bottom; cur <= end && i >= top; i-- {
			ans[cur] = matrix[i][left]
			cur++
		}
		left++
	}
	return ans
}
