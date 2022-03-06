// Source : https://leetcode.com/problems/balanced-binary-tree/
// Author : Zhonghuan Gao
// Date   : 2022-03-06

/**********************************************************************************
*
Given a binary tree, determine if it is height-balanced.
For this problem, a height-balanced binary tree is defined as:
a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: true

Example 2:
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

Example 3:
Input: root = []
Output: true

Constraints:
The number of nodes in the tree is in the range [0, 5000].
-104 <= Node.val <= 104
*
***********************************************************************************/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
解法一：递归
*/
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := getDepth(root.Left)
	if lDepth == -1 {
		return -1
	}
	rDepth := getDepth(root.Right)
	if rDepth == -1 {
		return -1
	}

	diff := 0
	if lDepth > rDepth {
		diff = lDepth - rDepth
	} else {
		diff = rDepth - lDepth
	}

	// 左右子书高度差，不满足则返回-1
	if diff > 1 {
		return -1
	}
	// 满足，返回具体高度
	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}

func isBalanced(root *TreeNode) bool {
	if getDepth(root) == -1 {
		return false
	}
	return true
}
