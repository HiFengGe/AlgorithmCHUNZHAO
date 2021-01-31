/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	//左-根-右
	result := []int{}
	var dfs func(*TreeNode) 
	
    dfs = func(root *TreeNode) {
        if root == nil {return}
        dfs(root.Left)
        result = append(result, root.Val)
        dfs(root.Right)
    }
    dfs(root)
    
    return result;
}

// @lc code=end

