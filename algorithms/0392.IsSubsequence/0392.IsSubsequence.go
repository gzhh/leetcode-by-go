// Source : https://leetcode.com/problems/is-subsequence/
// Author : Zhonghuan Gao
// Date   : 2020-01-02

/**********************************************************************************
*
Given a string s and a string t, check if s is subsequence of t.
A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ace" is a subsequence of "abcde" while "aec" is not).

Follow up:
If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you want to check one by one to see if T has its subsequence. In this scenario, how would you change your code?

Credits:
Special thanks to @pbrother for adding this problem and creating all test cases.

Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false

Constraints:
0 <= s.length <= 100
0 <= t.length <= 10^4
Both strings consists only of lowercase characters.
*
***********************************************************************************/

package main

/**
解法一：双指针
*/
func isSubsequence(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	if sLen == 0 {
		return true
	}
	if tLen == 0 {
		return false
	}
	j := 0
	for i := 0; i < tLen; i++ {
		if t[i] == s[j] {
			j++
		}
		if j == sLen {
			return true
		}
	}
	return false
}

/**
解法二：动态规划（编辑距离）
思路：https://programmercarl.com/0392.%E5%88%A4%E6%96%AD%E5%AD%90%E5%BA%8F%E5%88%97.html
*/
func isSubsequence2(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	if sLen == 0 {
		return true
	}
	if tLen == 0 {
		return false
	}

	dp := make([][]int, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([]int, tLen+1)
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[sLen][tLen] == sLen
}
