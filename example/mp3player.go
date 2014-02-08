package main

import (
	"fmt"
	. "github.com/enderlu/vfp"
	"os"
)

var zsong string = ""

func main() {

	fmt.Println("playing sound...")

	zdir := `C:\Kugou\Listen\`
	if len(os.Args) == 2 {
		zdir = os.Args[1]
		fmt.Printf("Read %v's file...\n", zdir)
	}
	zdir = Addbs(zdir)
	zfiles, _ := Adir(zdir + `*.mp3`)
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
	if len(zfiles) == 0 {
		return
	}
	zx := "1"
	znotfirst := false
	for {

		if znotfirst {
			fmt.Print("press song number or r = resume ,s = stop ,x = exit :")
			zx = Wait()
		}
		znotfirst = true
		switch zx {

		case "s":
			MCISendString("stop " + Md5(zsong))
		case "r":
			MCISendString("play " + Md5(zsong))
		case "x":

			return
		default:
			zid := int(Val(zx)) - 1
			if zid >= 0 && zid < len(zfiles) {
				MCISendString("close " + Md5(zsong))
				zsong = zdir + zfiles[zid].Name()
				fmt.Println("正在播放：", zsong)
				PlayX(zsong)
			} else {
				fmt.Printf("合法范围：[%v ~ %v]\n", 1, len(zfiles))
				continue
			}

		}

	}

}
