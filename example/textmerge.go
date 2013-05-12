package main

import (
	"fmt"
	"log"
	"text/template"
)

func main() {

	const zsql = `
		insert into dbo.product(pid ,pname ,price ,categid) 
		values('{{.Id}}' ,'{{.Name}}' ,{{.Price}} ,'{{.Cat}}')
						`
	t := template.Must(template.New("zsql").Parse(zsql))
	ztext := &TextM{}
	err := t.Execute(ztext, struct {
		Id    string
		Name  string
		Price float64
		Cat   string
	}{Id: "01", Name: "P1", Price: 66.4, Cat: "C01"})

	if err != nil {
		log.Println("executing template:", err)
	} else {
		fmt.Println(ztext)
	}

}

type p struct {
	Id    string
	Name  string
	Price float64
	Cat   string
}

type TextM struct {
	mtext []byte
}

func (t *TextM) Write(p []byte) (n int, err error) {
	t.mtext = append(t.mtext, p...)
	return len(t.mtext), nil
}

func (t *TextM) String() string {
	return string(t.mtext)
}
