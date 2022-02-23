// Source : https://leetcode.com/problems/spiral-matrix-ii/
// Author : Zhonghuan Gao
// Date   : 2022-02-22

/**********************************************************************************
*
Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.

Example 1:
Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]

Example 2:
Input: n = 1
Output: [[1]]

Constraints:
1 <= n <= 20
*
***********************************************************************************/

package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	cur, end := 1, n*n
	for cur <= end {
		for i := left; i <= right; i++ {
			matrix[top][i] = cur
			cur++
		}
		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = cur
			cur++
		}
		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = cur
			cur++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			matrix[i][left] = cur
			cur++
		}
		left++
	}
	return matrix
}
