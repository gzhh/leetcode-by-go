// Source : https://leetcode.com/problems/decode-ways/
// Author : Zhonghuan Gao
// Date   : 2020-01-16

/**********************************************************************************
*
A message containing letters from A-Z can be encoded into numbers using the following mapping:
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode an encoded message, all the digits must be mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "111" can have each of its "1"s be mapped into 'A's to make "AAA", or it could be mapped to "11" and "1" ('K' and 'A' respectively) to make "KA". Note that "06" cannot be mapped into 'F' since "6" is different from "06".
Given a non-empty string num containing only digits, return the number of ways to decode it.
The answer is guaranteed to fit in a 32-bit integer.

Example 1:
Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).

Example 2:
Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

Example 3:
Input: s = "0"
Output: 0
Explanation: There is no character that is mapped to a number starting with 0. The only valid mappings with 0 are 'J' -> "10" and 'T' -> "20".
Since there is no character, there are no valid ways to decode this since all digits need to be mapped.

Example 4:
Input: s = "1"
Output: 1

Constraints:
1 <= s.length <= 100
s contains only digits and may contain leading zero(s).
*
***********************************************************************************/

package main

/**
解法一：动态规划
dp[i]为s[0..i]的译码方法总数
状态转移方程：
if s[i] == '0':
	dp[i] = dp[i-1]
else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6'):
	dp[i] = dp[i-1] + dp[i-2]
else:
	dp[i] = d[i-1]
时间复杂度：O(n)
空间复杂度：由分析可知，dp[i]仅与前两项有关，可用单变量代替dp数组，故空间复杂度可由O(n)降到O(1)
Ref: https://leetcode-cn.com/problems/decode-ways/solution/c-wo-ren-wei-hen-jian-dan-zhi-guan-de-jie-fa-by-pr/
*/
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur = pre + cur
		}
		pre = tmp
	}
	return cur
}