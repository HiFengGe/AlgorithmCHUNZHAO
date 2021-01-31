/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil{
		return [][]int{}
	}
	result := [][]int{}
	dfs(root,0,&result)//层级 及传地址
	return result
}

func dfs(root *Node,level int,result *[][]int){
	if len(*result) <= level{ //0
		(*result) = append((*result),[]int{root.Val})
	}else{
		//合并叠加对应的层
		(*result)[level] = append((*result)[level],root.Val)
	}
	
	if root.Children != nil{
		for _,n := range root.Children {
			dfs(n,level+1,result)
		}
	}
	
}
// @lc code=end

