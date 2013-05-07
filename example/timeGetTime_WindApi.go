package main

import (
	"fmt"
	"syscall"
)

func main() {
	fmt.Println(TimeGetTime())

}

//函数以毫秒计的系统时间。该时间为从系统开启算起所经过的时间
func TimeGetTime() int64 {
	zlib := MustLoadLibrary("winmm.dll")

	zproc := MustGetProcAddress(zlib, "timeGetTime")

	zret, _, _ := syscall.Syscall(zproc, 0, 0, 0, 0)
	return int64(zret)
}
func MustLoadLibrary(name string) uintptr {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}

	return uintptr(lib)
}

func MustGetProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}

	return uintptr(addr)
}
