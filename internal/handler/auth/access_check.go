package auth

import (
	"github.com/gin-gonic/gin"
)

const (
	adminRole   = "admin"
	userRoleKey = "user_role"
)

func AccessCheck(req gin.Context, expectedContextID, expectedIDKey string) bool {
	role, exists := req.Get(userRoleKey)
	if !exists {
		return false
	}

	userRole, ok := role.(string)
	if !ok {
		return false
	}

	userID := req.Param(expectedIDKey)
	if userRole == adminRole || expectedContextID == userID && expectedContextID != "" {
		return true
	}

	return false
}
