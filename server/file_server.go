package main

import (
	"log"
	"net/http"
)

var ROOT_PATH string = "E:\\cache1\\"

func main() {
	err := http.ListenAndServe(":8181", http.FileServer(http.Dir(ROOT_PATH)))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
