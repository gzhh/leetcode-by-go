// Source : https://leetcode.com/problems/symmetric-tree/
// Author : Zhonghuan Gao
// Date   : 2020-07-29

/**********************************************************************************
*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Follow up: Solve it both recursively and iteratively.
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSubSymmetric(root.Left, root.Right)
}

func isSubSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) || (left.Val != right.Val) {
		return false
	}
	return isSubSymmetric(left.Left, right.Right) && isSubSymmetric(right.Left, left.Right)
}


/**
解法二：迭代
用两个队列来判断左右子树是否对称
 */
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var left, right *TreeNode
	q1, q2 := []*TreeNode{}, []*TreeNode{}
	q1 = append(q1, root.Left)
	q2 = append(q2, root.Right)
	for len(q1) > 0 && len(q2) > 0 {
		left = q1[0]
		q1 = q1[1:]
		right = q2[0]
		q2 = q2[1:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil  || left.Val != right.Val {
			return false
		}
		q1 = append(q1, left.Left)
		q1 = append(q1, left.Right)
		q2 = append(q2, right.Right)
		q2 = append(q2, right.Left)
	}

	return true
}


/**
解法三：综合
先翻转左子树，然后再与右子树比较
参考第 226 题翻转二叉树，第 100 题判断两个二叉树是否相等
 */