package main

import (
	"bufio"
	"github.com/axgle/mahonia"

	"log"
	"net/http"
	"os"
)

func checkError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	//DisplayFile(`test.txt`)

	//DisplayUrl(`http://www.qdwenxue.com/BookReader/2403629,40213317.aspx`)
	//DisplayUrl(`http://www.baidu.com`)
	DisplayUrl(`http://xbrl.cninfo.com.cn/XBRL/allinfo.jsp?stkid=000005&getyear=2008`)
	
}

func DisplayFile(zfile string) {
	f, err := os.Open(zfile)
	checkError(err)
	defer f.Close()
	decoder := mahonia.NewDecoder("gb18030")
	//decoder := mahonia.NewDecoder("utf8")
	//decoder := mahonia.NewDecoder("big5")
	r := bufio.NewReader(decoder.NewReader(f))
	line, _, err := r.ReadLine()
	for err == nil {
		checkError(err)
		println(string(line))
		line, _, err = r.ReadLine()
	}
}

func DisplayUrl(zurl string) {
	println("loading ..." ,zurl)
	r, err1 := http.Get(zurl)

	checkError(err1)
	println("complete !" ,r.StatusCode)
	decoder := mahonia.NewDecoder("gb18030")//gb18030可以适用于gb2312
	//decoder := mahonia.NewDecoder("utf8")
	//decoder := mahonia.NewDecoder("big5")
	bs := make([]byte, 5056)
	
	var buf []byte
	defer r.Body.Close()
	
	n, err1 := r.Body.Read(bs)
	zi  := 0
	for ;n>0 ; {
		if n > 0 {
			zi++
			buf = append(buf, bs[:n]...)
		}
		n, err1 = r.Body.Read(bs)
	}
	println("read complete !" ,"times:" ,zi )
	
	_, line, _ := decoder.Translate(buf, true)
	println(n, string(line))
}
