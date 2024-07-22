package middleware

import (
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
)

const userRole = "user"

func AccessCheck(role string, param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if role == userRole {
			contextParam := c.GetString(param)
			pathParam := c.Param(param)

			if contextParam != pathParam {
				c.AbortWithStatusJSON(http.StatusForbidden, response.Error{Error: "access denied"})
				return
			}
		}
		c.Next()
	}
}
