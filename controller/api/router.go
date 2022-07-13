package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/articles", listArticles)
	r.GET("/articles/:id", getArticle)
	r.GET("/articles/:id/words", getArticleWords)
	r.GET("words", getWordByPosition)
	r.POST("/articles", createArticle)
	return r
}
