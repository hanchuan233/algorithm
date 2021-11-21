package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{}

	// 先构建中序遍历的哈希表
	inorderMap := make(map[int]int)
	for idx, val := range inorder {
		inorderMap[val] = idx
	}

	var makeTree func(curNode *TreeNode, curIdx int, left, right int)
	makeTree = func(curNode *TreeNode, curIdx int, left, right int) {
		curNode.Val = preorder[curIdx]
		idx := inorderMap[preorder[curIdx]]
		leftSize := idx - left
		rightSize := right - idx
		if leftSize > 0 {
			curNode.Left = &TreeNode{}
			makeTree(curNode.Left, curIdx+1, left, idx-1)
		}
		if rightSize > 0 {
			curNode.Right = &TreeNode{}
			makeTree(curNode.Right, curIdx+leftSize+1, idx+1, right)
		}
	}

	makeTree(root, 0, 0, len(inorder)-1)

	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	if root.Left == nil && root.Right != nil {
		fmt.Print("null ")
		printTree(root.Right)
	} else {
		printTree(root.Left)
		printTree(root.Right)
	}
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	ret := buildTree(preorder, inorder)
	printTree(ret)
}
