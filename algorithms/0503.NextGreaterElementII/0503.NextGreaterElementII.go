// Source : https://leetcode.com/problems/next-greater-element-ii/
// Author : Zhonghuan Gao
// Date   : 2022-02-14

/**********************************************************************************
*

*
***********************************************************************************/

package main

/**
解法一：单调栈
思路: 将两个数据拼接成一个大数组，与739解法一致
*/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	newNums := make([]int, 2*n)
	for i := 0; i < n; i++ {
		newNums[i], newNums[i+n] = nums[i], nums[i]
	}
	answer := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		answer[i] = -1
	}

	stack := []int{}
	for i, num := range newNums {
		for len(stack) > 0 && num > newNums[stack[len(stack)-1]] {
			// top
			top := stack[len(stack)-1]
			answer[top] = num
			// pop
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answer[:n]
}

/**
解法二：单调栈（解法一的优化）
思路：https://programmercarl.com/0503.%E4%B8%8B%E4%B8%80%E4%B8%AA%E6%9B%B4%E5%A4%A7%E5%85%83%E7%B4%A0II.html
*/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	for i := 0; i < n; i++ {
		answer[i] = -1
	}

	stack := []int{}
	for i := 0; i < 2*n; i++ {
		// 模拟遍历两遍nums，注意都是用i%n来操作
		for len(stack) > 0 && nums[i%n] > nums[stack[len(stack)-1]] {
			// top
			top := stack[len(stack)-1]
			answer[top] = nums[i%n]
			// pop
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return answer
}
