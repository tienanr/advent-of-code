package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isInvalidV2(n int) bool {
	s := strconv.Itoa(n)
	m := len(s)

	for _, factor := range calculateFactors(m) {
		seg := s[:factor]
		if strings.Repeat(seg, m/factor) == s {
			return true
		}
	}
	return false
}

func calculateFactors(n int) []int {
	fs := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			fs = append(fs, i)
		}
	}
	return fs
}

func isInvalidV1(n int) bool {
	s := strconv.Itoa(n)
	m := len(s)

	return s[:m/2] == s[m/2:]
}

func solveRange(low int, hi int) (int, int) {
	ans1, ans2 := 0, 0

	for i := low; i <= hi; i++ {
		if isInvalidV1(i) {
			ans1 += i
		}
		if isInvalidV2(i) {
			ans2 += i
		}
	}

	return ans1, ans2
}

func solve(fn string) (int, int) {
	content, _ := os.ReadFile(fn)
	rgs := strings.Split(string(content), ",")

	ans1, ans2 := 0, 0
	for _, rg := range rgs {
		pts := strings.Split(rg, "-")

		low, err := strconv.Atoi(strings.TrimSpace(pts[0]))
		if err != nil {
			log.Fatal(err)
		}

		high, err := strconv.Atoi(strings.TrimSpace(pts[1]))
		if err != nil {
			log.Fatal(err)
		}

		d1, d2 := solveRange(low, high)

		ans1 += d1
		ans2 += d2
	}

	return ans1, ans2
}

func main() {
	fmt.Println(solve("example.txt"))
}
