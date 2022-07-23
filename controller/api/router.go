package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	r.GET("/articles", listArticles)
	r.POST("/articles", createArticle)
	r.GET("/articles/:id", getArticle)
	r.GET("/articles/:id/ling_expr_pos", getLingExprPos)

	r.GET("/article_words", getWordByPosition)
	r.GET("/article_words/index", getWordsIndex)

	r.GET("/word_group", listWordGroups)
	r.POST("/word_group", createWordGroup)
	r.POST("/word_group/:id", addWordToWordGroup)

	r.GET("/ling_expr", listLinguisticExpr)
	r.POST("/ling_expr", createLinguisticExpr)
	return r
}
