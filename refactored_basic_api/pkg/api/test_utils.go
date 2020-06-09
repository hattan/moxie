package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/golobby/container"
	"github.com/google/uuid"
	d "github.com/hattan/moxie/pkg/api/data"
	m "github.com/hattan/moxie/pkg/api/models"
)

type APITestCase struct {
	Tag              string
	Method           string
	URL              string
	Body             string
	Status           int
	ExpectedResponse string
	ExpectedJSON     string
	ActualResponse   string
}

func testAPI(router *ApiRouter, method, URL, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, URL, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func SetUpTestFixtures(data m.Articles) {
	container.Singleton(func() d.ArticleRepository {
		repository := NewFakeArticleRepository()
		repository.articles = data
		return repository
	})
	container.Singleton(func() ArticleHandler {
		return NewArticleService()
	})
}

type FakeArticleRepository struct {
	articles m.Articles
}

func NewFakeArticleRepository() *FakeArticleRepository {
	repository := &FakeArticleRepository{}
	return repository
}

func (repository *FakeArticleRepository) ReturnAllArticles() m.Articles {
	return repository.articles
}

func (repository *FakeArticleRepository) AddArticle(article m.Article) {
	repository.articles = append(repository.articles, article)
}

func (repository *FakeArticleRepository) ReturnSingleArticle(id uuid.UUID) *m.Article {
	for _, article := range repository.articles {
		if article.Id == id {
			return &article
		}
	}
	return nil
}
