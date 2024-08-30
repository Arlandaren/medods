package handlers

import (
	"server/pkg/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, s *services.Services) {
	api_router := r.Group("api")
	{
		auth := api_router.Group("auth")
		{
			auth.GET("/token/:user_id", GenerateTokens(s))
			auth.POST("/token/refresh", RefreshAccessToken(s))
		}
	}

}