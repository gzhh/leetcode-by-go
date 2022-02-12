// Source : https://leetcode.com/problems/palindromic-substrings/
// Author : Zhonghuan Gao
// Date   : 2020-01-23

/**********************************************************************************
*
Given a string, your task is to count how many palindromic substrings in this string.
The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

Example 1:
Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

Example 2:
Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

Note:
The input string length won't exceed 1000.
*
***********************************************************************************/

package main

/**
解法一：动态规划
dp[i][j]标记s[i:j]的是否是回文串
状态转移方程：
	s[i:j]由单个字符组成；则 dp[i][j] = true
	s[i:j]由2个字符组成，且两个字符相同 s[i] = s[j]；则 dp[i][j] = true
	s[i:j]由超过2个字符组成，且首尾字符相同 s[i] = s[j]，且除去头尾剩余子串是一个回文串 dp[i+1][j-1] = true；则 dp[i][j] = true
时间复杂度：O(n*n)
空间复杂度：O(n*n)
Ref: https://leetcode-cn.com/problems/palindromic-substrings/solution/shou-hua-tu-jie-dong-tai-gui-hua-si-lu-by-hyj8/
*/
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	cnt := 0
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
				cnt++
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
				cnt++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				cnt++
			}
		}
	}
	return cnt
}

/**
思路：https://programmercarl.com/0647.%E5%9B%9E%E6%96%87%E5%AD%90%E4%B8%B2.html
*/
func countSubstrings1(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	cnt := 0
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
				cnt++
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
				cnt++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				cnt++
			}
		}
	}
	return cnt
}

/**
解法二：中心拓展
Ref: https://leetcode-cn.com/problems/palindromic-substrings/solution/hui-wen-zi-chuan-by-leetcode-solution/
思路：https://programmercarl.com/0647.%E5%9B%9E%E6%96%87%E5%AD%90%E4%B8%B2.html
*/
func countSubstrings2(s string) int {
	n := len(s)
	var cnt int
	for i := 0; i < n; i++ {
		cnt += extend(s, i, i, n)   // 以i为中心
		cnt += extend(s, i, i+1, n) // 以i和i+1为中心
	}
	return cnt
}

func extend(s string, l, r, n int) int {
	var cnt int
	for l >= 0 && r < n && s[l] == s[r] {
		l--
		r++
		cnt++
	}
	return cnt
}

/**
解法三：马拉车
Ref: https://leetcode-cn.com/problems/palindromic-substrings/solution/hui-wen-zi-chuan-by-leetcode-solution/
*/
