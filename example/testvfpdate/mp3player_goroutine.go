package main

import (
	"fmt"
	. "github.com/enderlu/vfp"
	"os"
	"sync"
	"time"
)

var zsong string = ""

var zfiles []os.FileInfo
var zdir string
var zx string = "1"
var changed bool = false
var lock *sync.Mutex = new(sync.Mutex)

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

	for {

		switch zx {
		case "l":
			listSound()
		case "q":
			fmt.Println("\nstatus:", MCIStatus(Md5(zsong)))
		case "s":
			MCISendString("stop " + Md5(zsong))
		case "r":
			MCISendString("play " + Md5(zsong))
		case "x":
			return
		default:
			go playList()
		}

		zx = Wait()
		lock.Lock()
		changed = true
		lock.Unlock()
		time.Sleep(120 * time.Millisecond)
	}

}
func playList() {
	for {

		zid := int(Val(zx)) - 1
		if zid >= 0 && zid < len(zfiles) {
			MCISendString("close " + Md5(zsong))
			zsong = zdir + zfiles[zid].Name()
			fmt.Println("\nPlaying Song：", zx+"."+zsong)
			fmt.Print("press song number or r = resume ,s = stop ,l = list song ,x = exit :")
			PlayX(zsong)
		} else {
			fmt.Printf("\nSong Range：[%v ~ %v]\n", 1, len(zfiles))
		}
		for MCIStatus(Md5(zsong)) != "stopped" && !changed {
			time.Sleep(100 * time.Millisecond)
		}
		if changed {
			lock.Lock()
			changed = false
			lock.Unlock()
			break
		}

		zid++
		if zid >= 0 && zid < len(zfiles) {
			zx = Transform(zid + 1)
		} else {
			break
		}
	}
	//fmt.Println("Close song", zsong)

}

func listSound() {
	zfiles, _ = Adir(zdir + `*.mp3`)
	zlen := len(zfiles)
	for zi := 0; zi < zlen; zi++ {
		for zk := zi; zk <= zi+3; zk++ {
			if zk >= zlen {
				break
			}
			zv := zfiles[zk]
			fmt.Print(Padl(Transform(zk+1), 3, " "), ".", Padr(zv.Name(), 60, " "), `  `)

		}
		fmt.Println("")
		zi += 3
	}
}
