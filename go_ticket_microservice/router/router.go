package router

import (
	"go_ticket_microservice/router/api/v_1"
	"go_ticket_microservice/utils/settings"

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
		apiv1.GET("/tickets", v_1.GetAllTickets)
		apiv1.GET("/tickets/:id", v_1.GetTicket)
		apiv1.GET("/tickets/search", v_1.SearchTickets)
	}

	return r
}
