package main

import (
	"crypto/md5"
	"fmt"
	//"html"
	"net/http"
	"strings"
	"time"
)

import "math/rand"

/*
	可用的服务器IP及端口
　　218.18.95.203:21001
　　119.147.10.10:14000
　　119.147.10.11:14000
　　119.147.14.253:14000
　　119.147.7.16:14000
　　121.14.102.159:14000


	登陆：
　　VER=1.4&CON=1&CMD=Login&SEQ=131&UIN=1402607192&PS=7877E0E69A15427080C5CF2BCB1E64CE&M5=1&LG=0&LC=812822641C978097&GD=EX4RLS2GFYGR6T1R&CKE=

	验证码：
　　VER=1.4&CON=1&CMD=VERIFYCODE&SEQ=132&UIN=1402607192&SID=null&XP=null&SC=2&VC=DBVP
	在线：
	　　VER=1.4&CON=1&CMD=Change_Stat&SEQ=141&UIN=1402607192&SID=&XP=C4CA4238A0B92382&ST=10
	隐身
	　　VER=1.4&CON=1&CMD=Change_Stat&SEQ=140&UIN=1402607192&SID=&XP=C4CA4238A0B92382&ST=40
	注销：
	　　VER=1.4&CON=1&CMD=Logout&SEQ=143&UIN=1402607192&SID=&XP=C4CA4238A0B92382

	发送消息：
	VER=1.4&CMD=CLTMSG&SEQ=&UIN=&UN=&MG=

	文档链接：
		http://tieba.baidu.com/p/2169489666?pn=1
		http://www.cnblogs.com/czjone/archive/2010/03/16/1636017.html
		http://wenku.baidu.com/view/dc2edd360b4c2e3f57276334.html

*/

func main() {
	zurl := `http://tieba.baidu.com/f?ie=utf-8&kw=%E6%B1%82%E9%AD%94`
	err := SendQQMsg("4069685031", "tx2000", "1382319470", "日你妹的阿阿凡啊\n"+zurl)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("finished!")
}
func SendQQMsg(zuin, zpass, zun, zmsg string) error {
	var zseed int64 = time.Now().Unix()

	host := "http://121.14.102.159:14000"
	posttype := `text/plain;charset=UTF-8`
	rand.Seed(zseed)
	xrand := fmt.Sprintf("%v", rand.Intn(200))

	poststring := `VER=1.4&CON=1&CMD=Login&SEQ=` + xrand + `&UIN=` + zuin +
		`&PS=` + strings.ToUpper(Md5(zpass)) +
		`&M5=1&LG=0&LC=812822641C978097&GD=EX4RLS2GFYGR6T1R&CKE=`

	r, err := http.Post(host, posttype, strings.NewReader(poststring))
	if err != nil {
		return err
	}
	rand.Seed(zseed)
	xrand = fmt.Sprintf("%v", rand.Intn(200))

	poststring = `VER=1.4&CMD=CLTMSG&SEQ=` + xrand + `&UIN=` + zuin +
		`&UN=` + zun + `&MG=` + zmsg + xrand
	r, err = http.Post(host, posttype, strings.NewReader(poststring))
	if err != nil {
		return err
	}
	rand.Seed(zseed)
	xrand = fmt.Sprintf("%v", rand.Intn(200))

	poststring = `VER=1.4&CON=1&CMD=Logout&SEQ=` + xrand + `&UIN=` + zuin + `&SID=&XP=C4CA4238A0B92382`
	r, err = http.Post(host, posttype, strings.NewReader(poststring))
	if err != nil {
		return err
	}

	if 1 == 0 {
		fmt.Println(ReadBody(r))
	}

	return nil
}
func Md5(zstrArg interface{}) string {
	var zstr []byte
	switch zstrArg.(type) {
	case string:
		zstr = []byte(zstrArg.(string))
	case []byte:
		zstr = zstrArg.([]byte)
	default:
		zstr = []byte(fmt.Sprintf("%v", zstrArg))
	}
	zs := md5.New()
	zs.Write([]byte(zstr))
	return fmt.Sprintf("%x", zs.Sum(nil))
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
