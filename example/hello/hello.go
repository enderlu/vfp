package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path == "favicon.ico" {
		http.NotFound(w, r)
		return
	}
	if path == "" {
		path = "index.html"
	}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(w, "404")
		return
	}
	fmt.Fprintf(w, "%s\n", contents)
}

func Add(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	if name == "" || age == "" {
		fmt.Fprint(w, AddForm)
		return
	}
	fmt.Fprintf(w, "Save : Your name is  %s , You age is %s", name, age)
}

func MyTime(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%v", time.Now())
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/time", MyTime)
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

const AddForm = `
<html><body>
<form method="POST" action="/add">
Name: <input type="text" name="name">
Age: <input type="text" name="age">
<input type="submit" value="Add">
</form>
</body></html>
`
