package main

import "github.com/enderlu/vfp"
import (
	"syscall"
	"unsafe"
)

var (
	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

func PlaySound(zsound string) error {
	vfp.PlayX(zsound)
	return nil
}
func MessageBox(caption, text string, style uintptr) (result int) {
	ret, _, _ := syscall.Syscall6(uintptr(messageBox),
		6, 0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		style, 0, 0)
	return int(ret)
}
