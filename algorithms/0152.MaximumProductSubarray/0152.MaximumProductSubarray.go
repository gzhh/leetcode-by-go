// Source : https://leetcode.com/problems/maximum-product-subarray/
// Author : Zhonghuan Gao
// Date   : 2020-02-13

/**********************************************************************************
*
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:
Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*
***********************************************************************************/

package main

/**
解法一：动态规划
状态转移方程：
	维护两个变量，当前最大和最小的连续数乘积
	maxDp[i] = max(maxDp[i-1] * nums[i], max(nums[i], minDp[i-1] * nums[i]))
	minDp[i] = min(minDp[i-1] * nums[i], min(nums[i], maxDp[i-1] * nums[i]))
时间复杂度：O(n)
空间复杂度：O(n)
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

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	maxDp := make([]int, n)
	minDp := make([]int, n)
	maxDp[0] = nums[0]
	minDp[0] = nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		maxDp[i] = max(maxDp[i-1] * nums[i], max(nums[i], minDp[i-1] * nums[i]))
		minDp[i] = min(minDp[i-1] * nums[i], min(nums[i], maxDp[i-1] * nums[i]))
		res = max(res, maxDp[i])
	}
	return res
}

/**
解法一的优化
时间复杂度：O(n)
空间复杂度：O(1)
 */
func maxProduct2(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx * nums[i], max(nums[i], mn * nums[i]))
		minF = min(mn * nums[i], min(nums[i], mx * nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}
