package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// BSTIterator 中序遍历
//type BSTIterator struct {
//	root *TreeNode
//	arr  []int
//}
//
//func Constructor(root *TreeNode) BSTIterator {
//	it := BSTIterator{
//		root: root,
//	}
//	var inorder func(node *TreeNode)
//	inorder = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//
//		inorder(node.Left)
//		it.arr = append(it.arr, node.Val)
//		inorder(node.Right)
//	}
//	inorder(root)
//	return it
//}
//
//func (this *BSTIterator) Next() int {
//	ret := this.arr[0]
//	this.arr = this.arr[1:]
//	return ret
//}
//
//func (this *BSTIterator) HasNext() bool {
//	return len(this.arr) > 0
//}

// BSTIterator 迭代
type BSTIterator struct {
	root *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator {
		root: root,
	}
	node := root
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
	return it
}

func (this *BSTIterator) Next() int {
	n := len(this.stack)
	retNode := this.stack[n-1]
	this.stack = this.stack[0:n-1]

	node := retNode.Right
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}

	return retNode.Val
}


func (this *BSTIterator) HasNext() bool {
	if len(this.stack) == 0 {
		return false
	} else {
		return true
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
