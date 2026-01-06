package main

import (
	"fmt"
	"github.com/tienanr/advent-of-code/utils"
	"slices"
)

func solve(fn string) {
	coordinates := [][]int{}
	for line := range utils.ReadFile(fn) {
		coordinates = append(coordinates, utils.ParseInts(line))
	}

	rectangles := [][][]int{}
	for i, p := range coordinates {
		for _, q := range coordinates[i+1:] {
			rectangles = append(rectangles, [][]int{p, q})
		}
	}

	slices.SortFunc(rectangles, func(r1, r2 [][]int) int {
		return area(r2) - area(r1)
	})

	fmt.Println("part 1 solution: ", area(rectangles[0]))

	// collect all edges
	edges := [][][]int{}
	for i, p := range coordinates {
		j := (i + 1) % len(coordinates)
		q := coordinates[j]
		edges = append(edges, [][]int{p, q})
	}

	// incomplete solution: only detect if any edge go across rectangles
	for _, r := range rectangles {
		found := true
		for _, e := range edges {
			if intersect(r, e) {
				found = false
				break
			}
		}

		if found {
			fmt.Println("part 2 solution: ", area(r))
			break
		}
	}
}

func area(r [][]int) int {
	w := abs(r[0][0]-r[1][0]) + 1
	h := abs(r[0][1]-r[1][1]) + 1

	return w * h
}

func intersect(r [][]int, e [][]int) bool {
	minX := min(r[0][0], r[1][0])
	maxX := max(r[0][0], r[1][0])
	minY := min(r[0][1], r[1][1])
	maxY := max(r[0][1], r[1][1])

	eMinX := min(e[0][0], e[1][0])
	eMaxX := max(e[0][0], e[1][0])
	eMinY := min(e[0][1], e[1][1])
	eMaxY := max(e[0][1], e[1][1])

	edgeOnLeft := eMaxX <= minX
	edgeOnRight := eMinX >= maxX
	edgeOnTop := eMaxY <= minY
	edgeOnBottom := eMinY >= maxY

	return !(edgeOnLeft || edgeOnRight || edgeOnTop || edgeOnBottom)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	solve("example.txt")
}
