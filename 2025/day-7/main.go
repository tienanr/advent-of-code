package main

import (
	"fmt"
	"github.com/tienanr/advent-of-code/utils"
)

func main() {
	// load map
	m := [][]rune{}
	for line := range utils.ReadFile("example.txt") {
		m = append(m, []rune(line))
	}
	w := len(m[0])
	h := len(m)

	// part 1 - simulation
	ans1 := 0
	for j := 0; j < h-1; j++ {
		for i := 0; i < w; i++ {
			switch m[j][i] {
			case 'S', '|':
				switch m[j+1][i] {
				case '.':
					m[j+1][i] = '|'
				case '^':
					m[j+1][i-1] = '|'
					m[j+1][i+1] = '|'
					ans1++
				}
			}
		}
	}
	fmt.Println("part 1 solution: ", ans1)

	// part 2 - dynamic programming
	s := make([][]int, h)
	for j := range s {
		s[j] = make([]int, w)
	}

	var ans2 int
	for j := h - 1; j >= 0; j-- {
		for i := 0; i < w; i++ {
			if j == h-1 { // no split at bottom
				s[j][i] = 1
			} else {
				if m[j+1][i] == '^' { //split!
					s[j][i] = s[j+1][i-1] + s[j+1][i+1]
				} else {
					s[j][i] = s[j+1][i]
				}
			}
			if m[j][i] == 'S' {
				ans2 = s[j][i]
			}
		}
	}
	fmt.Println("part 2 solution: ", ans2)
}
