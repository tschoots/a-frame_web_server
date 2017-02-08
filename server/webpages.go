package main

import (
	"path/filepath"
	//"regexp"
	"os"
	"fmt"
	"strings"
)


type Webpages struct {
	
}


func (w *Webpages) findPages(searchPath string) (mapping *map[string]string){
	
	filepath.Walk(searchPath, func(path string, fileInfo os.FileInfo, err error) error {
			fmt.Printf("path : %s\n", path)
			
			if strings.Compare(filepath.Base(path), "index.html") == 0 {
				fmt.Println("yes")
			}
			
			return nil
	})
	
	
	the_map := map[string]string{
		"test" : "filepath",
	}
	
	return &the_map;
}

