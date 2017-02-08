package main

import (
    "testing"
    "fmt"
    "path/filepath"
)

func TestSearch(t *testing.T) {
	
	webpages := Webpages{}
	
	path, _ := filepath.Abs(".")
	
	pageMap := webpages.findPages(path)
	
	pageMap = webpages.findPages(".")
	
	
	
	
	for k, v := range *pageMap {
		fmt.Printf("key : %s , value : %s\n", k, v)
	}

}

