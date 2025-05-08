package routes

import (
	controllers "go-rest-application/src/controllers/user-controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("/users", controllers.GetAllUsers)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.PATCH("/users/:id", controllers.PartialUpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	return r
}
