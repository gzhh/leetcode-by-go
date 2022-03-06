// Source : https://leetcode.com/problems/binary-tree-paths/
// Author : Zhonghuan Gao
// Date   : 2022-03-06

/**********************************************************************************
*
Given the root of a binary tree, return all root-to-leaf paths in any order.
A leaf is a node with no children.

Example 1:
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]

Example 2:
Input: root = [1]
Output: ["1"]

Constraints:
The number of nodes in the tree is in the range [1, 100].
-100 <= Node.val <= 100
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
var res []string

func tranversal(root *TreeNode, path string) {
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, path)
		return
	}

	if root.Left != nil {
		tranversal(root.Left, path+"->")
	}
	if root.Right != nil {
		tranversal(root.Right, path+"->")
	}
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var path string
	res = []string{}
	tranversal(root, path)
	return res
}
