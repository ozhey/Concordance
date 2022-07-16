package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/articles", createArticle)
	r.GET("/articles", listArticles)
	r.GET("/articles/index", getWordsIndex)
	r.GET("/articles/:id", getArticle)
	r.GET("/articles/:id/ling_expr_pos", getLingExprPos)

	r.GET("/article_words", getWordByPosition)

	r.GET("/word_group", listWordGroups)
	r.POST("/word_group", createWordGroup)
	r.POST("/word_group/:id", addWordToWordGroup)

	r.GET("/ling_expr", listLinguisticExpr)
	r.POST("/ling_expr", createLinguisticExpr)
	return r
}
