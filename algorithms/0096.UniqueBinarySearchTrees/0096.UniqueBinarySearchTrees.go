// Source : https://leetcode.com/problems/unique-binary-search-trees/
// Author : Zhonghuan Gao
// Date   : 2020-01-16

/**********************************************************************************
*
Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

Example:
Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3


Constraints:
1 <= n <= 19
*
***********************************************************************************/

package main

/**
解法一：动态规划
dp[i]为s[0..i]的译码方法总数
状态转移方程：
dp[i] = dp[j-1] * dp[i-j] (dp[0]=1, dp[1]=1; i:2->n, j:1->i)
时间复杂度：O(n*n)
空间复杂度：O(n)
Ref: https://leetcode-cn.com/problems/unique-binary-search-trees/solution/bu-tong-de-er-cha-sou-suo-shu-by-leetcode-solution/
PS: 待复盘
*/
func numTrees(n int) int {
	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

/**
解法二：数学方法
卡塔兰数
 */