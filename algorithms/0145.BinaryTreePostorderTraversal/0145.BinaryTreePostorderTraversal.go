// Source : https://leetcode.com/problems/binary-tree-postorder-traversal/
// Author : Zhonghuan Gao
// Date   : 2020-07-22

/**********************************************************************************
*
Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]

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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
解法一：递归
*/
func postorderTraversal(root *TreeNode) []int {
	var res []int
	rec(root, &res)
	return res
}

func rec(root *TreeNode, output *[]int) {
	if root != nil {
		rec(root.Left, output)
		rec(root.Right, output)
		*output = append(*output, root.Val)
	}
}

/**
解法一：递归（匿名函数）
*/
func postorderTraversal1(root *TreeNode) []int {
	var res []int
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}

/**
解法二：
用栈模拟递归过程
*/
func postorderTraversal2(root *TreeNode) []int {
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
		res = append([]int{node.Val}, res...) // 前插
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	return res
}

func postorderTraversal3(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    stack := []*TreeNode{}
    var res []int
    stack = append(stack, root)
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, node.Val)
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
    }
    n := len(res)
    for i := 0; i < n / 2; i++ {
        res[i], res[n-1-i] = res[n-1-i], res[i]
    }
    return res
