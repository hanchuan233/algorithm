package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	num := 0
	var getNumber func(node *TreeNode, num int)
	getNumber = func(node *TreeNode, num int) {
		if node.Left == nil && node.Right == nil {
			sum += num*10 + node.Val
			return
		}
		num = num*10 + node.Val
		defer func() {
			num /= 10
		}()
		if node.Left != nil {
			getNumber(node.Left, num)
		}

		if node.Right != nil {
			getNumber(node.Right, num)
		}
	}

	getNumber(root, num)

	return sum
}

func MakeTree() *TreeNode {
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
	root := MakeTree()
	ret := sumNumbers(root)
	fmt.Println(ret)
}
