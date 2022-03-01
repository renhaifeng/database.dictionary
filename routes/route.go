package routers

import (
	"database.dictionary/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("resource/*")

	r.GET("/", controllers.IndexHandler)

	return r
}
