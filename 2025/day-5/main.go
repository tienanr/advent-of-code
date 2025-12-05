package main

import (
	"fmt"
	"github.com/tienanr/advent-of-code/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Low  uint64
	High uint64
}

func parseId(s string) uint64 {
	ret, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func parseRange(s string) Range {
	parts := strings.Split(s, "-")
	lo := parseId(parts[0])
	hi := parseId(parts[1])
	return Range{Low: lo, High: hi}
}

func main() {
	ranges := []Range{}
	ids := []uint64{}

	// input parsing
	parsing_ranges := true
	for line := range utils.ReadFile("question.txt") {
		if parsing_ranges {
			if len(line) == 0 {
				parsing_ranges = false
			} else {
				ranges = append(ranges, parseRange(line))
			}
		} else {
			ids = append(ids, parseId(line))
		}
	}

	// solve part 1
	ans1 := 0
	for _, id := range ids {
		for _, rg := range ranges {
			if rg.Low <= id && id <= rg.High {
				ans1++
				break
			}
		}
	}
	fmt.Println("Part 1 solution is: ", ans1)

	// solve part 2
	ans2 := uint64(0)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Low < ranges[j].Low
	})
	i, j := 0, 0
	for i < len(ranges) {
		// find the last range with overlap
		lo := ranges[i].Low
		hi := ranges[i].High
		for j+1 < len(ranges) && hi >= ranges[j+1].Low {
			j++
			if ranges[j].High > hi {
				hi = ranges[j].High
			}
		}
		// count valid ids between range[i] and range[j]
		ans2 += hi - lo + 1
		// move cursor to larger ranges
		i, j = j+1, j+1
	}
	fmt.Println("Part 2 solution is: ", ans2)
}
