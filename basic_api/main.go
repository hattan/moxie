package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Invalid key format, %s", vars["id"])
		w.WriteHeader(500)
		fmt.Fprintf(w, "Invalid key format %s", vars["id"])
		return
	}

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	article.Id = uuid.New()
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("here")
	vars := mux.Vars(r)

	key, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Invalid key format, %s", vars["id"])
		w.WriteHeader(500)
		fmt.Fprintf(w, "Invalid key format %s", vars["id"])
		return
	}
	for index, article := range Articles {
		if article.Id == key {
			log.Printf("found a key %s", key)
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers!")

	Articles = []Article{
		{Id: uuid.New(), Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: uuid.New(), Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}

type Article struct {
	Id      uuid.UUID `json:"Id"`
	Title   string    `json:"Title"`
	Desc    string    `json:"desc"`
	Content string    `json:"content"`
}
