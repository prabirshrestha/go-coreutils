// +build windows
package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"syscall"
	"unsafe"
)

var (
	rpcrt4Dll = syscall.NewLazyDLL("rpcrt4.dll")

	procUuidCreate = rpcrt4Dll.NewProc("UuidCreate")
)

const (
	RPC_S_OK = 0
)

func uuidgen() ([]byte, error) {
	var uuid [16]byte
	rc, _, e := syscall.Syscall(procUuidCreate.Addr(), 1,
		uintptr(unsafe.Pointer(&uuid[0])), 0, 0)
	if int(rc) != RPC_S_OK {
		if e != 0 {
			return nil, error(e)
		} else {
			return nil, syscall.EINVAL
		}
	}
	return uuid[:], nil
}

func generateUUID(w io.Writer) error {
	data, err := uuidgen()
	if err != nil {
		return err
	}

	uuid := hex.EncodeToString(data)

	fmt.Fprintf(w, "%s-%s-%s-%s-%s", uuid[:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])
	fmt.Fprintln(w)

	return nil
}

func generateUUIDs(w io.Writer, count int) error {
	for i := 0; i < count; i++ {
		err := generateUUID(w)
		if err != nil {
			return err
		}
	}
	return nil
}
