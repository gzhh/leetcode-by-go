// Source : https://leetcode.com/problems/climbing-stairs/
// Author : Zhonghuan Gao
// Date   : 2020-01-01

/**********************************************************************************
*
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2

Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3

Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:
1 <= n <= 45
*
***********************************************************************************/

package main

/**
解法一：动态规划-斐波那契数列
状态转移方程 f[i] = f[i-1] + f[i-2]
时间复杂度 O(n)
*/
func climbStairs(n int) int {
    f := make([]int, n)
    f[0] = 1
    f[1] = 2
    i := 2
    for i++ < n {
        f[i] = f[i-1] + f[i-2]
        i++
    }
    return f[n-1]
}
