package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
)

//产生10个随机数
func main() {
	ExampleRead()

}

// This example reads 10 cryptographically secure pseudorandom numbers from
// rand.Reader and writes them to a byte slice.
func ExampleRead() {
	c := 10
	b := make([]byte, c)
	n, err := io.ReadFull(rand.Reader, b)
	if n != len(b) || err != nil {
		fmt.Println("error:", err)
		return
	}
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, c)))
	fmt.Println("rand:", b, len(b))

	// Output:
	// false
}
