// Source : https://leetcode.com/problems/partition-equal-subset-sum
// Author : Zhonghuan Gao
// Date   : 2021-12-15

/**********************************************************************************
**
Given a non-empty array nums containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Example 1:
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.

Constraints:
1 <= nums.length <= 200
1 <= nums[i] <= 100
**
**********************************************************************************/

package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canPartition(nums []int) bool {
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([]int, 10001)
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	if dp[target] == target {
		return true
	}
	return false
}
