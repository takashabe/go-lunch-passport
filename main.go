package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var indexTemplate *template.Template

type Page struct {
	Title    string
	Contents string
}

func gpsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lat := params["lat"]
	//params := r.URL.Query()
	//lat := params.Get(":lat")
	//lon := params.Get(":lon")
	//store_id := params.Get(":store_id")

	//fmt.Fprintf(w, "lat=%s, lon=%s, store_id=%s", lat, lon, store_id)
	fmt.Fprintf(w, "lat = %s", lat)
}

func init() {
	indexTemplate = template.Must(template.ParseFiles("views/index.html"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{lat}", gpsHandler)
	// r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
