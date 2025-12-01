package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func countZero(rot int) int {
	if rot%100 == 0 {
		return 1
	}
	return 0
}

func zerosBetween(rot int, nrot int) int {
	if nrot > 0 {
		return nrot / 100
	}
	if nrot < 0 && rot == 0 {
		return -nrot / 100
	}
	if nrot < 0 {
		return -nrot/100 + 1
	}
	return 1
}

var Pattern = regexp.MustCompile(`(\w)(\d+)`)

func parseLine(line string) (int, error) {
	match := Pattern.FindStringSubmatch(line)
	if match == nil {
		return 0, fmt.Errorf("invalid input: %s", line)
	}
	n, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, fmt.Errorf("invalid input: %s", line)
	}
	switch match[1] {
	case "L":
		return -n, nil
	case "R":
		return n, nil
	default:
		return 0, fmt.Errorf("invalid input: %s", line)
	}
}

func solve(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rot, ans1, ans2 := 50, 0, 0
	for scanner.Scan() {
		delta, err := parseLine(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		nrot := rot + delta
		ans1 += countZero(nrot)
		ans2 += zerosBetween(rot, nrot)
		rot = (nrot%100 + 100) % 100
	}
	return ans1, ans2
}

func main() {
	ans1, ans2 := solve("example.txt")
	fmt.Println("part 1 solution is:", ans1)
	fmt.Println("part 2 solution is:", ans2)
}
