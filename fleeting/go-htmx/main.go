package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Go With HTMX")

	handler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := getMockedDataForTesting()
		tmpl.Execute(w, films)
	}

	apiHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMX request")
		log.Print(r.Header.Get("HX-Request"))
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-light'>%s - %s</li>", title, director)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/add-film/", apiHandler)
	log.Fatal(http.ListenAndServe(":8005", nil))
}

func getMockedDataForTesting() map[string][]Film {
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenters"},
		},
	}
	return films
}
