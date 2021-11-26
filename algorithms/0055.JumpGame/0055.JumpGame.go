// Source : https://leetcode.com/problems/jump-game/
// Author : Zhonghuan Gao
// Date   : 2021-11-26

/**********************************************************************************
*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
Return true if you can reach the last index, or false otherwise.

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
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
解法一：
*/
func canJump(nums []int) bool {
	dis, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if i <= dis {
			dis = max(dis, i+nums[i])
			if dis >= n-1 {
				return true
			}
		}
	}
	return false
}
