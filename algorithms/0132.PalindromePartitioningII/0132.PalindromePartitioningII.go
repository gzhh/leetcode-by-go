// Source : https://leetcode.com/problems/palindrome-partitioning-ii/
// Author : Zhonghuan Gao
// Date   : 2020-01-30

/**********************************************************************************
*
Given a string s, partition s such that every substring of the partition is a palindrome.
Return the minimum cuts needed for a palindrome partitioning of s.

Example 1:
Input: s = "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

Example 2:
Input: s = "a"
Output: 0

Example 3:
Input: s = "ab"
Output: 1

Constraints:
1 <= s.length <= 2000
s consists of lower-case English letters only.
*
***********************************************************************************/

package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
解法一：动态规划
isPalindromic[i][j] 表示s[i,j]是不是回文子串(参考 5.最长回文子串、647.回文子串)
dp[i] 表示s[0,i]的回文子串，最少分割次数
状态转移方程
if isPalindromic[j+1][i]:
	dp[i] = min(dp[i], dp[j] + 1)
时间复杂度：O(n*n)
空间复杂度：O(n*n)
Ref: https://leetcode-cn.com/problems/palindrome-partitioning-ii/solution/132-fen-ge-hui-wen-chuan-iidong-tai-gui-6fyn1/
 */
func minCut(s string) int {
	n := len(s)

	isPalindromic := make([][]bool, n)
	for i := range isPalindromic {
		isPalindromic[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j - i <= 1 || isPalindromic[i+1][j-1]) {
				isPalindromic[i][j] = true
			}
		}
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = i
	}
	for i := 1; i < n; i++ {
		if isPalindromic[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if isPalindromic[j+1][i] {
				dp[i] = min(dp[i], dp[j] + 1)
			}
		}
	}
	return dp[n-1]
}