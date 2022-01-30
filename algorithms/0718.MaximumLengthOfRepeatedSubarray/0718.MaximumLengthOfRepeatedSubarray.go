// Source : https://leetcode.com/problems/maximum-length-of-repeated-subarray/
// Author : Zhonghuan Gao
// Date   : 2022-01-30

/**********************************************************************************
*
Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.

Example 1:
Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
Output: 3
Explanation: The repeated subarray with maximum length is [3,2,1].

Example 2:
Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
Output: 5

Constraints:
1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 100
*
***********************************************************************************/

package main

/**
解法一：动态规划
思路：
1.dp[i][j] ：以下标i - 1为结尾的num1，和以下标j - 1为结尾的nums2，最长重复子数组长度为dp[i][j]。
2.根据dp[i][j]的定义，dp[i][j]的状态只能由dp[i - 1][j - 1]推导出来。即当nums1[i - 1] 和nums2[j - 1]相等的时候，dp[i][j] = dp[i - 1][j - 1] + 1
3.dp[i][0]和dp[0][j]初始化为0
Ref: https://programmercarl.com/0718.%E6%9C%80%E9%95%BF%E9%87%8D%E5%A4%8D%E5%AD%90%E6%95%B0%E7%BB%84.html
*/
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
