// Source : https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
// Author : Zhonghuan Gao
// Date   : 2020-11-28

/**********************************************************************************
*
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
Initially, all next pointers are set to NULL.

Follow up:
You may only use constant extra space.
Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.

Example 1:
Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

Constraints:
The number of nodes in the given tree is less than 4096.
-1000 <= node.val <= 1000
*
***********************************************************************************/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

/**
题意：将一个完全二叉树每层的节点用链表连接起来

解法一：层次遍历
使用一个队列来来存储每层的节点，将每层的节点链接起来，依次拓展下一层的新队列。
复杂度，时间O(N) 空间O(N)
 */
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		tmpQueue := queue
		queue = nil
		for i, node := range tmpQueue {
			if i < len(tmpQueue) - 1 {
				node.Next = tmpQueue[i + 1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

/**
解法二：使用已建立的 next 指针
复杂度，时间O(N) 空间O(1)
参考：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/solution/tian-chong-mei-ge-jie-dian-de-xia-yi-ge-you-ce-2-4/
 */
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}

	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		for node := leftmost; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	return root
}

/**
解法三：拉拉链解法（递归）
从根节点开始遍历，然后依次遍历每个节点的左右子树，每次只处理当前树的每层中心两个节点的链接
参考：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/solution/la-la-lian-jie-fa-by-sorcerer/
 */
func connect3(root *Node) *Node {
	if root == nil {
		return root
	}

	left := root.Left
	right := root.Right
	for left != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}
	connect(root.Left)
	connect(root.Right)

	return root
}

