// Source : https://leetcode.com/problems/last-stone-weight-ii/
// Author : Zhonghuan Gao
// Date   : 2021-12-15

/**********************************************************************************
**
You are given an array of integers stones where stones[i] is the weight of the ith stone.
We are playing a game with the stones. On each turn, we choose any two stones and smash them together. Suppose the stones have weights x and y with x <= y. The result of this smash is:
- If x == y, both stones are destroyed, and
- If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.
Return the smallest possible weight of the left stone. If there are no stones left, return 0.

Example 1:
Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We can combine 2 and 4 to get 2, so the array converts to [2,7,1,8,1] then,
we can combine 7 and 8 to get 1, so the array converts to [2,1,1,1] then,
we can combine 2 and 1 to get 1, so the array converts to [1,1,1] then,
we can combine 1 and 1 to get 0, so the array converts to [1], then that's the optimal value.

Example 2:
Input: stones = [31,26,33,21,40]
Output: 5

Example 3:
Input: stones = [1,2]
Output: 1

Constraints:
1 <= stones.length <= 30
1 <= stones[i] <= 100
**
**********************************************************************************/

package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lastStoneWeightII(stones []int) int {
	sum := 0
	n := len(stones)
	for i := 0; i < n; i++ {
		sum += stones[i]
	}
	target := sum / 2

	dp := make([]int, 15001)
	for i := 0; i < n; i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - 2*dp[target]
}
