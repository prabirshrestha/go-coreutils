package main

import (
    "fmt"
    "os"
)

func main() {
    env := os.Environ()

    for _, envar := range env {
        fmt.Println(envar)
    }
}
