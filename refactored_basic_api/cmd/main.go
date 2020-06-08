package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golobby/container"
	"github.com/hattan/moxie/pkg/api"
	"github.com/hattan/moxie/pkg/api/data"
)

func initialize() {
	container.Singleton(func() data.ArticleRepository {
		return data.NewArticleMemoryRepository()
	})
	container.Singleton(func() api.ArticleHandler {
		return api.NewArticleService()
	})

}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers!")
	initialize()

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":10000", router))
}
