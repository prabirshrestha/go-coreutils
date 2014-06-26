package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    file, err := ioutil.TempFile(os.TempDir(), "");

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(file.Name())

}