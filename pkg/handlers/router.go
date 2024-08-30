package handlers

import (
	"server/pkg/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, s *services.Services) {
	api_router := r.Group("api")
	{
		api_router.Group("auth")
		{
			api_router.GET("/token/:user_id", GenerateTokens(s))
			api_router.POST("/token/refresh", RefreshAccessToken(s))
		}
	}

}