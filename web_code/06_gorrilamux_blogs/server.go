package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	if r.Method == "GET" {
		articles := Articles{
			Article{Title: "How good is go??", Description: "This article covers aspects of go programming language !!", Slug: "how-good-is-go"},
			Article{Title: "How good is go??", Description: "This article covers aspects of go programming language !!", Slug: "how-good-is-go"},
		}

		fmt.Println("Endpoint : getAllArticles")
		json.NewEncoder(w).Encode(articles)

	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	var1 := vars["var1"]
	fmt.Fprintf(w, "Key: "+key+" var "+var1)

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", getAllArticles)
	myRouter.HandleFunc("/article/{key}/{var1}", getArticle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
