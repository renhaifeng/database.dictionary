package routers

import (
	"database.dictionary/controllers"
	"database.dictionary/setting"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.LoadHTMLGlob("resource/*")

	r.GET("/", controllers.IndexHandler)

	return r
}
