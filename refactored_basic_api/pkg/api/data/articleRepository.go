package data

import (
	"github.com/google/uuid"
	m "github.com/hattan/moxie/pkg/api/models"
)

type ArticleRepository interface {
	ReturnAllArticles() m.Articles
	ReturnSingleArticle(id uuid.UUID) *m.Article
	AddArticle(article m.Article)
}
