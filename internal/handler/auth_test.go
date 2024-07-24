package handler

import (
	"testing"

	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	userRole = "user"
)

func TestAccessCheck(t *testing.T) {
	tests := map[string]struct {
		role              any
		expectedContextID string
		expectedIDKey     string
		paramValue        string
		expectedResult    bool
	}{
		"admin role": {
			role:              adminRole,
			expectedContextID: "1",
			expectedIDKey:     "user_id",
			paramValue:        "2",
			expectedResult:    true,
		},
		"user role, same IDs": {
			role:              userRole,
			expectedContextID: "1",
			expectedIDKey:     "user_id",
			paramValue:        "1",
			expectedResult:    true,
		},
		"user role, not same IDs": {
			role:              userRole,
			expectedContextID: "1",
			expectedIDKey:     "user_id",
			paramValue:        "2",
			expectedResult:    false,
		},
		"role key not set": {
			expectedContextID: "1",
			expectedIDKey:     "user_id",
			paramValue:        "1",
			expectedResult:    false,
		},
		"role not a string": {
			role:              123,
			expectedContextID: "1",
			expectedIDKey:     "user_id",
			paramValue:        "1",
			expectedResult:    false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			ctx, _ := gin.CreateTestContext(nil)
			if tt.role != "" {
				ctx.Set(middleware.UserRoleKey, tt.role)
			}

			ctx.Params = gin.Params{
				{Key: tt.expectedIDKey, Value: tt.paramValue},
			}

			result := AccessCheck(*ctx, tt.expectedContextID, tt.expectedIDKey)
			if tt.expectedResult != result {
				t.Errorf("expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
