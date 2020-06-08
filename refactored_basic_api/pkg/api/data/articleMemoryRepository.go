package data

import (
	"github.com/google/uuid"
	m "github.com/hattan/moxie/pkg/api/models"
)

type ArticleMemoryRepository struct {
	articles m.Articles
}

func NewArticleMemoryRepository() *ArticleMemoryRepository {
	repository := &ArticleMemoryRepository{}
	repository.articles = m.Articles{
		{Id: uuid.New(), Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: uuid.New(), Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	return repository
}

func (repository *ArticleMemoryRepository) ReturnAllArticles() m.Articles {
	return repository.articles
}

func (repository *ArticleMemoryRepository) ReturnSingleArticle(id uuid.UUID) *m.Article {
	for _, article := range repository.articles {
		if article.Id == id {
			return &article
		}
	}
	return nil
}
