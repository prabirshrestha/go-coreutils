package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    pwd, err := os.Getwd()

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    files, err := ioutil.ReadDir(pwd)
    for _, f := range files {
        fmt.Println(f.Name())
    }
}
