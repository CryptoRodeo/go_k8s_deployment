package router

import (
	"go_user_microservice/router/api/v_1"
	"go_user_microservice/utils/settings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(settings.Cors))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/users", v_1.GetAllUsers)
		apiv1.GET("/user/:id", v_1.GetUserByID)
		apiv1.GET("/user/search", v_1.GetUsersByID)
		apiv1.GET("/user/:id/tickets", v_1.GetUsersTickets)
	}

	return r
}
