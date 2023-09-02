package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title string
	Director string
}

func main() {
    fmt.Println("hello world")

	h1:= func (w http.ResponseWriter, r *http.Request)  {
		// io.WriteString(w, "Hello World\n")
		// io.WriteString(w, r.Method)
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Oppenheimer", Director: "Christoper Nolan"},
				{Title: "Barbie", Director: "Christoper Nolun"},
				{Title: "Transformer", Director: "Michael Bay"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func (w http.ResponseWriter, r *http.Request)  {
		// Dynamic Render 
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})

		// Hard Code Test

		// Check for the value sent from input
		// fmt.Print(title)
		// fmt.Print(director)

		// Template for adding film lists
		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary-subtle'>%s - %s</li>", title, director)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		// tmpl.Execute(w, nil)
		
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}