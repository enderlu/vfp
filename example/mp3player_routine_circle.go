package main

import (
	"fmt"
	"os"
	"time"
	. "vfp"
)

var zsong string = ""
var zfiles []os.FileInfo
var zdir string
var changed bool
var zx string = "1"

func main() {
	fmt.Println("X6.M Playing sound...")

	if len(os.Args) == 2 {
		zdir = os.Args[1]
		fmt.Printf("Read %v's file...\n", zdir)
		Strtofile(zdir, "songdir.txt")
	} else {
		if File("songdir.txt") {
			zdata, _ := Filetostr("songdir.txt")
			zdir = string(zdata)
		} else {
			zdir = Justpath(Program())
		}

	}

	zdir = Addbs(zdir)
	listSound()
	if len(zfiles) == 0 {
		fmt.Println("把播放器放在音乐目录中运行即可,回车退出")
		Wait()
		return
	}

	go commandList()

	for {
		zx = Wait()
		changed = true
		if zx == "x" {
			break
		}
	}
}

func commandList() {
	zid := 0
	for {

		switch zx {
		case "l":
			listSound()
		case "q":
			fmt.Println("\nstatus:", MCIStatus(Md5(zsong)), zsong)
		case "s":
			fmt.Println("stopping song.. ")
			MCISendString("stop " + Md5(zsong))

		case "r":
			fmt.Println("resume song.. ")
			MCISendString("play " + Md5(zsong))
		case "x":
			break
		case "p":
			fmt.Println("\nreplay song ", zsong)
			PlayX(zsong)
		default:
			zid = int(Val(zx)) - 1
			if zid >= 0 && zid < len(zfiles) {
				MCISendString("close " + Md5(zsong))
				zsong = zdir + zfiles[zid].Name()
				fmt.Println("\nPlaying Song：", zx+"."+zsong)
				PlayX(zsong)
			} else {
				fmt.Printf("\nSong Range：[%v ~ %v]\n", 1, len(zfiles))
			}
		}

		fmt.Print("press song number or r = resume ,s = stop ,l = list song ,x = exit :")

		for (zx == "s" && !changed) || (MCIStatus(Md5(zsong)) != "stopped" && !changed) {
			time.Sleep(100 * time.Millisecond)

		}

		if zx != "s" && !changed {
			zid++

			if !(zid >= 0 && zid < len(zfiles)) {
				zid = 0
			}
			zx = Transform(zid + 1)
		}
		changed = false
	}

}

func listSound() {
	zfiles, _ = Adir(zdir + `*.mp3`)
	zlen := len(zfiles)
	zmaxpad := 0
	for zi := 0; zi < zlen; zi++ {
		zmaxpad = int(Max(float64(zmaxpad), float64(len(zfiles[zi].Name()))))
	}
	zmaxpad += 5
	for zi := 0; zi < zlen; zi++ {
		for zk := zi; zk <= zi+3; zk++ {
			if zk >= zlen {
				break
			}
			zv := zfiles[zk]
			fmt.Print(Padl(Transform(zk+1), 3, " "), ".", Padr(zv.Name(), zmaxpad, " "), `  `)

		}
		fmt.Println("")
		zi += 3
	}
}
