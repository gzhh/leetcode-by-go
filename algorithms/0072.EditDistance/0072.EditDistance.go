// Source : https://leetcode.com/problems/edit-distance/
// Author : Zhonghuan Gao
// Date   : 2022-02-12

/**********************************************************************************
*
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
You have the following three operations permitted on a word:
	- Insert a character
	- Delete a character
	- Replace a character

Example 1:
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')

Constraints:
0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.
*
***********************************************************************************/

func min(a, b int)int {
    if a < b {
        return a
    }
    return b
}

/**
解法一：动态规划
思路：https://programmercarl.com/0072.%E7%BC%96%E8%BE%91%E8%B7%9D%E7%A6%BB.html
*/
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
        dp[i][0] = i
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(min(dp[i][j-1]+1, dp[i-1][j]+1), dp[i-1][j-1]+1)
            }
        }
    }

    return dp[m][n]
}
