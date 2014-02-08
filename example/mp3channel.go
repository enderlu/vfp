package main

import (
	"fmt"
	. "github.com/enderlu/vfp"
	"os"
)

var zsong string = ""

func main() {
	fmt.Println("X6.M Playing sound...")
	zsong = `D:\Kugou\杨幂 - 爱的供养.mp3`

	ch := make(chan string)

	go playList(ch)

	ch <- "p"
	for {
		fmt.Print("press song number or r = resume ,s = stop ,q = song state,x = exit :")
		zx := Wait()
		<-ch
		ch <- zx
		//fmt.Println("you input :", zx)
	}
}

func playList(ch chan string) {
	for {
		zx := <-ch
		switch zx {

		case "q":
			fmt.Println("\nstatus:", MCIStatus(Md5(zsong)))
		case "s":
			fmt.Println("stopping song.. ")
			MCISendString("stop " + Md5(zsong))
		case "r":
			fmt.Println("resume song.. ")
			MCISendString("play " + Md5(zsong))
		case "x":
			os.Exit(0)
		case "p":
			fmt.Println("play song ", zsong)
			PlayX(zsong)
		}

		ch <- "wait"

		//fmt.Println("clear wait.. ")
	}

}
