// Source : https://leetcode.com/problems/counting-bits/
// Author : Zhonghuan Gao
// Date   : 2020-03-21

/**********************************************************************************
*
Given an integer num, return an array of the number of 1's in the binary representation of every number in the range [0, num].

Example 1:
Input: num = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10

Example 2:
Input: num = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101


Constraints:
0 <= num <= 105
*
***********************************************************************************/

package main

/**
解法一：动态规划
思路：最低设置位，把最低位的1置为0
Ref: https://leetcode-cn.com/problems/counting-bits/solution/bi-te-wei-ji-shu-by-leetcode-solution-0t1i/
*/
func countBits(num int) []int {
	bits := make([]int, num + 1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}