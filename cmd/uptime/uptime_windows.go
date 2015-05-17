// +build windows
package main

import (
	"syscall"
)

var (
	kernel32Dll = syscall.NewLazyDLL("kernel32.dll")

	procGetTickCount64 = kernel32Dll.NewProc("GetTickCount64")
)

func getTickCount() (tick uint64, err error) {
	ret, _, _ := procGetTickCount64.Call()
	return uint64(ret), nil
}
