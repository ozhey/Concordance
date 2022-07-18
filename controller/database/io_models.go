package database

// models for possible request body payloads

type NewArticle struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	PublishedAt string `json:"published_at" binding:"required"`
	Source      string `json:"source" binding:"required"`
	RawContent  string `json:"content" binding:"required"`
}

// models for possible response bodies

type textLines []struct {
	Content    string `json:"content"`
	LineNumber int    `json:"line_number"`
}
type wordByPositionRes struct {
	Lines textLines `json:"lines"`
	Word  string    `json:"word"`
}

type wordsRes []struct {
	PageNumber int    `json:"page_number"`
	LineNumber int    `json:"line_number"`
	WordNumber int    `json:"word_number"`
	Word       string `json:"word"`
}
