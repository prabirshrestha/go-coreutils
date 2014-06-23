package main

import (
    "fmt"
    "os"
    "os/user"
)

func main() {
    usr, err := user.Current()

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(usr.Username)
}