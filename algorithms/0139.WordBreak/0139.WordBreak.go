// Source : https://leetcode.com/problems/word-break/
// Author : Zhonghuan Gao
// Date   : 2020-02-17

/**********************************************************************************
*
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.
Note:
The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:
Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
Example 3:
Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
*
***********************************************************************************/

package main

/**
解法一：动态规划
思路：
	定义 dp[i] 表示字符串 s 前 i 个字符组成的字符串 s[0..i-1] 是否能被空格拆分成若干个字典中出现的单词
	枚举 s[0..i-1] 中的分割点 j ，看 s[0..j−1] 组成的字符串 s1(默认 j = 0 时 s1 为空串）和 s[j..i−1] 组成的字符串 s2​是否都合法
状态转移方程：
	dp[i] = dp[j] && check(s[j..i−1])
时间复杂度：O(n*n)
空间复杂度：O(n)
Ref: https://leetcode-cn.com/problems/word-break/solution/dan-ci-chai-fen-by-leetcode-solution/
*/
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, word := range wordDict {
		wordDictSet[word] = true
	}
	n := len(s)
	dp := make([]bool, n + 1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}