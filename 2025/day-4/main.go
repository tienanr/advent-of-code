package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solvePart1(m [][]rune) int {
	// check accessable rolls
	w := len(m[0])
	h := len(m)
	ans := 0

	for y := 0; y < h; y++ { // check all cells
		for x := 0; x < w; x++ {
			if m[y][x] == '@' {
				count := -1 // discount self
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if 0 <= x+dx && x+dx < w && 0 <= y+dy && y+dy < h && m[y+dy][x+dx] == '@' {
							count++
						}
					}
				}
				if count < 4 {
					ans++
				}
			}
		}
	}
	return ans
}

func solvePart2(m [][]rune) int {
	// check accessable rolls
	w := len(m[0])
	h := len(m)
	ans := 0
	stop := false

	for !stop {
		stop = true
		for y := 0; y < h; y++ { // check all cells
			for x := 0; x < w; x++ {
				if m[y][x] == '@' {
					count := -1 // discount self
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							if 0 <= x+dx && x+dx < w && 0 <= y+dy && y+dy < h && m[y+dy][x+dx] == '@' {
								count++
							}
						}
					}
					if count < 4 {
						m[y][x] = 'x'
						ans++
						stop = false
					}
				}
			}
		}
	}
	return ans
}

func solve(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// initialize map
	m := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, []rune(line))
	}

	fmt.Println("part 1 solution:", solvePart1(m))
	fmt.Println("part 2 solution:", solvePart2(m))
}

func main() {
	solve("question.txt")
}
