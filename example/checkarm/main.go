//package main

//import (
//	"fmt"
//	"net/http"
//)

//func main() {
//	zurl := `http://tieba.baidu.com/f?kw=%C3%A7%BB%C4%BC%CD`
//	fmt.Println("\nchecking ...", zurl)
//	r, err1 := http.Get(zurl)

//	if err1 != nil {
//		fmt.Println(" arm error:", err1)
//		return
//	}
//	fmt.Println("\n内容:", ReadBody(r))
//	return

//}
//func ReadBody(r *http.Response) string {
//	bs := make([]byte, 512)
//	defer r.Body.Close()
//	var buf []byte

//	n, _ := r.Body.Read(bs)
//	zi := 0
//	for n > 0 {
//		if n > 0 {
//			zi++
//			buf = append(buf, bs[:n]...)
//		}
//		n, _ = r.Body.Read(bs)
//	}
//	return string(buf)
//}

package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	//"github.com/axgle/service"
	//"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"time"

	"github.com/enderlu/vfp"
)
import "math/rand"

var zurl, zw, zsound string
var zfindMap map[string]string = make(map[string]string)
var zpath string = vfp.Addbs(vfp.Justpath(vfp.Program()))

func checkError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
func readStr(zsound_b []byte) (zsound string) {
	zsound = ""
	if zsound_b[0] == 239 && zsound_b[1] == 187 && zsound_b[2] == 191 {
		zsound = string(zsound_b[3:])
	} else {
		zsound = string(zsound_b)
	}
	return
}
func main() {

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
	Wait()

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

		if err1 != nil {
			fmt.Println(" arm error:", err1)
		}
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
		zfound_word := ""
		znot_found := ""
		//zh, zn := 0, 0
		zset := false
		znew_found := ""
		znew_set := false

		for _, zv = range zwords {
			//zh++
			if zv == "" {
				continue
			}
			if At(zv, string(buf_line)) > 0 {
				_, zok := zfindMap[zv]
				if !zok {
					if znew_set {
						znew_found = znew_found + "," + zv
					} else {
						znew_found = znew_found + zv
					}
					znew_set = true
					zfindMap[zv] = zv
				} else {
					if zfound {
						zfound_word = zfound_word + "," + zv
					} else {
						zfound_word = zfound_word + zv
					}

				}
				zfound = true

				//zn++
			} else {
				if zset {
					znot_found = znot_found + "," + zv
				} else {
					znot_found = znot_found + zv
				}
				zset = true
			}
		}

		//zmind := ""

		//if znot_found != "" {
		//	zmind = "未找到:" + znot_found + "\n未找全，隔10分钟继续下一轮查找\n"
		//} else {
		//	zmind = "已经找全了，隔10分钟继续下一轮查找\n"
		//}

		if znew_set && zfound {
			fmt.Println("found! ", zv)

			zuin_b, _ := vfp.Filetostr(zpath + "uin.txt")
			zun_b, _ := vfp.Filetostr(zpath + "un.txt")
			zpass_b, _ := vfp.Filetostr(zpath + "pass.txt")

			SendQQMsg(readStr(zuin_b), readStr(zpass_b), readStr(zun_b),
				"新找到: "+znew_found+" \n"+
					"已找到: "+zfound_word+" \n"+
					"未找到: "+znot_found+" \n")

			//SendQQMsg(readStr(zuin_b), readStr(zpass_b), readStr(zun_b),
			//	"x6.GoLang~\n"+
			//		fmt.Sprintf("%v\n", vfp.Datetime())+
			//		"找到: "+zfound_word+" \n"+
			//		zmind+
			//		`url: `+html.EscapeString(zurl))

			//MessageBox("查找结果", "找到了\n"+zw, 0)
			//if zn == zh {
			//	break
			//}

		} else {
			fmt.Println("can not find ", zw)
		}
		time.Sleep(10 * 60 * time.Second)
	}

}
func At(zsubstring, zwholestring string) int {
	return strings.Index(zwholestring, zsubstring) + 1
}

func SendQQMsg(zuin, zpass_md5, zun, zmsg string) error {
	var zseed int64 = time.Now().Unix()

	host := "http://121.14.102.159:14000"
	posttype := `text/plain;charset=UTF-8`
	rand.Seed(zseed)
	xrand := fmt.Sprintf("%v", rand.Intn(200))

	poststring := `VER=1.4&CON=1&CMD=Login&SEQ=` + xrand + `&UIN=` + zuin +
		`&PS=` + zpass_md5 +
		`&M5=1&LG=0&LC=812822641C978097&GD=EX4RLS2GFYGR6T1R&CKE=`

	r, err := http.Post(host, posttype, strings.NewReader(poststring))
	if err != nil {
		return err
	}
	rand.Seed(zseed)
	xrand = fmt.Sprintf("%v", rand.Intn(200))

	poststring = `VER=1.4&CMD=CLTMSG&SEQ=` + xrand + `&UIN=` + zuin +
		`&UN=` + zun + `&MG=` + zmsg
	r, err = http.Post(host, posttype, strings.NewReader(poststring))
	if err != nil {
		return err
	}
	//rand.Seed(zseed)
	//xrand = fmt.Sprintf("%v", rand.Intn(200))

	//poststring = `VER=1.4&CON=1&CMD=Logout&SEQ=` + xrand + `&UIN=` + zuin + `&SID=&XP=C4CA4238A0B92382`
	//r, err = http.Post(host, posttype, strings.NewReader(poststring))
	//if err != nil {
	//	return err
	//}

	if 1 == 0 {
		fmt.Println(ReadBody(r))
	}

	return nil
}

func ReadBody(r *http.Response) string {
	bs := make([]byte, 5056)
	defer r.Body.Close()
	var buf []byte

	n, _ := r.Body.Read(bs)
	zi := 0
	for n > 0 {
		if n > 0 {
			zi++
			buf = append(buf, bs[:n]...)
		}
		n, _ = r.Body.Read(bs)
	}
	return string(buf)
}
