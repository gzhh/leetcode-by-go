// Source : https://leetcode.com/problems/binary-tree-preorder-traversal/
// Author : Zhonghuan Gao
// Date   : 2020-07-22

/**********************************************************************************
*
Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]

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
func preorderTraversal(root *TreeNode) []int {
	var res []int
	rec(root, &res)
	return res
}

func rec(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		rec(root.Left, output)
		rec(root.Right, output)
	}
}

/**
解法二：
用栈模拟递归过程
 */
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*TreeNode{}, []int{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return res
}