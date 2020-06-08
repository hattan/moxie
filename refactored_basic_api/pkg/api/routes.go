package api

import (
	m "github.com/hattan/moxie/pkg/api/models"
)

func (router *ApiRouter) getRoutes() m.Routes {
	return m.Routes{
		m.Route{
			"Index",
			"GET",
			"/",
			router.articleHandler.homePage,
		},
		m.Route{
			"Get All Articles",
			"GET",
			"/articles",
			router.articleHandler.returnAllArticles,
		},
		m.Route{
			"Get Article By Id",
			"GET",
			"/article/{id}",
			router.articleHandler.returnSingleArticle,
		},
	}
}
