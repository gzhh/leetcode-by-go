// Source : https://leetcode.com/problems/palindrome-partitioning/
// Author : Zhonghuan Gao
// Date   : 2020-01-30

/**********************************************************************************
*
Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.
A palindrome string is a string that reads the same backward as forward.

Example 1:
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]
Example 2:
Input: s = "a"
Output: [["a"]]

Constraints:
1 <= s.length <= 16
s contains only lowercase English letters.
*
***********************************************************************************/

package main

/**
解法一：dfs回溯法
Ref: https://leetcode-cn.com/problems/palindrome-partitioning/solution/131-fen-ge-hui-wen-chuan-hui-su-sou-suo-suan-fa-xi/
 */
func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func backtrack(res [][]string, cur[]string, s string, start int) [][]string {
	strLen := len(s)

	if start == strLen {
		curCopy := make([]string, len(cur))
		copy(curCopy, cur)
		res = append(res, curCopy)
		return res
	}

	for i := start; i < strLen; i++ {
		if isPalindrome(s, start, i) {
			cur = append(cur, s[start:i+1])
			res = backtrack(res, cur, s, i+1)
			cur = cur[:len(cur)-1]
		}
	}

	return res
}

func partition(s string) [][]string {
	res := [][]string{}
	cur := []string{}
	res = backtrack(res, cur, s, 0)
	return res
}



/**
解法二：动态规划+DFS
dp[start][i] 表示字符串从位置 start 到位置 i(闭区间)是否为回文子串
*/
func dfs(res [][]string, dp [][]bool, cur[]string, s string, start int) [][]string {
	strLen := len(s)

	if start == strLen {
		curCopy := make([]string, len(cur))
		copy(curCopy, cur)
		res = append(res, curCopy)
		return res
	}
	for i := start; i < strLen; i++ {
		if dp[start][i] {
			cur = append(cur, s[start:i+1])
			res = dfs(res, dp, cur, s, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	return res
}

func partition2(s string) [][]string {
	res := [][]string{}

	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i - j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
			}
		}
	}

	res = dfs(res, dp, []string{}, s, 0)
	return res
}