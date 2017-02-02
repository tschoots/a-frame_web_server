package main 

import (
	"net/http"

)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./a-frame")))
	http.ListenAndServe(":8080", nil)

}

