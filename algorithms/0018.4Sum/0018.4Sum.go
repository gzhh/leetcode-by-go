// Source : https://leetcode.com/problems/4sum/
// Author : Zhonghuan Gao
// Date   : 2022-02-26

/**********************************************************************************
*
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

Example 1:
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Example 2:
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]

Constraints:
1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*
***********************************************************************************/

package leetcode

import "sort"

/**
解法一：双指针
类似15.3Sum，多了层循环
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for j := 0; j < n; j++ {
		// unique from j
		if j > 0 && nums[j] == nums[j-1] {
			continue
		}
		for i := j + 1; i < n; i++ {
			// unique from i
			if i > j+1 && nums[i] == nums[i-1] {
				continue
			}
			left, right := i+1, n-1
			for left < right {
				if nums[j]+nums[i]+nums[left]+nums[right] < target {
					left++
				} else if nums[j]+nums[i]+nums[left]+nums[right] > target {
					right--
				} else {
					ans = append(ans, []int{nums[j], nums[i], nums[left], nums[right]})
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
	}
	return ans
}
