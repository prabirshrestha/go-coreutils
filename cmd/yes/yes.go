package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for {
			fmt.Println(os.Args[1])
		}
	} else {
		for {
			fmt.Println("y")
		}
	}
}
