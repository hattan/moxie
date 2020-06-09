package models

import "github.com/google/uuid"

type Article struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Desc    string    `json:"desc"`
	Content string    `json:"content"`
}

type Articles []Article
