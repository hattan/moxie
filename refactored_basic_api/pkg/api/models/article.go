package models

import "github.com/google/uuid"

type Article struct {
	Id      uuid.UUID `json:"Id"`
	Title   string    `json:"Title"`
	Desc    string    `json:"desc"`
	Content string    `json:"content"`
}

type Articles []Article
