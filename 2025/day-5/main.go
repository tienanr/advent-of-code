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
	for line := range utils.ReadFile("example.txt") {
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
	for i:=0; i<len(ranges); {
		// combine all ranges that overlaps starting range[i]
		lo := ranges[i].Low
		hi := ranges[i].High
		for i++;i < len(ranges) && hi >= ranges[i].Low; i++ {
			if ranges[i].High > hi {
				hi = ranges[i].High
			}
		}
		// count valid ids between lo and hi
		ans2 += hi - lo + 1
	}
	fmt.Println("Part 2 solution is: ", ans2)
}
