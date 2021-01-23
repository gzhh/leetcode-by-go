// Source : https://leetcode.com/problems/wiggle-subsequence/
// Author : Zhonghuan Gao
// Date   : 2020-01-23

/**********************************************************************************
*
A sequence of numbers is called a wiggle sequence if the differences between successive numbers strictly alternate between positive and negative. The first difference (if one exists) may be either positive or negative. A sequence with fewer than two elements is trivially a wiggle sequence.
For example, [1,7,4,9,2,5] is a wiggle sequence because the differences (6,-3,5,-7,3) are alternately positive and negative. In contrast, [1,4,7,2,5] and [1,7,4,5,5] are not wiggle sequences, the first because its first two differences are positive and the second because its last difference is zero.
Given a sequence of integers, return the length of the longest subsequence that is a wiggle sequence. A subsequence is obtained by deleting some number of elements (eventually, also zero) from the original sequence, leaving the remaining elements in their original order.

Example 1:
Input: [1,7,4,9,2,5]
Output: 6
Explanation: The entire sequence is a wiggle sequence.

Example 2:
Input: [1,17,5,10,13,15,10,5,16,8]
Output: 7
Explanation: There are several subsequences that achieve this length. One is [1,17,10,13,10,16,8].

Example 3:
Input: [1,2,3,4,5,6,7,8,9]
Output: 2
Follow up:
Can you do it in O(n) time?
*
***********************************************************************************/

// 题意：给你一个整数数组，求它的摆动子序列的最大长度（摆动序列，相邻元素的差值正负交替）
package main

/**
解法一：贪心法
思路：遍历数组，计算当前元素与前一个元素的差，使用 preDiff 和 curDiff 分别表示上一个元素和当前元素的相邻元素差
当出现 curDiff 与 preDiff 的不同为正数或负数或零时，则最大摇摆子序列长度加一，并用 curDiff 替换 preDiff。
PS：注意差为零的情况
时间复杂度：O(n)
空间复杂度：O(1)
*/
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	res, preDiff, curDiff := 1, 0, 0
	for i := 1; i < n; i++ {
		curDiff = nums[i] - nums[i-1]
		if (curDiff > 0 && preDiff <= 0) || (curDiff < 0 && preDiff >= 0) {
			res++
			preDiff = curDiff
		}
	}
	return res
}

/**
解法二：动态规划
Ref: https://leetcode-cn.com/problems/wiggle-subsequence/solution/bai-dong-xu-lie-by-leetcode-solution-yh2m/
 */