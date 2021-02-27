// Source : https://leetcode.com/problems/ugly-number/
// Author : Zhonghuan Gao
// Date   : 2020-02-27

/**********************************************************************************
Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example:
Input: n = 10
Output: 12
Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.

Note:
1 is typically treated as an ugly number.
n does not exceed 1690.
*
***********************************************************************************/

package main

/**
解法一：动态规划（三指针）
思路：先排序后存
Ref:
	https://leetcode-cn.com/problems/ugly-number-ii/solution/chou-shu-ii-by-leetcode/
	https://leetcode-cn.com/problems/ugly-number-ii/solution/bao-li-you-xian-dui-lie-xiao-ding-dui-dong-tai-gui/
	https://leetcode-cn.com/problems/ugly-number-ii/solution/san-zhi-zhen-fang-fa-de-li-jie-fang-shi-by-zzxn/
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = min(min(2 * dp[p2], 3 * dp[p3]), 5 * dp[p5])
		if dp[i] == 2 * dp[p2] {
			p2++
		}
		if dp[i] == 3 * dp[p3] {
			p3++
		}
		if dp[i] == 5 * dp[p5] {
			p5++
		}
	}
	return dp[n-1]
}

/**
解法二 C++：最小堆（优先队列）
思路：先存后排序

class Solution {
public:
    int nthUglyNumber(int n) {
        priority_queue <long long, vector<long long>, greater<long long> > q;
        long long ans = 1;
        for (int i = 1; i < n; i++) {
            q.push(ans * 2);
            q.push(ans * 3);
            q.push(ans * 5);
            ans = q.top();
            q.pop();
            while (!q.empty() && ans == q.top())
                q.pop();
        }
        return ans;
    }
};
 */
