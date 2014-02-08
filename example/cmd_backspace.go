package main

import "fmt"

import . "github.com/enderlu/vfp"
import "os"

func main() {

	go RCount()
	Wait()

}

func RCount() {
	fmt.Print("please input:")
	znotfirst := false
	zx := ""
	for zi := 0; zi < 10000000; zi++ {
		zold := zx
		zx = fmt.Sprintf("%v/%v", zi, 10000000)

		if znotfirst {
			fmt.Print(Replicate("\b", len(zold)) + zx)
		} else {
			fmt.Print(zx)
			znotfirst = true
		}
	}
}
