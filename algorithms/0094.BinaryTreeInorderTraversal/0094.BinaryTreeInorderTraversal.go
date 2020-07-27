// Source : https://leetcode.com/problems/binary-tree-inorder-traversal/
// Author : Zhonghuan Gao
// Date   : 2020-07-24

/**********************************************************************************
*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]

Follow up: Recursive solution is trivial, could you do it iteratively?
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

/**
解法一：递归
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	rec(root, &res)
	return res
}

func rec(root *TreeNode, output *[]int) {
	if root != nil {
		rec(root.Left, output)
		*output = append(*output, root.Val)
		rec(root.Right, output)
	}
}
/**
解法二：
用栈模拟递归过程
 */
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*TreeNode{}, []int{}
	node := root
	for node != nil || len(stack) != 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			tmpNode := stack[len(stack)-1]
			res = append(res, tmpNode.Val)
			stack = stack[:len(stack)-1]
			node = tmpNode.Right
		}
	}
	return res
}