// Source : https://leetcode.com/problems/largest-rectangle-in-histogram/
// Author : Zhonghuan Gao
// Date   : 2022-02-15

/**********************************************************************************
*
Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

Example 1:
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.

Example 2:
Input: heights = [2,4]
Output: 4

Constraints:
1 <= heights.length <= 105
0 <= heights[i] <= 104
*
***********************************************************************************/

package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
解法一：双指针（超时）
*/
func largestRectangleArea(heights []int) int {
	n := len(heights)
	sum := 0
	for i := 0; i < n; i++ {
		left, right := i, i
		for ; left >= 0; left-- {
			if heights[left] < heights[i] {
				break
			}
		}
		for ; right < n; right++ {
			if heights[right] < heights[i] {
				break
			}
		}
		w := right - left - 1
		h := heights[i]
		sum = max(sum, w*h)
	}
	return sum
}

/**
解法二：动态规划
*/
func largestRectangleArea2(heights []int) int {
	n := len(heights)

	minLeftIndex := make([]int, n)
	minLeftIndex[0] = -1
	for i := 1; i < n; i++ {
		cur := i - 1
		for cur >= 0 && heights[cur] >= heights[i] {
			cur = minLeftIndex[cur]
		}
		minLeftIndex[i] = cur
	}

	minRightIndex := make([]int, n)
	minRightIndex[n-1] = n
	for i := n - 2; i >= 0; i-- {
		cur := i + 1
		for cur < n && heights[cur] >= heights[i] {
			cur = minRightIndex[cur]
		}
		minRightIndex[i] = cur
	}

	sum := 0
	for i := 0; i < n; i++ {
		w := minRightIndex[i] - minLeftIndex[i] - 1
		h := heights[i]
		sum = max(sum, w*h)
	}
	return sum
}

/**
解法三：单调栈
参考：https://programmercarl.com/0084.%E6%9F%B1%E7%8A%B6%E5%9B%BE%E4%B8%AD%E6%9C%80%E5%A4%A7%E7%9A%84%E7%9F%A9%E5%BD%A2.html
*/
func largestRectangleArea3(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	n := len(heights)
	stack := []int{0}
	sum := 0
	for i := 1; i < n; i++ {
		if heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if heights[i] == heights[stack[len(stack)-1]] {
			// pop
			stack = stack[:len(stack)-1]
			// push
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				// top
				top := stack[len(stack)-1]
				// pop
				stack = stack[:len(stack)-1]

				if len(stack) > 0 {
					left := stack[len(stack)-1]
					right := i
					w := right - left - 1
					h := heights[top]
					sum = max(sum, w*h)
				}
			}
			// push
			stack = append(stack, i)
		}
	}
	return sum
}
