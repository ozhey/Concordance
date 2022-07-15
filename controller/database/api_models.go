package database

// this file contains models for possible request body payloads

type NewArticle struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	PublishedAt string `json:"published_at" binding:"required"`
	Source      string `json:"source" binding:"required"`
	RawContent  string `json:"content" binding:"required"`
}
