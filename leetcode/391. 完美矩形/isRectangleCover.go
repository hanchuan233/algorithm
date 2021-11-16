package main

import "fmt"

type point struct {
	x, y int
}

func isRectangleCover(rectangles [][]int) bool {
	// 左下角的点和右上角的点组成的矩形面积等于所有矩形的面积总和
	// 并且左下角的点的x, y都是最小的；右上角的点的x, y都是最大的
	// 所有顶点只可能出现过1次
	// 所有除顶点外的点只可能出现2或4次
	pointMap := make(map[point]int)
	rectangle := rectangles[0]
	area := 0
	for i := 0; i < len(rectangles); i++ {
		area += getArea(rectangles[i])

		points := getPoints(rectangles[i])

		for _, point := range points {
			pointMap[point] += 1
		}

		if rectangles[i][0] <= rectangle[0] && rectangles[i][1] <= rectangle[1] {
			rectangle[0] = rectangles[i][0]
			rectangle[1] = rectangles[i][1]
		}

		if rectangles[i][2] >= rectangle[2] && rectangles[i][3] >= rectangle[3] {
			rectangle[2] = rectangles[i][2]
			rectangle[3] = rectangles[i][3]
		}
	}
	// 求最外层矩形的四个顶点
	outsidePoints := getPoints(rectangle)
	for _, point := range outsidePoints {
		if pointMap[point] != 1 {
			return false
		}
		delete(pointMap, point)
	}

	for _, cnt := range pointMap {
		if cnt != 2 && cnt != 4 {
			return false
		}
	}

	sumArea := getArea(rectangle)
	if area == sumArea {
		return true
	} else {
		return false
	}

}

func getArea(rectangle []int) int {
	return (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])
}

func getPoints(rectangle []int) []point {
	pointA := point{x: rectangle[0], y: rectangle[1]} // 左下角
	pointB := point{x: rectangle[2], y: rectangle[1]} // 右下角
	pointC := point{x: rectangle[2], y: rectangle[3]} // 右上角
	pointD := point{x: rectangle[0], y: rectangle[3]} // 左上角
	return []point{pointA, pointB, pointC, pointD}
}

func main() {
	rectangles := [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}
	ret := isRectangleCover(rectangles)
	fmt.Println(ret)
}
