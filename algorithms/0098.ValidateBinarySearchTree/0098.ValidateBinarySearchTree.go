// Source : https://leetcode.com/problems/validate-binary-search-tree/
// Author : Zhonghuan Gao
// Date   : 2020-10-28

/**********************************************************************************
*
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:
The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Example 1:
    2
   / \
  1   3
Input: [2,1,3]
Output: true

Example 2:
    5
   / \
  1   4
     / \
    3   6
Input: [5,1,4,null,null,3,6]
Output: false

Explanation: The root node's value is 5 but its right child's value is 4.
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

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

/**
验证二叉搜索树
解法一：用二叉搜索树本身的性质来做，即左<根<右，根大于所有的左子树且小于所有的右子树
初始化时带入int类型最大值和最小值，在递归过程中换成它们自己的节点值
 */
func isValidBST(root *TreeNode) bool {
	return validate(root, INT_MIN, INT_MAX)
}

func validate(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return validate(root.Left, min, root.Val) && validate(root.Right, root.Val, max)
}