package database

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func ListArticles() (any, error) {
	var articles []Article
	res := DB.Raw(getArticles).Scan(&articles)
	return handleQueryResult(articles, res)
}

func ListWordGroups() (any, error) {
	var wordGroups []WordGroup
	res := DB.Raw(getWordGroups).Scan(&wordGroups)
	return handleQueryResult(wordGroups, res)
}

func ListLinguisticExpr() (any, error) {
	var lingExprs []LinguisticExpr
	res := DB.Raw(getLinguisticExprs).Scan(&lingExprs)
	return handleQueryResult(lingExprs, res)
}

func GetArticle(id string) (any, error) {
	var article struct {
		Article     string `json:"content"`
		WordsInLine string `json:"avg_words_in_line"`
		CharsInWord string `json:"avg_chars_in_word"`
		PagesNum    string `json:"avg_pages_num"`
	}
	res := DB.Raw(getArticleByID, id).Scan(&article)
	return handleQueryResult(article, res)

}

func GetWordsIndex(articleID string, wordGroupName string) (any, error) {
	var articleWords []struct {
		Word  string
		Count int
		Index string
	}
	articleFilter := "1=1"
	if articleID != "" {
		articleFilter = fmt.Sprintf("a.id = %s", articleID)
	}
	wordGroupFilter := "1=1"
	if wordGroupName != "" {
		wordGroupFilter = fmt.Sprintf(wordsIndexWithWordGroup, wordGroupName)
	}
	res := DB.Raw(getWordsIndex, gorm.Expr(articleFilter), gorm.Expr(wordGroupFilter)).Scan(&articleWords)
	return handleQueryResult(articleWords, res)

}

func GetWordByPosition(articleID string, pageNum string, lineNum string, wordNum string) (any, error) {
	if articleID == "" || pageNum == "" || lineNum == "" || wordNum == "" {
		return nil, errors.New("one of the parameters is missing")
	}

	var word string
	res := DB.Raw(getWordByPosition, articleID, pageNum, lineNum, wordNum).Scan(&word)
	return handleQueryResult(word, res)
}

func CreateArticle(newArticle NewArticle) (any, error) {
	articleToInsert, err := parseArticle(newArticle)
	if err != nil {
		return nil, err
	}

	return handleCreate(articleToInsert)
}

func CreateWordGroup(group WordGroup) (any, error) {
	return handleCreate(group)
}

func CreateLinguisticExpr(expr LinguisticExpr) (any, error) {
	return handleCreate(expr)
}

func handleCreate(obj any) (any, error) {
	res := DB.Create(&obj)
	if res.Error != nil {
		return nil, res.Error
	}
	return obj, nil
}

func handleQueryResult(res any, tx *gorm.DB) (any, error) {
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return res, nil
}
