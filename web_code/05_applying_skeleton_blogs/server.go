package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Slug        string `json:"Slug"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage !!")
	fmt.Println("Endpoint : homePage")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "How good is go??", Description: "This article covers aspects of go programming language !!", Slug: "how-good-is-go"},
		Article{Title: "How good is go??", Description: "This article covers aspects of go programming language !!", Slug: "how-good-is-go"},
	}

	fmt.Println("Endpoint : getAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", getAllArticles)

	log.Fatal(http.ListenAndServe(":8080", nil))

	/*
	  http://localhost:8080/
	  http://localhost:8080/all
	  http://localhost:8080/single
	  http://localhost:8080/add
	  http://localhost:8080/del
	  http://localhost:8080/edit
	*/

}
