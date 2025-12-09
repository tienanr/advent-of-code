package utils

import (
    "bufio"
    "os"
    "log"
	"strconv"
)

func ReadFile(fn string) <-chan string {
    ch := make(chan string)

    go func() {
        file, err := os.Open(fn)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            ch <- scanner.Text()
        }
        close(ch)
    }()

    return ch
}

func ParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
