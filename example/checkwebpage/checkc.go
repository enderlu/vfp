package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	//"github.com/axgle/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
	"vfp"
)

var zurl, zw, zsound string

var (
	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

func checkError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
func Main_Sit() {
	zpath := vfp.Addbs(vfp.Justpath(vfp.Program()))
	zsound_b, _ := ioutil.ReadFile(zpath + `sound.txt`)
	if zsound_b[0] == 239 && zsound_b[1] == 187 && zsound_b[2] == 191 {
		zsound = string(zsound_b[3:])
	} else {
		zsound = string(zsound_b)
	}

	zurl_b, _ := ioutil.ReadFile(zpath + "url.txt")
	zurl = string(zurl_b)
	fmt.Print("\nsearch url:", zurl)

	zw_b, _ := ioutil.ReadFile(zpath + "search.txt")
	//如果是notepad录入的，则有个前缀
	if zw_b[0] == 239 && zw_b[1] == 187 && zw_b[2] == 191 {
		zw = string(zw_b[3:])
	} else {
		zw = string(zw_b)
	}

	fmt.Print("\nsearch word:", zw)

	go DisplayUrl(zurl)
	////Wait()

}

func Wait() string {
	reader := bufio.NewReader(os.Stdin)
	zr, _, _ := reader.ReadLine()
	zline := string(zr)
	return zline
}

func DisplayUrl(zurl string) {
	for {
		fmt.Println("\nchecking ...", zurl)
		r, err1 := http.Get(zurl)

		zcode := vfp.Strextract(r.Header["Content-Type"][0], "charset=", "charset=", 1)
		zcode = vfp.Lower(zcode)
		switch zcode {
		case "gbk", "gb2312", "gb18030":
			zcode = "gb18030"
		default:
			zcode = "utf8"
		}
		checkError(err1)
		fmt.Println("search status :", func() string {
			if r.StatusCode == 200 {
				return "OK"
			}
			return string(r.StatusCode)
		}())

		bs := make([]byte, 5056)

		var buf []byte
		defer r.Body.Close()

		n, err1 := r.Body.Read(bs)
		zi := 0
		for n > 0 {
			if n > 0 {
				zi++
				buf = append(buf, bs[:n]...)
			}
			n, err1 = r.Body.Read(bs)
		}
		decoder := mahonia.NewDecoder(zcode)
		_, buf_line, _ := decoder.Translate(buf, true)

		zfound := false
		zwords := vfp.Aline(zw)
		zv := ""
		for _, zv = range zwords {
			if At(zv, string(buf_line)) > 0 {
				zfound = true
				break
			}
		}

		if zfound {
			fmt.Println("found! ", zv)
			go vfp.PlayX(zsound)
			//MessageBox("查找结果", "找到了\n"+zw, 0)
			break

		} else {
			fmt.Println("can not find ", zw)
		}
		time.Sleep(10 * time.Second)
	}

}
func At(zsubstring, zwholestring string) int {
	return strings.Index(zwholestring, zsubstring) + 1
}

func MessageBox(caption, text string, style uintptr) (result int) {
	ret, _, _ := syscall.Syscall6(uintptr(messageBox),
		6, 0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		style, 0, 0)
	return int(ret)
}
