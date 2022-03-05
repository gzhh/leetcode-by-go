// Source : https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
// Author : Zhonghuan Gao
// Date   : 2022-03-05

/**********************************************************************************
*
Evaluate the value of an arithmetic expression in Reverse Polish Notation.
Valid operators are +, -, *, and /. Each operand may be an integer or another expression.
Note that division between two integers should truncate toward zero.
It is guaranteed that the given RPN expression is always valid. That means the expression would always evaluate to a result, and there will not be any division by zero operation.

Example 1:
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Example 2:
Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6

Example 3:
Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

Constraints:
1 <= tokens.length <= 104
tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
*
***********************************************************************************/

package leetcode

/**
解法一：数组模拟栈
*/
func evalRPN(tokens []string) int {
	stack := []string{}
	for _, s := range tokens {
		n := len(stack)
		if s == "+" {
			top1, _ := strconv.Atoi(stack[n-1])
			top2, _ := strconv.Atoi(stack[n-2])
			stack = stack[:n-2]
			stack = append(stack, strconv.Itoa(top1+top2))
		} else if s == "-" {
			top1, _ := strconv.Atoi(stack[n-1])
			top2, _ := strconv.Atoi(stack[n-2])
			stack = stack[:n-2]
			stack = append(stack, strconv.Itoa(top2-top1))
		} else if s == "*" {
			top1, _ := strconv.Atoi(stack[n-1])
			top2, _ := strconv.Atoi(stack[n-2])
			stack = stack[:n-2]
			stack = append(stack, strconv.Itoa(top1*top2))
		} else if s == "/" {
			top1, _ := strconv.Atoi(stack[n-1])
			top2, _ := strconv.Atoi(stack[n-2])
			stack = stack[:n-2]
			stack = append(stack, strconv.Itoa(top2/top1))
		} else {
			stack = append(stack, s)
		}
	}
	ans, _ := strconv.Atoi(stack[0])
	return ans
}
