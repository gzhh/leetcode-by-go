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
	// dp[i][j] 表示字符串 s[i:j] 是否是回文串
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	ans := ""
	// 维护k，k(0->n)
	for k := 0; k < n; k++ {
		for i := 0; i + k < n; i++ {
			j := i + k
			if k == 0 {
				dp[i][j] = 1
			} else if k == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				// 只有 s[i+1:j-1] 是回文串，并且 s 的第 i 和 j 个字母相同时，s[i:j] 才会是回文串
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && k + 1 > len(ans) {
				ans = s[i:i+k+1]
			}
		}
	}

	// 只会是所有最大长度回文串中最先找到的那个回文串
	return ans
}


/**
解法二：中心扩展算法
时间复杂度：O(n*n)
*/
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