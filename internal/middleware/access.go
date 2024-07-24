package middleware

import (
	"log"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
)

const UserRoleKey = "user_role"

func AccessCheck(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(UserRoleKey)
		if !exists {
			log.Println("22222222222222222222")
			c.AbortWithStatusJSON(http.StatusForbidden, response.Error{Error: "access denied"})
			return
		}

		userRole, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, response.Error{Error: "invalid role type"})
			return
		}

		log.Println("11111111111111111111", userRole)
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, response.Error{Error: "access denied"})
	}
}
