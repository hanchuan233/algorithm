package main

import "fmt"

type TrieNode struct {
	childMap map[interface{}]*TrieNode
	val      int
	isEnd    bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		childMap: make(map[interface{}]*TrieNode),
		val:      0,
		isEnd:    false,
	}
}

type MapSum struct {
	root *TrieNode
	cnt  map[string]int
}

func Constructor() MapSum {
	return MapSum{
		root: newTrieNode(),
		cnt:  make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	// 存储key的数量
	delta := val
	if cnt, has := this.cnt[key]; has {
		delta -= cnt
	}

	curNode := this.root
	for i := 0; i < len(key); i++ {
		if _, has := curNode.childMap[key[i]]; !has {
			curNode.childMap[key[i]] = newTrieNode()
		}
		curNode = curNode.childMap[key[i]]
		curNode.val += delta
	}

	curNode.isEnd = true
	this.cnt[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	curNode := this.root
	for i := 0; i < len(prefix); i++ {
		if _, has := curNode.childMap[prefix[i]]; !has {
			return 0
		}
		curNode = curNode.childMap[prefix[i]]
	}

	return curNode.val
}

func main() {
	key1 := "apple"
	val1 := 3

	key2 := "app"
	val2 := 2

	key3 := "apple"
	val3 := 1

	prefix := "ap"

	ms := Constructor()
	ms.Insert(key1, val1)
	fmt.Println(ms.Sum(prefix))
	ms.Insert(key2, val2)
	fmt.Println(ms.Sum(prefix))
	ms.Insert(key3, val3)
	fmt.Println(ms.Sum(prefix))
}
