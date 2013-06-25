	package main

	import (
		"io"
		"net/http"
		"log"
	)

	// hello world, the web server
	func HelloServer(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello, world!你好\n")
	}

	func main() {
		http.HandleFunc("/", HelloServer)
		err := http.ListenAndServe(":88", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}