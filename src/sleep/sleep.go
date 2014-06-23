package main

import (
    "os"
    "strconv"
    "time"
)

func main() {
    if len(os.Args) != 2 {
        os.Exit(1)
    }

    duration, err := strconv.ParseInt(os.Args[1], 0, 64)

    if err != nil {
        os.Exit(1)
    }

    time.Sleep(time.Duration(duration) * 1000 * 1000 * 1000)
}