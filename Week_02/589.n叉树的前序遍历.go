/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 var result = []int{}
func preorder(root *Node) []int {
	//根左右 root left right
	result = []int{}
	dfs(root)
	return result
}

func dfs(root *Node){
	if root == nil{
		return 
	}
	result = append(result,root.Val)
	for _,n := range root.Children{
		dfs(n)
	}
}
// @lc code=end

