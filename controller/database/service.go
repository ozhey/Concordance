package database

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
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
	lineNumInt, err := strconv.Atoi(lineNum)
	if err != nil {
		return nil, errors.Wrap(err, "convert line number to int")
	}

	lines, err := getWordContext(articleID, pageNum, lineNumInt)
	if err != nil {
		return nil, err
	}

	wordNumInt, err := strconv.Atoi(wordNum)
	if err != nil {
		return nil, errors.Wrap(err, "convert word number to int")
	}

	var word string
	for _, line := range lines {
		if line.LineNumber == lineNumInt {
			words := strings.Split(line.Content, " ")
			word = words[wordNumInt-1]
		}
	}
	return wordByPositionResult{
		Lines: lines,
		Word:  word,
	}, nil
}

func getWordContext(articleID string, pageNum string, lineNumInt int) (textLines, error) {
	linesToGet := fmt.Sprintf("(%d,%d,%d)", lineNumInt-1, lineNumInt, lineNumInt+1)
	if lineNumInt%10 == 1 { // first line in page
		linesToGet = fmt.Sprintf("(%d,%d)", lineNumInt, lineNumInt+1)
	} else if lineNumInt%10 == 0 { // last line in page
		linesToGet = fmt.Sprintf("(%d,%d)", lineNumInt-1, lineNumInt)
	}

	lines := textLines{}
	tx := DB.Raw(getContextByPosition, articleID, pageNum, gorm.Expr(linesToGet)).Scan(&lines)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return lines, nil
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
