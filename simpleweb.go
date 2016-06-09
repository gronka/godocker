package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	fmt.Fprintf(w, "hello")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("./file")
	check(err)
	fmt.Print(string(dat))
	fmt.Fprintf(w, string(dat))
}

func main() {
	http.HandleFunc("/hello", sayhelloName)
	http.HandleFunc("/read", readFile)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
