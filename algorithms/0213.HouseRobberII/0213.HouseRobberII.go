// Source : https://leetcode.com/problems/house-robber-ii/
// Author : Zhonghuan Gao
// Date   : 2020-02-11

/**********************************************************************************
*
you are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.
Given a list of non-negative integers nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.

Example 2:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 3:
Input: nums = [0]
Output: 0

Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 1000
*
***********************************************************************************/

// 题意：打家劫舍II
package main

/**
解法一：动态规划，参考198，是198的扩展

思路：
设dp[i]为前i个房子能偷的最大金额
状态转移方程：dp[i] = max(dp[i-1], dp[i-2]+nums[i])
环状排列意味着第一个房子和最后一个房子中只能选择一个偷窃，因此可以把此环状排列房间问题约化为两个单排排列房间子问题：
在不偷窃第一个房子的情况下（即 nums[1:]nums[1:]），最大金额是 p1
1.在不偷窃最后一个房子的情况下（即 nums[:n-1]nums[:n−1]），最大金额是 p2
2.综合偷窃最大金额： 为以上两种情况的较大值，即 max(p1,p2)max(p1,p2)

时间复杂度：O(n)
空间复杂度：O(n)
参考：https://leetcode-cn.com/problems/house-robber-ii/solution/213-da-jia-jie-she-iidong-tai-gui-hua-jie-gou-hua-/
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rec(nums []int) int {
	pprev := 0
	prev := 0
	for _, num := range nums {
		pprev, prev = prev, max(prev, pprev+num)
	}
	return prev
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(rec(nums[:n-1]), rec(nums[1:]))
}