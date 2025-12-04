package main

import (
    "github.com/tienanr/advent-of-code/utils"
    "fmt"
)

func parseLine(line string) []int {
    bank := []int{}
    for _, ch := range line {
        bank = append(bank, int(ch - '0'))
    }
    return bank
}

func solveBank(bank []int, n int) int {
    ans, pos := 0, 0

    for n>0 {
        for i:=pos+1; i<len(bank)-n+1; i++ {
            if bank[i] > bank[pos] {
                pos = i
            }
        }
        ans = ans * 10 + bank[pos]
        n--
        pos++
    }
    return ans
}

func solve(fn string) {
    ans1, ans2 := 0, 0
    for line := range utils.ReadFile(fn) {
        bank := parseLine(line)
        ans1 += solveBank(bank, 2)
        ans2 += solveBank(bank, 12)
    }

    fmt.Println("part 1 answer is:", ans1)
    fmt.Println("part 2 answer is:", ans2)
}

func main() {
    solve("example.txt")
}
