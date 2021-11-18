package main

import (
	"fmt"
	_ "kruskal"
	"prim"
)

func main() {
	points := [][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}
	//ret := kruskal.MinCostConnectPoints(points)
	ret := prim.MinCostConnectPoints(points)
	fmt.Print(ret)
}

