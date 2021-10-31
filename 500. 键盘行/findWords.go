package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	keys := "qwertyuiopasdfghjklzxcvbnm"

	keyMap := make(map[byte]int, len(keys))
	for i := 0; i < len(keys); i++ {
		keyMap[keys[i]] = i
	}

	var ret []string
	for _, word := range words {
		w := strings.ToLower(word)
		line := getLine(keyMap[w[0]])
		for i := 1; i < len(w); i++ {
			if getLine(keyMap[w[i]]) != line {
				goto NOT
			}
		}
		ret = append(ret, word)
	NOT:
	}

	return ret
}

func getLine(idx int) int {
	line1 := 10
	line2 := 19
	if idx < line1 {
		return 1
	} else if idx < line2 {
		return 2
	} else {
		return 3
	}
}

func main() {
	words := []string{"Hello","Alaska","Dad","Peace"}
	ret := findWords(words)
	fmt.Println(ret)
}
