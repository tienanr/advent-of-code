package main

import (
	"fmt"
	"github.com/tienanr/advent-of-code/utils"
	"strconv"
	"strings"
)

func parseInt(s string) int {
	ret, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return ret
}

func solve1(fn string) {
	m := [][]string{}
	for line := range utils.ReadFile(fn) {
		row := strings.Fields(line)
		m = append(m, row)
	}

	w := len(m[0])
	h := len(m)
	ans := 0
	for i := 0; i < w; i++ {
		op := m[h-1][i]
		switch op {
		case "+":
			for j := 0; j < h-1; j++ {
				ans += parseInt(m[j][i])
			}
		case "*":
			tmp := 1
			for j := 0; j < h-1; j++ {
				tmp *= parseInt(m[j][i])
			}
			ans += tmp
		}
	}

	fmt.Println("part 1 solution is:", ans)
}

func solve2(fn string) {
	m := [][]rune{}
	for line := range utils.ReadFile(fn) {
		m = append(m, []rune(line))
	}

	w := len(m[0])
	h := len(m)
	ans := 0

	var op rune
	var tmp int
	for i := 0; i < w; i++ {
		line := []rune{}
		for j := 0; j < h-1; j++ {
			line = append(line, m[j][i])
		}
		if m[h-1][i] != ' ' {
			op = m[h-1][i]
			// determine `zero` for op
			switch op {
			case '+':
				tmp = 0
			case '*':
				tmp = 1
			}
		}

		s := strings.TrimSpace(string(line))
		if len(s) > 0 {
			switch op {
			case '+':
				tmp += parseInt(s)
			case '*':
				tmp *= parseInt(s)
			}
		} else {
			ans += tmp //collect result
		}
	}
	ans += tmp
	fmt.Println("part 2 solution is:", ans)
}

func main() {
	fn := "example.txt"
	solve1(fn)
	solve2(fn)
}
