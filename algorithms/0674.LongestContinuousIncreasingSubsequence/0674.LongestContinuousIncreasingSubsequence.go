// Source : https://leetcode.com/problems/longest-continuous-increasing-subsequence/
// Author : Zhonghuan Gao
// Date   : 2022-01-25

/**********************************************************************************
*
Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence (i.e. subarray). The subsequence must be strictly increasing.
A continuous increasing subsequence is defined by two indices l and r (l < r) such that it is [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].

Example 1:
Input: nums = [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
Even though [1,3,5,7] is an increasing subsequence, it is not continuous as elements 5 and 7 are separated by element
4.

Example 2:
Input: nums = [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2] with length 1. Note that it must be strictly
increasing.

Constraints:
1 <= nums.length <= 104
-109 <= nums[i] <= 109
*
***********************************************************************************/

package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
方法一：暴力法
时间复杂度：O(n*n)
空间复杂度：O(n)
*/
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	lcis := 1

	for i := 0; i < n; i++ {
		cnt := 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[j+1] {
				cnt = 1
				lcis = max(lcis, cnt)
				continue
			}
			cnt++
		}
		lcis = max(lcis, cnt)
	}

	return lcis
}

/**
方法二：动态规划
时间复杂度：O(n)
空间复杂度：O(n)
*/
func findLengthOfLCIS2(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	lcis := 1
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			dp[i+1] = dp[i] + 1
		}
		lcis = max(lcis, dp[i+1])
	}

	return lcis
}

/**
方法三：贪心
时间复杂度：O(n)
空间复杂度：O(1)
*/
func findLengthOfLCIS3(nums []int) int {
	n := len(nums)

	lcis := 1
	cnt := 1
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			cnt++
		} else {
			cnt = 1
		}
		lcis = max(lcis, cnt)
	}

	return lcis
}
