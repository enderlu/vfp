package vfp

import (
	"syscall"
	"unsafe"
)

type (
	BOOL    int32
	HRESULT int32
	HMODULE int32
	DWORD   int32
	LPCTSTR string
)

var zlibwinmm uintptr = MustLoadLibrary(`winmm.dll`)
var playSound uintptr = MustGetProcAddress(zlibwinmm, "PlaySound")
var mciSendString uintptr = MustGetProcAddress(zlibwinmm, "mciSendStringW")

const (
	SND_SYNC      = 0x0000     /* play synchronously (default) */
	SND_ASYNC     = 0x0001     /* play asynchronously */
	SND_NODEFAULT = 0x0002     /* silence (!default) if sound not found */
	SND_MEMORY    = 0x0004     /* pszSound points to a memory file */
	SND_LOOP      = 0x0008     /* loop the sound until next sndPlaySound */
	SND_NOSTOP    = 0x0010     /* don't stop any currently playing sound */
	SND_NOWAIT    = 0x00002000 /* don't wait if the driver is busy */
	SND_ALIAS     = 0x00010000 /* name is a registry alias */
	SND_ALIAS_ID  = 0x00110000 /* alias is a predefined ID */
	SND_FILENAME  = 0x00020000 /* name is file name */
	SND_RESOURCE  = 0x00040004 /* name is resource name or atom */

	SND_PURGE       = 0x0040 /* purge non-static events for task */
	SND_APPLICATION = 0x0080 /* look for application specific association */
)

/*
package main
import (
	"fmt"
	. "vfp"
)
func main() {
	fmt.Println("playing sound...")
	MCISendString(`C:\Kugou\Listen\tn.mp3`)
	fmt.Println("end play sound...",Wait())
}
*/
func MCISendString(zcmd string) int {

	zret, _, _ := syscall.Syscall6(mciSendString, 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(zcmd))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(""))),
		uintptr(0), uintptr(0), 0, 0)

	return int(zret)
}

func MCIStatus(zsong string) string {
	zs := new(Mst)
	zdd := 20
	zcmd := "status " + zsong + " mode"
	syscall.Syscall6(mciSendString, 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(zcmd))),
		uintptr(unsafe.Pointer(zs)),
		uintptr(unsafe.Pointer(&zdd)), uintptr(0), 0, 0)

	return Strtran(string(zs.Data[0:20]), Chr(0), "", -1)
}

func PlayX(zfile string) {
	MCISendString("close " + Md5(zfile))
	MCISendString(`open "` + zfile + `" alias ` + Md5(zfile))
	MCISendString("play " + Md5(zfile))
}

/*
PlaySound(gsoundFile, 0, SND_FILENAME|SND_ASYNC)
PlaySound("", 0, 0)
*/
func PlaySound(pszSound string, hmod HMODULE, fdwSound DWORD) BOOL {
	zret, _, _ := syscall.Syscall(playSound, 3,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(pszSound))),
		uintptr(hmod), uintptr(fdwSound))
	return BOOL(zret)

}

type Mst struct {
	Data [20]byte
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
