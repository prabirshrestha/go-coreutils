package main

import (
    "fmt"
    "syscall"
)

var (
    kernel32Dll = syscall.NewLazyDLL("kernel32.dll")

    procGetTickCount64 = kernel32Dll.NewProc("GetTickCount64")
)

func main() {
    fmt.Println(getTickCount())
}

func getTickCount() (tick uint64) {
    ret, _, _ := procGetTickCount64.Call()
    return uint64(ret)
}