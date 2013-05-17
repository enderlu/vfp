package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	. "vfp"
)

func main() {
	ExampleWriter()
	ExampleReader()
}
func ExampleWriter() {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	Strtofile(buf.Bytes(), `c:\xx.zip`)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleReader() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(`c:\xx.zip`)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Println("\nopen err:", err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			fmt.Println("\ncopy err:", err)
		}
		rc.Close()
		fmt.Println()
	}
	// Output:
	// Contents of README:
	// This is the source code repository for the Go programming language.
}