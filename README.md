# Leetcode solutions with golang

## Data Structure
* [Linked List](#linked-list)
* [Tree](#tree)
* [String](#string)

## Algorithm
* [Dynamic Programing](#dynamic-programing)
* [Depth-first Search](#depth-first-search)

## Classical Problems
* [Ugly Number](#ugly-number)
* [House Robber](#house-robber)
* [Best Time to Buy and Sell Stock](#best-time-to-buy-and-sell-stock)

## Problem List 
| # | Title | Solution | Difficulty | Importantly |
|--- | --- | ---| --- | --- |
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | [Golang](algorithms/0001.TwoSum/0001.TwoSum.go) | Easy |
| 5 | [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) | [Golang](algorithms/0005.LongestPalindromicSubstring/0005.LongestPalindromicSubstring.go) | Medium | ✅ ✅ ✅ |
| 7 | [Reverse Integer](https://leetcode.com/problems/reverse-integer/) | [Golang](algorithms/0007.ReverseInteger/0007.ReverseInteger.go) | Easy |
| 21 | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | [Golang](algorithms/0021.MergeTwoSortedLists/0021.MergeTwoSortedLists.go) | Easy | ✅ |
| 42 | [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) | [Golang](algorithms/0042.TrappingRainWater/0042.TrappingRainWater.go) | Hard | ✅ ✅ ✅ |
| 53 | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) | [Golang](algorithms/0053.MaximumSubarray/0053.MaximumSubarray.go) | Easy | ✅ ✅  |
| 55 | [Jump Game](https://leetcode.com/problems/jump-game/) | [Golang](algorithms/0055.JumpGame/0055.JumpGame.go) | Medium | ✅ ✅  |
| 62 | [Unique Paths](https://leetcode.com/problems/unique-paths/) | [Golang](algorithms/0062.UniquePaths/0062.UniquePaths.go) | Medium | ✅ ✅ |
| 63 | [Unique Paths II](https://leetcode.com/problems/unique-paths-ii/) | [Golang](algorithms/0063.UniquePathsII/0063.UniquePathsII.go) | Medium | ✅ ✅ |
| 64 | [Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/) | [Golang](algorithms/0064.MinimumPathSum/0064.MinimumPathSum.go) | Medium | ✅ ✅ |
| 70 | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) | [Golang](algorithms/0070.ClimbingStairs/0070.ClimbingStairs.go) | Easy | ✅ ✅ |
| 75 | [Sort Colors](https://leetcode.com/problems/sort-colors/) | [Golang](algorithms/0075.SortColors/0075.SortColors.go) | Medium | ✅ ✅ |
| 83 | [Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/) | [Golang](algorithms/0083.RemoveDuplicatesFromSortedList/0083.RemoveDuplicatesFromSortedList.go) | Easy | ✅ |
| 91 | [Decode Ways](https://leetcode.com/problems/decode-ways/) | [Golang](algorithms/0091.DecodeWays/0091.DecodeWays.go) | Medium | ✅ ✅ ✅ |
| 94 | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) | [Golang](algorithms/0094.BinaryTreeInorderTraversal/0094.BinaryTreeInorderTraversal.go) | Medium | ✅ ✅ |
| 95 | [Unique Binary Search Trees II](https://leetcode.com/problems/unique-binary-search-trees-ii/) | [Golang](algorithms/0095.UniqueBinarySearchTreesII/0095.UniqueBinarySearchTreesII.go) | Medium | ✅ ✅ |
| 96 | [Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/) | [Golang](algorithms/0096.UniqueBinarySearchTrees/0096.UniqueBinarySearchTrees.go) | Medium | ✅ ✅ ✅ |
| 98 | [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) | [Golang](algorithms/0098.ValidateBinarySearchTree/0098.ValidateBinarySearchTree.go) | Medium | ✅ ✅ |
| 100 | [Same Tree](https://leetcode.com/problems/same-tree/) | [Golang](algorithms/0100.SameTree/0100.SameTree.go) | Easy | ✅ ✅ |
| 101 | [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) | [Golang](algorithms/0101.SymmetricTree/0101.SymmetricTree.go) | Easy | ✅ ✅ |
| 102 | [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) | [Golang](algorithms/0102.BinaryTreeLevelOrderTraversal/0102.BinaryTreeLevelOrderTraversal.go) | Medium | ✅ ✅ |
| 104 | [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | [Golang](algorithms/0104.MaximumDepthOfBinaryTree/0104.MaximumDepthOfBinaryTree.go) | Easy | ✅ ✅ |
| 120 | [Triangle](https://leetcode.com/problems/triangle/) | [Golang](algorithms/0120.Triangle/0120.Triangle.go) | Medium | ✅ ✅ ✅ |
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [Golang](algorithms/0121.BestTimeToBuyAndSellStock/0121.BestTimeToBuyAndSellStock.go) | Easy | ✅ ✅ |
| 122 | [Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [Golang](algorithms/0122.BestTimeToBuyAndSellStockII/0122.BestTimeToBuyAndSellStockII.go) | Easy | ✅ ✅ |
| 123 | [Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/) | [Golang](algorithms/0123.BestTimeToBuyAndSellStockIII/0123.BestTimeToBuyAndSellStockIII.go) | Hard | ✅ ✅ ✅ |
| 131 | [Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/) | [Golang](algorithms/0131.PalindromePartitioning/0131.PalindromePartitioning.go) | Medium | ✅ ✅ ✅ |
| 132 | [Palindrome Partitioning II](https://leetcode.com/problems/palindrome-partitioning-ii/) | [Golang](algorithms/0132.PalindromePartitioningII/0132.PalindromePartitioningII.go) | Hard | ✅ ✅ ✅ |
| 139 | [Word Break](https://leetcode.com/problems/word-break/) | [Golang](algorithms/0139.WordBreak/0139.WordBreak.go) | Medium | ✅ ✅ ✅ |
| 141 | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) | [Golang](algorithms/0141.LinkedListCycle/0141.LinkedListCycle.go) | Easy | ✅ |
| 144 | [Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/) | [Golang](algorithms/0144.BinaryTreePreorderTraversal/0144.BinaryTreePreorderTraversal.go) | Medium | ✅ ✅ |
| 145 | [Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/) | [Golang](algorithms/0145.BinaryTreePostorderTraversal/0145.BinaryTreePostorderTraversal.go) | Hard | ✅ ✅ |
| 152 | [Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/) | [Golang](algorithms/0152.MaximumProductSubarray/0152.MaximumProductSubarray.go) | Medium | ✅ ✅ |
| 160 | [Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/) | [Golang](algorithms/0160.IntersectionOfTwoLinkedLists/0160.IntersectionOfTwoLinkedLists.go) | Easy | ✅ ✅ |
| 188 | [Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/) | [Golang](algorithms/0188.BestTimeToBuyAndSellStockIV/0188.BestTimeToBuyAndSellStockIV.go) | Hard | ✅ ✅ ✅ |
| 198 | [House Robber](https://leetcode.com/problems/house-robber/) | [Golang](algorithms/0198.HouseRobber/0198.HouseRobber.go) | Medium | ✅ ✅ |
| 203 | [Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/) | [Golang](algorithms/0203.RemoveLinkedListElements/0203.RemoveLinkedListElements.go) | Easy | ✅ |
| 206 | [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) | [Golang](algorithms/0206.ReverseLinkedList/0206.ReverseLinkedList.go) | Easy | ✅ ✅ |
| 213 | [House Robber II](https://leetcode.com/problems/house-robber-ii/) | [Golang](algorithms/0213.HouseRobberII/0213.HouseRobberII.go) | Medium | ✅ ✅ |
| 221 | [Maximal Square](https://leetcode.com/problems/maximal-square/) | [Golang](algorithms/0221.MaximalSquare/0221.MaximalSquare.go) | Medium | ✅ ✅ |
| 234 | [Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/) | [Golang](algorithms/0234.PalindromeLinkedList/0234.PalindromeLinkedList.go) | Easy | ✅ ✅ |
| 237 | [Delete Node in a Linked List](https://leetcode.com/problems/delete-node-in-a-linked-list/) | [Golang](algorithms/0237.DeleteNodeInALinkedList/0237.DeleteNodeInALinkedList.go) | Easy | |
| 263 | [Ugly Number](https://leetcode.com/problems/ugly-number/) | [Golang](algorithms/0263.UglyNumber/0263.UglyNumber.go) | Easy | ✅ ✅ ✅ |
| 264 | [Ugly Number II](https://leetcode.com/problems/ugly-number-ii/) | [Golang](algorithms/0264.UglyNumberII/0264.UglyNumberII.go) | Medium | ✅ ✅ ✅ |
| 279 | [Perfect Squares](https://leetcode.com/problems/perfect-squares/) | [Golang](algorithms/0279.PerfectSquares/0279.PerfectSquares.go) | Medium | ✅ ✅ ✅ |
| 300 | [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) | [Golang](algorithms/0300.LongestIncreasingSubsequence/0300.LongestIncreasingSubsequence.go) | Medium | ✅ ✅ ✅ |
| 309 | [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | [Golang](algorithms/0309.BestTimeToBuyAndSellStockWithCooldown/0309.BestTimeToBuyAndSellStockWithCooldown.go) | Medium | ✅ ✅ ✅ |
| 322 | [Coin Change](https://leetcode.com/problems/coin-change/) | [Golang](algorithms/0322.CoinChange/0322.CoinChange.go) | Medium | ✅ ✅ ✅ |
| 337 | [House Robber III](https://leetcode.com/problems/house-robber-iii/) | [Golang](algorithms/0337.HouseRobberIII/0337.HouseRobberIII.go) | Medium | ✅ ✅ |
| 338 | [Counting Bits](https://leetcode.com/problems/counting-bits/) | [Golang](algorithms/0338.CountingBits/0338.CountingBits.go) | Medium | ✅ ✅ |
| 344 | [Reverse String](https://leetcode.com/problems/reverse-string/) | [Golang](algorithms/0344.ReverseString/0344.ReverseString.go) | Easy | ✅ |
| 376 | [Wiggle Subsequence](https://leetcode.com/problems/wiggle-subsequence/) | [Golang](algorithms/0376.WiggleSubsequence/0376.WiggleSubsequence.go) | Medium | ✅ ✅ |
| 392 | [Is Subsequence](https://leetcode.com/problems/is-subsequence/) | [Golang](algorithms/0392.IsSubsequence/0392.IsSubsequence.go) | Easy | ✅ |
| 516 | [Longest Palindromic Subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/) | [Golang](algorithms/0516.LongestPalindromicSubsequence/0516.LongestPalindromicSubsequence.go) | Medium | ✅ ✅ ✅ |
| 647 | [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/) | [Golang](algorithms/0647.PalindromicSubstrings/0647.PalindromicSubstrings.go) | Medium | ✅ ✅ ✅ |
| 714 | [Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | [Golang](algorithms/0714.BestTimeToBuyAndSellStockWithTransactionFee/0714.BestTimeToBuyAndSellStockWithTransactionFee.go) | Medium | ✅ ✅ ✅ |
| 746 | [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/) | [Golang](algorithms/0746.MinCostClimbingStairs/0746.MinCostClimbingStairs.go) | Easy | ✅ ✅ |
| 867 | [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) | [Golang](algorithms/0876.MiddleOfTheLinkedList/0876.MiddleOfTheLinkedList.go) | Easy | ✅ |
| 1143 | [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/) | [Golang](algorithms/1143.LongestCommonSubsequence/1143.LongestCommonSubsequence.go) | Medium | ✅ ✅ ✅ |
| 1290 | [Convert Binary Number in a Linked List to Integer](https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/) | [Golang](algorithms/1290.ConvertBinaryNumberInALinkedListToInteger/1290.ConvertBinaryNumberInALinkedListToInteger.go) | Easy | ✅ ✅ |

## Classical Problems
### Ugly Number
| # | Title | Solution | Difficulty | Importantly |
|--- | --- | ---| --- | --- |
| 263 | [Ugly Number](https://leetcode.com/problems/ugly-number/) | [Golang](algorithms/0263.UglyNumber/0263.UglyNumber.go) | Easy | ✅ ✅ ✅ |
| 264 | [Ugly Number II](https://leetcode.com/problems/ugly-number-ii/) | [Golang](algorithms/0264.UglyNumberII/0264.UglyNumberII.go) | Medium | ✅ ✅ ✅ |

### House Robber
| # | Title | Solution | Difficulty | Importantly |
|--- | --- | ---| --- | --- |
| 198 | [House Robber](https://leetcode.com/problems/house-robber/) | [Golang](algorithms/0198.HouseRobber/0198.HouseRobber.go) | Medium | ✅ ✅ |
| 213 | [House Robber II](https://leetcode.com/problems/house-robber-ii/) | [Golang](algorithms/0213.HouseRobberII/0213.HouseRobberII.go) | Medium | ✅ ✅ |
| 337 | [House Robber III](https://leetcode.com/problems/house-robber-iii/) | [Golang](algorithms/0337.HouseRobberIII/0337.HouseRobberIII.go) | Medium | ✅ ✅ |

### Best Time to Buy and Sell Stock
| # | Title | Solution | Difficulty | Importantly |
|--- | --- | ---| --- | --- |
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [Golang](algorithms/0121.BestTimeToBuyAndSellStock/0121.BestTimeToBuyAndSellStock.go) | Easy | ✅ ✅ |
| 122 | [Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [Golang](algorithms/0122.BestTimeToBuyAndSellStockII/0122.BestTimeToBuyAndSellStockII.go) | Easy | ✅ ✅ |
| 123 | [Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/) | [Golang](algorithms/0123.BestTimeToBuyAndSellStockIII/0123.BestTimeToBuyAndSellStockIII.go) | Hard | ✅ ✅ ✅ |
| 188 | [Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/) | [Golang](algorithms/0188.BestTimeToBuyAndSellStockIV/0188.BestTimeToBuyAndSellStockIV.go) | Hard | ✅ ✅ ✅ |
| 309 | [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | [Golang](algorithms/0309.BestTimeToBuyAndSellStockWithCooldown/0309.BestTimeToBuyAndSellStockWithCooldown.go) | Medium | ✅ ✅ ✅ |
| 714 | [Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | [Golang](algorithms/0714.BestTimeToBuyAndSellStockWithTransactionFee/0714.BestTimeToBuyAndSellStockWithTransactionFee.go) | Medium | ✅ ✅ ✅ |


## Topic
### Dynamic Programing
| # | Title | Solution | Difficulty | Importantly |
|--- | --- | ---| --- | --- |
| 5 | [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) | [Golang](algorithms/0005.LongestPalindromicSubstring/0005.LongestPalindromicSubstring.go) | Medium | ✅ ✅ ✅ |
| 42 | [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) | [Golang](algorithms/0042.TrappingRainWater/0042.TrappingRainWater.go) | Hard | ✅ ✅ ✅ |
| 53 | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) | [Golang](algorithms/0053.MaximumSubarray/0053.MaximumSubarray.go) | Easy | ✅ ✅ |
| 55 | [Jump Game](https://leetcode.com/problems/jump-game/) | [Golang](algorithms/0055.JumpGame/0055.JumpGame.go) | Medium | ✅ ✅  |
| 62 | [Unique Paths](https://leetcode.com/problems/unique-paths/) | [Golang](algorithms/0062.UniquePaths/0062.UniquePaths.go) | Medium | ✅ ✅ |
| 63 | [Unique Paths II](https://leetcode.com/problems/unique-paths-ii/) | [Golang](algorithms/0063.UniquePathsII/0063.UniquePathsII.go) | Medium | ✅ ✅ |
| 64 | [Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/) | [Golang](algorithms/0064.MinimumPathSum/0064.MinimumPathSum.go) | Medium | ✅ ✅ |
| 70 | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) | [Golang](algorithms/0070.ClimbingStairs/0070.ClimbingStairs.go) | Easy | ✅ ✅ |
| 91 | [Decode Ways](https://leetcode.com/problems/decode-ways/) | [Golang](algorithms/0091.DecodeWays/0091.DecodeWays.go) | Medium | ✅ ✅ ✅ |
| 95 | [Unique Binary Search Trees II](https://leetcode.com/problems/unique-binary-search-trees-ii/) | [Golang](algorithms/0095.UniqueBinarySearchTreesII/0095.UniqueBinarySearchTreesII.go) | Medium | ✅ ✅ |
| 96 | [Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/) | [Golang](algorithms/0096.UniqueBinarySearchTrees/0096.UniqueBinarySearchTrees.go) | Medium | ✅ ✅ ✅ |
| 120 | [Triangle](https://leetcode.com/problems/triangle/) | [Golang](algorithms/0120.Triangle/0120.Triangle.go) | Medium | ✅ ✅ ✅ |
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [Golang](algorithms/0121.BestTimeToBuyAndSellStock/0121.BestTimeToBuyAndSellStock.go) | Easy | ✅ ✅ |
| 122 | [Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [Golang](algorithms/0122.BestTimeToBuyAndSellStockII/0122.BestTimeToBuyAndSellStockII.go) | Easy | ✅ ✅ |
| 123 | [Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/) | [Golang](algorithms/0123.BestTimeToBuyAndSellStockIII/0123.BestTimeToBuyAndSellStockIII.go) | Hard | ✅ ✅ ✅ |
| 132 | [Palindrome Partitioning II](https://leetcode.com/problems/palindrome-partitioning-ii/) | [Golang](algorithms/0132.PalindromePartitioningII/0132.PalindromePartitioningII.go) | Hard | ✅ ✅ ✅ |
| 139 | [Word Break](https://leetcode.com/problems/word-break/) | [Golang](algorithms/0139.WordBreak/0139.WordBreak.go) | Medium | ✅ ✅ ✅ |
| 152 | [Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/) | [Golang](algorithms/0152.MaximumProductSubarray/0152.MaximumProductSubarray.go) | Medium | ✅ ✅ |
| 188 | [Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/) | [Golang](algorithms/0188.BestTimeToBuyAndSellStockIV/0188.BestTimeToBuyAndSellStockIV.go) | Hard | ✅ ✅ ✅ |
| 198 | [House Robber](https://leetcode.com/problems/house-robber/) | [Golang](algorithms/0198.HouseRobber/0198.HouseRobber.go) | Medium | ✅ ✅ |
| 213 | [House Robber II](https://leetcode.com/problems/house-robber-ii/) | [Golang](algorithms/0213.HouseRobberII/0213.HouseRobberII.go) | Medium | ✅ ✅ |
| 221 | [Maximal Square](https://leetcode.com/problems/maximal-square/) | [Golang](algorithms/0221.MaximalSquare/0221.MaximalSquare.go) | Medium | ✅ ✅ |
| 264 | [Ugly Number II](https://leetcode.com/problems/ugly-number-ii/) | [Golang](algorithms/0264.UglyNumberII/0264.UglyNumberII.go) | Medium | ✅ ✅ ✅ |
| 279 | [Perfect Squares](https://leetcode.com/problems/perfect-squares/) | [Golang](algorithms/0279.PerfectSquares/0279.PerfectSquares.go) | Medium | ✅ ✅ ✅ |
| 300 | [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) | [Golang](algorithms/0300.LongestIncreasingSubsequence/0300.LongestIncreasingSubsequence.go) | Medium | ✅ ✅ ✅ |
| 309 | [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | [Golang](algorithms/0309.BestTimeToBuyAndSellStockWithCooldown/0309.BestTimeToBuyAndSellStockWithCooldown.go) | Medium | ✅ ✅ ✅ |
| 322 | [Coin Change](https://leetcode.com/problems/coin-change/) | [Golang](algorithms/0322.CoinChange/0322.CoinChange.go) | Medium | ✅ ✅ ✅ |
| 337 | [House Robber III](https://leetcode.com/problems/house-robber-iii/) | [Golang](algorithms/0337.HouseRobberIII/0337.HouseRobberIII.go) | Medium | ✅ ✅ |
| 338 | [Counting Bits](https://leetcode.com/problems/counting-bits/) | [Golang](algorithms/0338.CountingBits/0338.CountingBits.go) | Medium | ✅ ✅ |
| 376 | [Wiggle Subsequence](https://leetcode.com/problems/wiggle-subsequence/) | [Golang](algorithms/0376.WiggleSubsequence/0376.WiggleSubsequence.go) | Medium | ✅ ✅ |
| 516 | [Longest Palindromic Subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/) | [Golang](algorithms/0516.LongestPalindromicSubsequence/0516.LongestPalindromicSubsequence.go) | Medium | ✅ ✅ ✅ |
| 647 | [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/) | [Golang](algorithms/0647.PalindromicSubstrings/0647.PalindromicSubstrings.go) | Medium | ✅ ✅ ✅ |
| 714 | [Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | [Golang](algorithms/0714.BestTimeToBuyAndSellStockWithTransactionFee/0714.BestTimeToBuyAndSellStockWithTransactionFee.go) | Medium | ✅ ✅ ✅ |
| 746 | [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/) | [Golang](algorithms/0746.MinCostClimbingStairs/0746.MinCostClimbingStairs.go) | Easy | ✅ ✅ |
| 1143 | [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/) | [Golang](algorithms/1143.LongestCommonSubsequence/1143.LongestCommonSubsequence.go) | Medium | ✅ ✅ ✅ |


### Linked List
| # | Title | Solution | Difficulty | Importantly |
|--- | --- | ---| --- | --- |
| 21 | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | [Golang](algorithms/0021.MergeTwoSortedLists/0021.MergeTwoSortedLists.go) | Easy | ✅ |
| 83 | [Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/) | [Golang](algorithms/0083.RemoveDuplicatesFromSortedList/0083.RemoveDuplicatesFromSortedList.go) | Easy | ✅ |
| 141 | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) | [Golang](algorithms/0141.LinkedListCycle/0141.LinkedListCycle.go) | Easy | ✅ |
| 160 | [Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/) | [Golang](algorithms/0160.IntersectionOfTwoLinkedLists/0160.IntersectionOfTwoLinkedLists.go) | Easy | ✅ ✅ |
| 203 | [Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/) | [Golang](algorithms/0203.RemoveLinkedListElements/0203.RemoveLinkedListElements.go) | Easy | ✅ |
| 206 | [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) | [Golang](algorithms/0206.ReverseLinkedList/0206.ReverseLinkedList.go) | Easy | ✅ ✅ |
| 234 | [Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/) | [Golang](algorithms/0234.PalindromeLinkedList/0234.PalindromeLinkedList.go) | Easy | ✅ ✅ |
| 237 | [Delete Node in a Linked List](https://leetcode.com/problems/delete-node-in-a-linked-list/) | [Golang](algorithms/0237.DeleteNodeInALinkedList/0237.DeleteNodeInALinkedList.go) | Easy | |
| 867 | [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) | [Golang](algorithms/0876.MiddleOfTheLinkedList/0876.MiddleOfTheLinkedList.go) | Easy | ✅ |
| 1290 | [Convert Binary Number in a Linked List to Integer](https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/) | [Golang](algorithms/1290.ConvertBinaryNumberInALinkedListToInteger/1290.ConvertBinaryNumberInALinkedListToInteger.go) | Easy | ✅ ✅ |


### Tree
| # | Title | Solution | Difficulty | Importantly |
|--- | --- | ---| --- | --- |
| 94 | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) | [Golang](algorithms/0094.BinaryTreeInorderTraversal/0094.BinaryTreeInorderTraversal.go) | Medium | ✅ ✅ |
| 100 | [Same Tree](https://leetcode.com/problems/same-tree/) | [Golang](algorithms/0100.SameTree/0100.SameTree.go) | Easy | ✅ ✅ |
| 101 | [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) | [Golang](algorithms/0101.SymmetricTree/0101.SymmetricTree.go) | Easy | ✅ ✅ |
| 102 | [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) | [Golang](algorithms/0102.BinaryTreeLevelOrderTraversal/0102.BinaryTreeLevelOrderTraversal.go) | Medium | ✅ ✅ |
| 104 | [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | [Golang](algorithms/0104.MaximumDepthOfBinaryTree/0104.MaximumDepthOfBinaryTree.go) | Easy | ✅ ✅ |
| 116 | [Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/) | [Golang](algorithms/0116.PopulatingNextRightPointersInEachNode/0116.PopulatingNextRightPointersInEachNode.go) | Medium | ✅ ✅ |
| 144 | [Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/) | [Golang](algorithms/0144.BinaryTreePreorderTraversal/0144.BinaryTreePreorderTraversal.go) | Medium | ✅ ✅ |
| 145 | [Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/) | [Golang](algorithms/0145.BinaryTreePostorderTraversal/0145.BinaryTreePostorderTraversal.go) | Hard | ✅ ✅ |

### String
| # | Title | Solution | Difficulty | Importantly |
|--- | --- | ---| --- | --- |
| 344 | [Reverse String](https://leetcode.com/problems/reverse-string/) | [Golang](algorithms/0344.ReverseString/0344.ReverseString.go) | Easy | ✅ |
| 392 | [Is Subsequence](https://leetcode.com/problems/is-subsequence/) | [Golang](algorithms/0392.IsSubsequence/0392.IsSubsequence.go) | Easy | ✅ |

### Depth-first Search
| # | Title | Solution | Difficulty | Importantly |
|--- | --- | ---| --- | --- |
| 98 | [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) | [Golang](algorithms/0098.ValidateBinarySearchTree/0098.ValidateBinarySearchTree.go) | Medium | ✅ ✅ |
| 129 | [Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/) | [Golang](algorithms/0129.SumRootToLeafNumbers/0129.SumRootToLeafNumbers.go) | Medium | ✅ ✅ |
| 131 | [Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/) | [Golang](algorithms/0131.PalindromePartitioning/0131.PalindromePartitioning.go) | Medium | ✅ ✅ ✅ |
| 199 | [Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/) | [Golang](algorithms/0199.BinaryTreeRightSideView/0199.BinaryTreeRightSideView.go) | Medium | ✅ ✅ |
| 200 | [Number of Islands](https://leetcode.com/problems/number-of-islands/) | [Golang](algorithms/0200.NumberOfIslands/0200.NumberOfIslands.go) | Medium | ✅ ✅ |
| 257 | [Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/) | [Golang](algorithms/0257.BinaryTreePaths/0257.BinaryTreePaths.go) | Easy | ✅ ✅ |
