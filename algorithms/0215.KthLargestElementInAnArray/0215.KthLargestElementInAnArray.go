// Source : https://leetcode.com/problems/kth-largest-element-in-an-array/
// Author : Zhonghuan Gao
// Date   : 2020-07-30

/**********************************************************************************
*
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:
Input: [3,2,1,5,6,4] and k = 2
Output: 5

Example 2:
Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4

Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*
***********************************************************************************/

package main

/**
解法一：先排序，再找第k大的数
 */
func minHeapify(nums []int, start int, end int) {
	dad := start
	son := dad * 2 + 1
	if son > end {
		return
	}

	if son + 1 <= end && nums[son] < nums[son + 1] {
		son++
	}

	if nums[son] > nums[dad] {
		nums[son], nums[dad] = nums[dad], nums[son]
		minHeapify(nums, son, end)
	}
}

func heapSort(nums []int, length int) {
	for i := length / 2 - 1; i >= 0; i-- {
		minHeapify(nums, i, length - 1)
	}
	for i := length - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		minHeapify(nums, 0, i - 1)
	}
}

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	heapSort(nums, length)
	return nums[length - k]
}

/**
解法二：局部排序
 */

/**
解法三：只找topk，不排序
 */



var sum = 0

func sumNumbers(root *TreeNode) int {
	var leafSum, level = 0, 0
	dfs(root, leafSum, level)
	return sum
}

func dfs(root *TreeNode, leafSum, level int) {
	if root == nil {
		sum += leafSum
		return
	}
	recurse(root.Left, leafSum, level)
	level++
	leafSum += 10 * level + root.Val
	recurse(root.Right, leafSum, level)

}