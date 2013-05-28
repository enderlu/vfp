package main

import (
	"fmt"
	"unsafe"
)

const N int = int(unsafe.Sizeof(0))

func main() {

	x := 0x1234

	p := unsafe.Pointer(&x) // *int -> Pointer

	p2 := (*[N]byte)(p) // Pointer -> *[4]int，注意 slice 的内存布局和 array 是不同的。

	// 数组类型元素⻓度必须是常量。

	for i, m := 0, len(p2); i < m; i++ {

		fmt.Printf("%02X ", p2[i])

	}
	p2[0] = 0x56

	fmt.Printf("%X ", x)

}

//result: 34 12 00 00 1256
