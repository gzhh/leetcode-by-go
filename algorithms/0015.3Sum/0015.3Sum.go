// Source : https://leetcode.com/problems/3sum/
// Author : Zhonghuan Gao
// Date   : 2022-02-26

/**********************************************************************************
*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Example 2:
Input: nums = []
Output: []

Example 3:
Input: nums = [0]
Output: []

Constraints:
0 <= nums.length <= 3000
-105 <= nums[i] <= 105
*
***********************************************************************************/

package leetcode

import "sort"

/**
解法一：双指针
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return ans
		}
		// unique from i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// unique from left
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// unique from right
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return ans
}

/**
解法二：哈希表
和1.Two Sum类似，可以用哈希表实现，但是比较复杂难实现
*/
