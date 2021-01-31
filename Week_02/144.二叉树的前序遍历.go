/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
 var result = []int{}
func preorderTraversal(root *TreeNode) []int {
	result = []int{}
	dfs(root)
	return result

}

func dfs(root *TreeNode){
	if root == nil{
		return 
	}
	// 根左右 
	result = append(result,root.Val)
	dfs(root.Left)
	dfs(root.Right)
}
// @lc code=end

