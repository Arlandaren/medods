package handlers

import (
	"fmt"
	"net/http"
	"os"
	"server/pkg/services"
	"server/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateTokens(s *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")

		ip := c.ClientIP()

		tokens, err := utils.GenerateTokens(user_id, ip, s)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tokens)

	}
}

func RefreshAccessToken(s *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		// err := godotenv.Load()
		// if err != nil {
		// 	c.JSON(500, gin.H{"reason": "could not load environment variables"})
		// 	return 
		// }
		ip := c.ClientIP()
		var request struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}

		claims := &utils.Claims{}
		_, err := jwt.ParseWithClaims(request.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if ip != claims.IP {
			fmt.Println(utils.SendWarningEmail(claims.UserID))
		}

		vaild, err := utils.ValidateRefreshToken(claims.UserID, request.RefreshToken, s)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "inactive refresh token"})
			return
		}

		if vaild {
			newAccess, err := utils.GenerateAccessToken(claims.UserID, ip)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to generate new access token"})
				return
			}
			request.AccessToken = newAccess
			c.JSON(http.StatusOK, request)
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "refresh in inactive"})
			return
		}
	}
}
