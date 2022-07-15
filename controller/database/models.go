package database

import (
	"time"
)

type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Article struct {
	Base
	PublishedAt  time.Time     `json:"published_at"`
	Author       string        `json:"author"`
	Title        string        `json:"title"`
	Source       string        `json:"source"`
	ArticlePages []ArticlePage `json:",omitempty"`
	PagesCount   int
}

type ArticlePage struct {
	Base
	ArticleID    uint
	PageNumber   int
	ArticleLines []ArticleLine `json:",omitempty"`
}

type ArticleLine struct {
	Base
	ArticlePageID uint
	LineNumber    int
	ArticleWords  []ArticleWord `json:",omitempty"`
	WordCount     int
}

type ArticleWord struct {
	Base
	ArticleLineID uint
	WordNumber    int
	Word          string
	CharCount     int
}

type WordGroup struct {
	Base
	Name  string `json:"name" binding:"required"`
	Words []Word `json:"words" binding:"required"`
}

type Word struct {
	Base
	WordGroupID uint
	Word        string `json:"word" binding:"required"`
}

type LinguisticExpr struct {
	Base
	Expression string `json:"expression" binding:"required"`
}

/* rename tables */

func (a ArticleLine) TableName() string {
	return "article_lines"
}

func (a ArticleWord) TableName() string {
	return "article_words"
}

func (a ArticlePage) TableName() string {
	return "article_pages"
}

func (a Word) TableName() string {
	return "words"
}
