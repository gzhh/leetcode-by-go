// Source : https://leetcode-cn.com/problems/daily-temperatures/
// Author : Zhonghuan Gao
// Date   : 2022-02-14

/**********************************************************************************
*
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]

Constraints:
1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*
***********************************************************************************/

package main

/**
解法一：暴力法
*/
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				answer[i] = j - i
				break
			}
		}
	}
	return answer
}

/**
解法二：单调栈
思路：https://programmercarl.com/0739.%E6%AF%8F%E6%97%A5%E6%B8%A9%E5%BA%A6.html
*/
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	stack := []int{}
	for i, num := range temperatures {
		for len(stack) > 0 && num > temperatures[stack[len(stack)-1]] {
			// top
			top := stack[len(stack)-1]
			answer[top] = i - top
			// pop
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answer
}
