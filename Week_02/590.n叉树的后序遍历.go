/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
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
func postorder(root *Node) []int {
	//左右根 left right root 
	result = []int{}
	dfs(root)
	return result
}

func dfs(root *Node) {
	if root == nil{
		return
	}

	for _,n := range root.Children{
		dfs(n)
	}
	//最后处理root
	result = append(result, root.Val) //后序输出
}
// @lc code=end

