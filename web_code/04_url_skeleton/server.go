package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage !!")
	fmt.Println("Endpoint : homePage")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All articles (CARD view) !!")
	fmt.Println("Endpoint : getAllArticles")
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Specific Article (Detail view with slug)!!")
	fmt.Println("Endpoint : getArticle")
}

func addArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Added a new article!!")
	fmt.Println("Endpoint : addArticle")
}

func delArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleted an article!!")
	fmt.Println("Endpoint : delArticle")
}

func editArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edited an article!!")
	fmt.Println("Endpoint : editArticle")
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", getAllArticles)
	http.HandleFunc("/single", getArticle)
	http.HandleFunc("/add", addArticle)
	http.HandleFunc("/del", delArticle)
	http.HandleFunc("/edit", editArticle)

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
