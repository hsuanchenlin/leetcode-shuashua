/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (25.57%)
 * Likes:    1829
 * Dislikes: 275
 * Total Accepted:    392.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[2,1,3]'
 *
 * Given a binary tree, determine if it is a valid binary search tree (BST).
 * 
 * Assume a BST is defined as follows:
 * 
 * 
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 
 * Input: [2,1,3]
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 
 * Input: [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is 4.
 * 
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
   return helper(root, math.MaxInt64, math.MinInt64)
}

func helper(root *TreeNode, max, min int) bool {
    if root.Val >= max || root.Val <= min {
        return false
    }
    left := true
    if root.Left != nil {
        left = helper(root.Left, root.Val, min)
    }
    right := true
    if root.Right != nil {
        right = helper(root.Right, max, root.Val)
    }
    return left && right  
}
