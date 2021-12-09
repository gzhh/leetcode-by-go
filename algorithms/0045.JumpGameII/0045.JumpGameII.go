// Source : https://leetcode.com/problems/jump-game-ii/
// Author : Zhonghuan Gao
// Date   : 2021-11-27

/**********************************************************************************
*
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Your goal is to reach the last index in the minimum number of jumps.
You can assume that you can always reach the last index.

Example 1:
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [2,3,0,1,4]
Output: 2
*
***********************************************************************************/

package main

/**
解法一：贪心
*/
func jump(nums []int) int {
	n := len(nums)
	step := 0
	pos := n - 1
	for pos > 0 {
		for i := 0; i < pos; i++ {
			if i+nums[i] >= pos {
				step++
				pos = i
				break
			}
		}
	}
	return step
}

/**
解法二：DP
*/
