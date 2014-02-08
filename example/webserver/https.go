package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	_ "path/filepath"
	"strconv"
)

var dir string
var port int
var staticHandler http.Handler

// 初始化参数
func Init() {
	dir = path.Dir(os.Args[0])
	flag.IntVar(&port, "port", 12345, "服务器端口")
	flag.Parse()
	staticHandler = http.FileServer(http.Dir(dir))
}

func main() {
	println("Server Started!")
	Init()
	http.HandleFunc("/", StaticServer)
	http.HandleFunc("/data1", data1)
	http.HandleFunc("/savedata", savedata)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func data1(w http.ResponseWriter, req *http.Request) {
	zstr, zerr := ioutil.ReadFile("miniui/tasks.txt")
	if zerr != nil {
		fmt.Fprintf(w, "%v", zerr)
	} else {
		fmt.Fprintf(w, "%v", string(zstr))
	}
	//fmt.Fprintf(w, "%v", ztasks)

}

func savedata(w http.ResponseWriter, req *http.Request) {
	data := req.FormValue("data")
	log.Println("保存数据：")
	log.Println(data)

}

// 静态文件处理
func StaticServer(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)
	if req.URL.Path != "/" {
		staticHandler.ServeHTTP(w, req)
		return
	}
	req.URL.Path = "/default.html"
	staticHandler.ServeHTTP(w, req)
}
