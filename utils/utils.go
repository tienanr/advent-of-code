package utils

import (
    "bufio"
    "os"
    "log"
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
