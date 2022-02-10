// Source : https://leetcode.com/problems/maximum-subarray/
// Author : Zhonghuan Gao
// Date   : 2020-01-01

/**********************************************************************************
*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [0]
Output: 0

Example 4:
Input: nums = [-1]
Output: -1

Example 5:
Input: nums = [-2147483647]
Output: -2147483647

Constraints:
1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
*
***********************************************************************************/

package main

/**
解法一：
固定右边界，左边界不断向1移动，求所有区间和的最大值；然后右边界也不断向1移动。
时间复杂度：O(n*n)
*/

/**
解法二：动态规划
状态转移方程：f[i] = max(f[i-1]+nums[i], nums[i])
时间复杂度：O(n)
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	dp[0] = nums[0]

	res := nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		res = max(res, dp[i])
	}

	return res
}
