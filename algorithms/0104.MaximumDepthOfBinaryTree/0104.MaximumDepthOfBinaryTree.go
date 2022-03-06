// Source : https://leetcode.com/problems/maximum-depth-of-binary-tree/
// Author : Zhonghuan Gao
// Date   : 2020-07-23

/**********************************************************************************
*
Given a binary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
Note: A leaf is a node with no children.

Example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
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

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
解法一：递归
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}

/**
解法二：层级遍历
*/
func maxDepth2(root *TreeNode) int {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	res := 0
	for len(queue) > 0 {
		len := len(queue)
		for i := 0; i < len; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res++
	}
	return res
}
