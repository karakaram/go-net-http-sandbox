package main

import (
	"encoding/json"
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
	router := http.NewServeMux()
	router.HandleFunc("/health/", CORSMiddleware(healthHandler))
	router.HandleFunc("/pages/", CORSMiddleware(pageHandler))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func CORSMiddleware(hander http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			return
		}
		hander(w, r)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		PageGetHandler(w, r)
	case "POST":
		PagePostHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func PageGetHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/pages/"):]
	page, _ := loadPage(title)

	json.NewEncoder(w).Encode(page)
}

func PagePostHandler(w http.ResponseWriter, r *http.Request) {
	raw, _ := ioutil.ReadAll(r.Body)
	var page Page
	json.Unmarshal(raw, &page)
	page.save()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, string(raw))
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
