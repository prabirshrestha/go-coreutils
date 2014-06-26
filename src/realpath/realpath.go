package main

import (
    "fmt"
    "path/filepath"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        os.Exit(1)
    }

    path, err := filepath.Abs(os.Args[1])

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(path)
}