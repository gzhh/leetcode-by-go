// Source : https://leetcode.com/problems/minimum-depth-of-binary-tree/
// Author : Zhonghuan Gao
// Date   : 2022-03-06

/**********************************************************************************
*
Given a binary tree, find its minimum depth.
The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
Note: A leaf is a node with no children.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 2

Example 2:
Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5

Constraints:
The number of nodes in the tree is in the range [0, 105].
-1000 <= Node.val <= 1000
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := minDepth(root.Left)
	rDepth := minDepth(root.Right)

	// 当一个左子树为空，右不为空，这时并不是最低点
	if root.Left == nil && root.Right != nil {
		return rDepth + 1
	}
	// 当一个右子树为空，左不为空，这时并不是最低点
	if root.Left != nil && root.Right == nil {
		return lDepth + 1
	}

	if lDepth < rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}

/**
解法二：层级遍历
*/
func minDepth2(root *TreeNode) int {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	res := 0
	for len(queue) > 0 {
		len := len(queue)
		res++
		for i := 0; i < len; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return res
			}
		}
	}
	return res
}
