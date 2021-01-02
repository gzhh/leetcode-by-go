// Source : https://leetcode.com/problems/longest-increasing-subsequence/
// Author : Zhonghuan Gao
// Date   : 2020-01-02

/**********************************************************************************
*
Given an integer array nums, return the length of the longest strictly increasing subsequence.
A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1

Constraints:
1 <= nums.length <= 2500
-104 <= nums[i] <= 104

Follow up:
Could you come up with the O(n2) solution?
Could you improve it to O(n log(n)) time complexity?
*
***********************************************************************************/

package main

/**
解法一：动态规划
状态转移方程：dp[i]=max(dp[j])+1,其中0≤j<i且num[j]<num[i]
Ref: https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-by-leetcode-soluti/
时间复杂度：O(n*n)
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	lis := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
			lis = max(dp[i], lis)
		}
	}
	return lis
}


/**
解法二：贪心 + 二分查找
时间复杂度：O(nlogn)
 */