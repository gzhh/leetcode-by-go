<<<<<<< HEAD
// Source : https://leetcode-cn.com/problems/fibonacci-number/
// Author : Zhonghuan Gao
// Date   : 2021-12-13
=======
// Source : https://leetcode.com/problems/fibonacci-number/
// Author : Zhonghuan Gao
// Date   : 2021-12-09
>>>>>>> bb2f200f8c69ade097372692dbd645fd8264c513

/**********************************************************************************
*
The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).

Example 1:
Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

Example 2:
Input: n = 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

Example 3:
Input: n = 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.

Constraints:
<<<<<<< HEAD
0 <= n <= 30
=======
    0 <= n <= 30
>>>>>>> bb2f200f8c69ade097372692dbd645fd8264c513
*
***********************************************************************************/

package main

<<<<<<< HEAD
// recursal
=======
/**
解法一：递归
*/
>>>>>>> bb2f200f8c69ade097372692dbd645fd8264c513
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

<<<<<<< HEAD
// iterative: array
func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// iterative: variable exchange
func fib(n int) int {
=======
/**
解法二：动态规划
*/
func fib2(n int) int {
	if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

/**
解法三：解法二的优化
*/
func fib3(n int) int {
>>>>>>> bb2f200f8c69ade097372692dbd645fd8264c513
	if n <= 1 {
		return n
	}
	a, b := 0, 1
<<<<<<< HEAD
	for n > 1 {
		a, b = b, a+b
		n--
=======
	for i := 2; i <= n; i++ {
		a, b = b, a+b
>>>>>>> bb2f200f8c69ade097372692dbd645fd8264c513
	}
	return b
}
