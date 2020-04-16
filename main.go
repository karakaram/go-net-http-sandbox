package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

const pageDir = "pages/"

func main() {
	//p1 := &Page{Title: "hello", Body: "Hello World"}
	//p1.save()
	//
	//p2, _ := loadPage("hello")
	//fmt.Println(p2.Body)
	http.HandleFunc("/health/", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func (p *Page) save() error {
	filename := pageDir + p.Title + ".txt"
	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func loadPage(title string) (*Page, error) {
	filename := pageDir + title + ".txt"
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(raw)}, nil
}
