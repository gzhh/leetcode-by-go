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
func maxSubArray(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	maxSum := f[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if f[i-1] > 0 {
			f[i] = f[i-1] + num
		} else {
			f[i] = num
		}
		if maxSum < f[i] {
			maxSum = f[i]
		}
	}
	return maxSum
}
