// Source : https://leetcode.com/problems/binary-tree-paths/
// Author : Zhonghuan Gao
// Date   : 2020-11-08

/**********************************************************************************
*
Given a binary tree, return all root-to-leaf paths.
Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3
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
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
解法一：递归
 */
var res []string

func binaryTreePaths(root *TreeNode) []string {
	res = make([]string, 0)
	if root == nil {
		return res
	}
	dfs(root, "")
	return res
}

func dfs(root *TreeNode, str string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		str = str + strconv.Itoa(root.Val)
		res = append(res, str)
		return
	}
	str = str + strconv.Itoa(root.Val) + "->"
	dfs(root.Left, str)
	dfs(root.Right, str)
}