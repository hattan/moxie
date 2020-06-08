package api

import (
	"net/http"

	"github.com/golobby/container"
	"github.com/gorilla/mux"
	m "github.com/hattan/moxie/pkg/api/models"
)

type ApiRouter struct {
	articleHandler ArticleHandler
	*mux.Router
}

func NewRouter() *ApiRouter {
	var articleHandler ArticleHandler
	container.Make(&articleHandler)

	router := &ApiRouter{
		articleHandler,
		mux.NewRouter().StrictSlash(true),
	}
	routes := router.getRoutes()

	for _, route := range routes {
		router.createHandler(route)
	}
	return router
}

func (router *ApiRouter) createHandler(route m.Route) {
	var handler http.Handler
	handler = route.HandlerFunc

	router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)
}
