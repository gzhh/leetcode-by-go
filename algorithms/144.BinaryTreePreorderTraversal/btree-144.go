package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归思想 先输出当前节点，再输出左节点，最后右节点
*/
//func preorderTraversal(root *TreeNode) []int {
//	result := make([]int,0,1)
//
//	if root == nil {
//		return result
//	}
//
//	result = append(result,root.Val)
//	if root.Left != nil {
//		result = append(result,preorderTraversal(root.Left)...)
//	}
//
//
//	if root.Right != nil {
//		result = append(result,preorderTraversal(root.Right)...)
//	}
//
//	return result
//}

/**
栈实现  递归的本质就是压栈
*/
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 0)
	stack := &Stack{
		Data: make([]*TreeNode, 0, 1),
		Size: 0,
	}
	stack.Push(root)

	node := stack.Pop()
	for node != nil {
		result = append(result, node.Val)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
		node = stack.Pop()
	}

	return result
}

type Stack struct {
	Data []*TreeNode
	Size int
}

func (s *Stack) Push(val *TreeNode) bool {
	s.Data = append(s.Data, val)
	s.Size++
	return true
}

func (s *Stack) Pop() *TreeNode {
	if s.Size == 0 {
		return nil
	}
	s.Size--
	ret := s.Data[s.Size]
	s.Data = s.Data[:s.Size]
	return ret
}

func main() {
	data := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	d := preorderTraversal(data)
	fmt.Println(d)
}
