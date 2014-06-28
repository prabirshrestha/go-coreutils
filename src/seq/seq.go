package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) != 2 {
        os.Exit(1)
    }

    start := int64(1)
    end, err := strconv.ParseInt(os.Args[1], 0, 64)

    if err != nil {
        os.Exit(1)
    }

    for i := start; i < end; i++ {
        fmt.Println(i)
    }
}