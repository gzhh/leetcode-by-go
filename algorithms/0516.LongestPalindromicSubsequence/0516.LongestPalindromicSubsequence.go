// Source : https://leetcode.com/problems/longest-palindromic-subsequence/
// Author : Zhonghuan Gao
// Date   : 2020-01-23

/**********************************************************************************
*
Given a string s, find the longest palindromic subsequence's length in s. You may assume that the maximum length of s is 1000.

Example 1:
Input:
"bbbab"
Output:
4
One possible longest palindromic subsequence is "bbbb".

Example 2:
Input:
"cbbd"
Output:
2
One possible longest palindromic subsequence is "bb".

Constraints:
1 <= s.length <= 1000
s consists only of lowercase English letters.
*
***********************************************************************************/

package main

/**
解法一：动态规划
dp[i][j]为s[i]到s[j]的字符串中最长子序列的长度
状态转移方程：
if s[i] == s[j]:
	dp[i] = dp[i-1][j-1] + 2
else:
	dp[i] = max(dp[i][j-1], dp[i+1][j])
遍历顺序，i从最后一个字符开始往前遍历，j从i+1开始往后遍历，初始化dp[i][i] = 1
时间复杂度：O(n*n)
空间复杂度：O(n*n)
Ref: https://leetcode-cn.com/problems/longest-palindromic-subsequence/solution/dong-tai-gui-hua-si-yao-su-by-a380922457-3/
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}