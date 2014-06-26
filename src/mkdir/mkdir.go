package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        os.Exit(1)
    }

    err := os.Mkdir(os.Args[1], 0777)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


