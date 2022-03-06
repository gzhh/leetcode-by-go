// Source : https://leetcode.com/problems/binary-tree-level-order-traversal/
// Author : Zhonghuan Gao
// Date   : 2020-07-26

/**********************************************************************************
*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7

return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
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
解法一：非递归(BFS)
用栈模拟递归过程
建立一个 queue，先把根节点放进去，当 queue 不为空时一直循环，循环内新建一个数组，把跟节点的值放进去，
然后找根节点的左右两个子节点，放入 queue 中，并去掉根节点，此时 queue 里的元素就是下一层的所有节点，
用一个 for 循环遍历它们，然后存到上面新建的数组 里，遍历完之后再把这个数组 存到全局二维 vector 里，
以此类推，可以完成层序遍历

层序遍历变种题：
102. 二叉树的层序遍历
107. 二叉树的层次遍历 II
199. 二叉树的右视图
429. N 叉树的层序遍历
116. 填充每个节点的下一个右侧节点指针
117. 填充每个节点的下一个右侧节点指针 II
104. 二叉树的最大深度
111. 二叉树的最小深度
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue, res := []*TreeNode{root}, [][]int{}

	for len(queue) > 0 {
		curLevel := []int{}
		for i := len(queue); i > 0; i-- {
			tmpNode := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, tmpNode.Val)
			if tmpNode.Left != nil {
				queue = append(queue, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				queue = append(queue, tmpNode.Right)
			}
		}
		res = append(res, curLevel)
	}

	return res
}

func levelOrder1(root *TreeNode) [][]int {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	res := [][]int{}
	for len(queue) > 0 {
		len := len(queue)
		item := []int{}
		for i := 0; i < len; i++ {
			node := queue[0]
			queue = queue[1:]
			item = append(item, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, item)
	}
	return res
}

/**
解法二：递归(DFS)
*/
func dfs(res *[][]int, root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(*res) == depth {
		*res = append(*res, []int{})
	}
	(*res)[depth] = append((*res)[depth], root.Val)
	dfs(res, root.Left, depth+1)
	dfs(res, root.Right, depth+1)
}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	dfs(&res, root, 0)
	return res
}
