// Source : https://leetcode.com/problems/sort-colors/
// Author : Zhonghuan Gao
// Date   : 2020-04-13

/**********************************************************************************
*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]

Example 3:
Input: nums = [0]
Output: [0]

Example 4:
Input: nums = [1]
Output: [1]

Constraints:
n == nums.length
1 <= n <= 300
nums[i] is 0, 1, or 2.

Follow up:
Could you solve this problem without using the library's sort function?
Could you come up with a one-pass algorithm using only O(1) constant space?
Output: false
*
***********************************************************************************/

package main

/**
解法一：单指针
 */
func sortColors(nums []int)  {
	count0 := swapColors(nums, 0)
	swapColors(nums[count0:], 1)
}

func swapColors(nums []int, target int) (countTarget int) {
	for i, num := range nums {
		if num == target {
			nums[i], nums[countTarget] = nums[countTarget], nums[i]
			countTarget++
		}
	}
	return
}

/**
解法二：双指针
 */
func sortColors2(nums []int)  {
	// 使用p0来交换0，p1来交换1
	p0, p1 := 0, 0
	for i, num := range nums {
		if num == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if num == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}