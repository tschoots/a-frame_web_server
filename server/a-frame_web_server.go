package main 

import (
	"net/http"
	//"path/filepath"
	"fmt"
	//"html/template"
	"os"

)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./a-frame")))
//	http.HandleFunc("/",  func(w http.ResponseWriter, r *http.Request) {
//		templ := template.Must(template.ParseFiles(filepath.Join("templates", "main.html")))
//		msg := "test"
//		templ.Execute(w, msg)
//	})
//	
//	http.HandleFunc("/basic",  func(w http.ResponseWriter, r *http.Request) {
//		templ := template.Must(template.ParseFiles(filepath.Join("a-frame/basic", "index.html")))
//		msg := "test"
//		templ.Execute(w, msg)
//	})
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("ERROR : %s\n", err)
		os.Exit(1)
		
	}
	
	

}
