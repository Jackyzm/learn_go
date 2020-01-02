package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"learn_go/pkg/response"
	"learn_go/pkg/utils"
)

// Jwt is jwt middleware
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := response.Gin{C: c}

		const bearerLength = len("Bearer ")
		token := c.GetHeader("Authorization")
		if len(token) < bearerLength {
			response.SetBadResponse(http.StatusUnauthorized, "Unauthorized")
			return
		}

		token = strings.TrimSpace(token[bearerLength:])

		code := http.StatusOK
		massage := "Unauthorized"

		if token == "" {
			code = http.StatusUnauthorized
			massage = "Unauthorized"
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
				if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
					massage = "token_expired"
				}
			}
		}

		if code != http.StatusOK {
			response.SetBadResponse(http.StatusUnauthorized, massage)
			return
		}

		c.Next()
	}
}
