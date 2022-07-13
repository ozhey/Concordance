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

func getArticle(c *gin.Context) {
	article, err := db.GetArticle(c.Params.ByName("id"))
	handleResponse(c, article, err)
}

func getArticleWords(c *gin.Context) {
	articleWords, err := db.GetArticleWords(c.Params.ByName("id"), "")
	handleResponse(c, articleWords, err)
}

func getWordByPosition(c *gin.Context) {

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

func handleResponse(c *gin.Context, res any, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res})
}
