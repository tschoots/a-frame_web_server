package main 

import (
	"net/http"
	"path/filepath"
	"html/template"

)

func main() {
	//http.Handle("/", http.FileServer(http.Dir("./a-frame")))
	http.HandleFunc("/",  func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles(filepath.Join("templates", "main.html")))
		msg := "test"
		templ.Execute(w, msg)
	})
	
	http.ListenAndServe(":8080", nil)

}

