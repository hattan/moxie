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
	fmt.Println("Endpoint Hit: returnSingleArticle")
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
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	article.Id = uuid.New()
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
	vars := mux.Vars(r)

	key, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Invalid key format, %s", vars["id"])
		w.WriteHeader(500)
		fmt.Fprintf(w, "Invalid key format %s", vars["id"])
		return
	}
	for index, article := range Articles {
		log.Printf("article.Id %s, key %s", article.Id, key)
		if article.Id == key {
			log.Printf("found a key %s", key)
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")
	vars := mux.Vars(r)
	key, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Invalid key format, %s", vars["id"])
		w.WriteHeader(500)
		fmt.Fprintf(w, "Invalid key format %s", vars["id"])
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	for index, a := range Articles {
		log.Printf("article.Id %s, key %s", article.Id, key)
		if a.Id == key {
			log.Printf("found a key %s", key)
			a.Title = article.Title
			a.Desc = article.Desc
			a.Content = article.Content
			Articles[index] = a
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
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
