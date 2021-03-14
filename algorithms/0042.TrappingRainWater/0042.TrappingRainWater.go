// Source : https://leetcode.com/problems/trapping-rain-water/
// Author : Zhonghuan Gao
// Date   : 2020-03-14

/**********************************************************************************
*
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it can trap after raining.

Example 1:
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
In this case, 6 units of rain water (blue section) are being trapped.

Example 2:
Input: height = [4,2,0,3,2,5]
Output: 9

Constraints:
n == height.length
0 <= n <= 3 * 104
0 <= height[i] <= 105
*
***********************************************************************************/

package main

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

/**
解法一：暴力法
思路：对于数组中的每个元素，我们找出下雨后水能达到的最高位置，等于两边最大高度的较小值减去当前高度的值。
Ref: https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/
*/
func trap(height []int) int {
	n := len(height)
	maxSum := 0
	for i := 1; i < n - 1; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		for j := i; j < n; j++ {
			maxRight = max(maxRight, height[j])
		}
		maxSum += min(maxLeft, maxRight) - height[i]
	}
	return maxSum
}

/**
解法二：动态规划（解法一的优化）
思路：使用数组预存解法一中的 maxLeft 和 maxRight
 */
func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	maxLeft := make([]int, n)
	maxLeft[0] = height[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	maxRight := make([]int, n)
	maxRight[n-1] = height[n-1]
	for i := n - 2; i >- 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	maxSum := 0
	for i := 1; i < n - 1; i++ {
		maxSum += min(maxLeft[i], maxRight[i]) - height[i]
	}
	return maxSum
}

/**
解法三：双指针解法
Ref: https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/
 */
func trap3(height []int) int {
	n := len(height)
	left, right := 0, n - 1
	maxLeft, maxRight, maxSum := 0, 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				maxSum += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				maxSum += maxRight - height[right]
			}
			right--
		}
	}

	return maxSum
}
