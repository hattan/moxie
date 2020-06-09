package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golobby/container"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	d "github.com/hattan/moxie/pkg/api/data"
	m "github.com/hattan/moxie/pkg/api/models"
)

type ArticleHandler interface {
	homePage(w http.ResponseWriter, r *http.Request)
	returnAllArticles(w http.ResponseWriter, r *http.Request)
	returnSingleArticle(w http.ResponseWriter, r *http.Request)
	createArticle(w http.ResponseWriter, r *http.Request)
}

type ArticleService struct {
	repository d.ArticleRepository
}

func NewArticleService() *ArticleService {
	var repository d.ArticleRepository
	container.Make(&repository)

	service := &ArticleService{}
	service.repository = repository
	return service
}

func (service ArticleService) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func (service ArticleService) returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(service.repository.ReturnAllArticles())
}

func (service ArticleService) createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article m.Article
	article.Id = uuid.New()
	json.Unmarshal(reqBody, &article)
	service.repository.AddArticle(article)
	w.WriteHeader(201)
}

func (service ArticleService) returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")

	vars := mux.Vars(r)
	key, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Invalid key format, %s", vars["id"])
		w.WriteHeader(500)
		fmt.Fprintf(w, "Invalid key format %s", vars["id"])
		return
	}

	article := service.repository.ReturnSingleArticle(key)
	if article == nil {
		w.WriteHeader(404)
		log.Printf("Couldn't find article with id: %s", vars["id"])
	}

	json.NewEncoder(w).Encode(article)
}
