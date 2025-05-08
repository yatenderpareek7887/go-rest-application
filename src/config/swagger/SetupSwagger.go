package swagger

import (
	_ "go-rest-application/src/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func SetupSwagger(router *gin.Engine) {
	router.GET("/swagger/api/docs/*any", gs.WrapHandler(swaggerFiles.Handler))
}
