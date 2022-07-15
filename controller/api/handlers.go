package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ozhey/concordance/controller/database"
	"net/http"
)

func listArticles(c *gin.Context) {
	articles, err := db.ListArticles()
	handleResponse(c, articles, err)
}

func listWordGroups(c *gin.Context) {
	wordGroups, err := db.ListWordGroups()
	handleResponse(c, wordGroups, err)
}

func listLinguisticExpr(c *gin.Context) {
	exprs, err := db.ListLinguisticExpr()
	handleResponse(c, exprs, err)
}

func getArticle(c *gin.Context) {
	article, err := db.GetArticle(c.Params.ByName("id"))
	handleResponse(c, article, err)
}

func getWordsIndex(c *gin.Context) {
	articleWords, err := db.GetWordsIndex(c.Query("article_id"), c.Query("word_group_id"))
	handleResponse(c, articleWords, err)
}

func getWordByPosition(c *gin.Context) {
	word, err := db.GetWordByPosition(c.Query("article_id"), c.Query("page_num"), c.Query("line_num"), c.Query("word_num"))
	handleResponse(c, word, err)
}

func createArticle(c *gin.Context) {
	var body db.NewArticle
	err := c.BindJSON(&body)
	if err != nil {
		handleResponse(c, nil, err)
	}

	article, err := db.CreateArticle(body)
	handleResponse(c, article, err)
}

func createWordGroup(c *gin.Context) {
	var body db.WordGroup
	err := c.BindJSON(&body)
	if err != nil {
		handleResponse(c, nil, err)
	}

	article, err := db.CreateWordGroup(body)
	handleResponse(c, article, err)
}

func createLinguisticExpr(c *gin.Context) {
	var body db.LinguisticExpr
	err := c.BindJSON(&body)
	if err != nil {
		handleResponse(c, nil, err)
	}

	article, err := db.CreateLinguisticExpr(body)
	handleResponse(c, article, err)
}

func handleResponse(c *gin.Context, res any, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res})
}
