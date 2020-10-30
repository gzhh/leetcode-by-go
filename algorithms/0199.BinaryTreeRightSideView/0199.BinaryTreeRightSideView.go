// Source : https://leetcode.com/problems/binary-tree-right-side-view/
// Author : Zhonghuan Gao
// Date   : 2020-10-30

/**********************************************************************************
*
Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes
you can see ordered from top to bottom.

Example:
Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
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
解法一：dfs
按照根结点->右子树->左子树的顺序访问，保证每层都是最先访问最右边的节点
定义全局结果列表，从根节点开始访问，先访问当前节点你，再递归地访问右子树和左子树；如果当前节点所在的深度还没有出现在res里，
说明在该深度下当前节点是第一个被访问的节点，因此将当前节点加入res中
 */
var res []int

func rightSideView(root *TreeNode) []int {
	res = make([]int, 0)
	if root == nil {
		return nil
	}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if depth == len(res) {
		res = append(res, root.Val)
	}
	dfs(root.Right, depth + 1)
	dfs(root.Left, depth + 1)
}