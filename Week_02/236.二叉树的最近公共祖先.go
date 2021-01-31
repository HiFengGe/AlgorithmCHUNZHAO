/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// （1） 如果当前结点 root root root 等于 NULL，则直接返回 NULL
	// （2） 如果 root root root 等于 ppp 或者 qqq ，那这棵树一定返回 ppp 或者 qqq
	// （3） 然后递归左右子树，因为是递归，使用函数后可认为左右子树已经算出结果，用 left left left 和 right right right 表示
	// （4） 此时若left left left为空，那最终结果只要看 right right right；若 right right right 为空，那最终结果只要看 left left left
	// （5） 如果 left left left 和 right right right 都非空，因为只给了 ppp 和 qqq 两个结点，都非空，说明一边一个，因此 root root root 是他们的最近公共祖先
	// （6） 如果 left left left 和 right right right 都为空，则返回空（其实已经包含在前面的情况中了）
	
	//如果为空
	if(root == nil){
		return nil
	}


	//根左右
	//如果是根
	if(root == p || root == q) {
		return root
	}
		
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if(left == nil){
		return right
	}

	if(right == nil){
		return left
	}

	// p和q在两侧
	if(left != nil && right != nil) { 
		return root
	}

	return nil; // 必须有返回值
}
// @lc code=end

