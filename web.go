package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

// Store file in loco / maybe NFS if needed ?
func (p *Page) save() error {
	filename := p.Title + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// load HTML pages provided by app
func loadPage(title string) *Page {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return &Page{Title: title, Body: body}
}

// Start handler to server HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Starting Server:  ", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/maps/"):]
	p := loadPage(title)
	t, _ := template.ParseFiles("maps.html")
	t.Execute(w, p)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	page1 := &Page{Title: "TEST", Body: []byte("This is a Test")}
	page1.save()

	page2 := loadPage("TEST")
	fmt.Println(w, "<hs>%s</h1><div>%s</div>", page2.Title, page2.Body)
}

func main() {
	http.HandleFunc("/maps/", viewHandler)
	http.HandleFunc("/test/", testHandler)
	http.ListenAndServe(":8080", nil)

}
