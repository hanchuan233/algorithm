package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func pathSum(root *TreeNode, targetSum int) [][]int {
	var ret [][]int
	var path []int
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right == nil && sum == targetSum {
			ret = append(ret, append([]int{}, path...))
			return
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, 0)
	return ret
}

func getTree() *TreeNode {
	node := &TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   11,
				Left:  &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	return node
}

func main() {
	root := getTree()
	ret := pathSum(root, 22)
	fmt.Println(ret)
}
