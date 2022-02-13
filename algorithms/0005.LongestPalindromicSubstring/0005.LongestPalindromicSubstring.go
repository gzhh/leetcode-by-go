// Source : https://leetcode.com/problems/longest-increasing-subsequence/
// Author : Zhonghuan Gao
// Date   : 2020-01-02

/**********************************************************************************
*
Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Example 3:
Input: s = "a"
Output: "a"

Example 4:
Input: s = "ac"
Output: "a"

Constraints:
1 <= s.length <= 1000
s consist of only digits and English letters (lower-case and/or upper-case).
*
***********************************************************************************/

package main

/**
解法一：动态规划
状态转移方程：P(i,j)=P(i+1,j−1)∧(Si==Sj)
时间复杂度：O(n*n)
Ref: https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
*/
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// dp[i][j] 表示字符串 s[i:j] 是否是回文串
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		// 初始化
		dp[i][i] = true
	}

	begin := 0
	maxLen := 1
	// 枚举子串长度
	for L := 2; L <= n; L++ {
		// 枚举左边界
		for i := 0; i < n; i++ {
			// 获取右边界
			j := i + L - 1
			if j >= n {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				begin = i
				maxLen = j - i + 1
			}
		}
	}

	return s[begin : begin+maxLen]
}

/**
解法二：中心扩展算法
时间复杂度：O(n*n)
*/
func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expand(s, i, i)
		left2, right2 := expand(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start : end+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left, right = left-1, right+1
	}
	return left + 1, right - 1
}

/**
// cpp solution
string longestPalindrome(string s) {
  if (s.size() < 2) return s;
  int n = s.size(), maxLen = 0, start = 0;
  for (int i = 0; i < n;) {
	  if (n - i <= maxLen / 2) break;
	  int left = i, right = i;
	  while (right < n - 1 && s[right + 1] == s[right]) ++right;
	  i = right + 1;
	  while (right < n - 1 && left > 0 && s[right + 1] == s[left - 1]) {
		  ++right; --left;
	  }
	  if (maxLen < right - left + 1) {
		  maxLen = right - left + 1;
		  start = left;
	  }
  }
  return s.substr(start, maxLen);
}
*/

/**
解法三：马拉车算法（复杂）
时间复杂度：O(n)
*/
