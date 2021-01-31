/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 前序 根左-右（根左右） preorder 
	// 中序 左根-右（左根右） inorder

	if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    // 中顺序列找根结点
    var root int
    for k, v :=  range inorder {
        if v == preorder[0] {
            root = k
            break
        }
    }

    // 左右子树归类
    // pre_left, pre_right := preorder[1: root+1], preorder[root+1:]
    // in_left, in_right := inorder[0: root], inorder[root+1:]
    
    // 左右子树递归
    return &TreeNode{
		Val:   preorder[0],
						// Next 
						// preorder => left=1:root+1, right = [root+1:]
						// inorder => left=0:root, right = [root+1:]
        Left:  buildTree(preorder[1: root+1], inorder[0: root]),
        Right: buildTree(preorder[root+1:], inorder[root+1:]),
    }
}
// @lc code=end

