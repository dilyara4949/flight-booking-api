package middleware

import (
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
)

func JWTAuth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error{Error: "authorization header required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error{Error: "Bearer token required"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}

			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error{Error: "invalid token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			claimedUID, ok := claims["user_id"].(string)
			if !ok {
				c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: "authorization: no user property in claims"})
				return
			}

			c.Set("user_id", claimedUID)

			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error{Error: "invalid token claims"})
		}
	}
}
