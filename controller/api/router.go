package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/articles", createArticle)
	r.GET("/articles", listArticles)
	r.GET("/articles/:id", getArticle)
	r.GET("/articles/words", getWordsIndex)

	r.GET("/article_words", getWordByPosition)

	r.GET("/word_group", listWordGroups)
	r.POST("/word_group", createWordGroup)

	r.GET("/linguistic_expr", listLinguisticExpr)
	r.POST("/linguistic_expr", createLinguisticExpr)
	return r
}
