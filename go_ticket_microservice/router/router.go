package router

import (
	"go_ticket_microservice/router/api/v_1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tickets", v_1.GetAllTickets)
		apiv1.GET("/tickets/search", v_1.SearchTickets)
	}

	return r
}
