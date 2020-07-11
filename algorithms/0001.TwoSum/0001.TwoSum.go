// Source : https://leetcode.com/problems/two-sum/
// Author : Zhonghuan Gao
// Date   : 2019-05-06

/********************************************************************************** 
*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*               
***********************************************************************************/

package leetcode

/**
 * 解法一：暴力法
 * 双层循环每两个不同元素相加的值与target比较
 */
func twoSum(nums []int, target int) []int {
    var res []int;
    for i, vi := range nums {
        for j, vj := range nums {
            if (i != j && vi + vj == target) {
                return append(res, i, j)
            }
        }
    }
    return res
}

/**
 * 解法二：哈希表
 * 利用map储存每个数组元素的值和索引，当target减掉当前数组元素的值在哈希表中存在的话，
 * 得出两个数组元素相加等于target，最后返回两个元素在数组中的索引
 */
func twoSumS2(nums []int, target int) []int {
    mp := make(map[int]int)
    
    for i, num := range(nums) {
        if idx, ok := mp[target - num]; ok {
            return []int{i, idx}
        }
        mp[num] = i
    }
    
    return []int{}
}
