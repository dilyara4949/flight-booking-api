package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dilyara4949/flight-booking-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func TestAccessCheck(t *testing.T) {
	tests := map[string]struct {
		role             interface{}
		allowedRoles     []string
		expectedStatus   int
		expectedResponse string
	}{
		"Allowed role": {
			role:             "admin",
			allowedRoles:     []string{"admin", "user"},
			expectedStatus:   http.StatusOK,
			expectedResponse: "success",
		},
		"Forbidden role": {
			role:             "guest",
			allowedRoles:     []string{"admin", "user"},
			expectedStatus:   http.StatusForbidden,
			expectedResponse: `{"error":"access denied"}`,
		},
		"No role": {
			role:             "",
			allowedRoles:     []string{"admin", "user"},
			expectedStatus:   http.StatusForbidden,
			expectedResponse: `{"error":"access denied"}`,
		},
		"Invalid role type": {
			role:             123,
			allowedRoles:     []string{"admin", "user"},
			expectedStatus:   http.StatusForbidden,
			expectedResponse: `{"error":"invalid role type"}`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.New()

			router.Use(func(c *gin.Context) {
				if tt.role != "" {
					c.Set(middleware.UserRoleKey, tt.role)
				}
				c.Next()
			})
			router.Use(middleware.AccessCheck(tt.allowedRoles...))

			router.GET("/", func(c *gin.Context) {
				c.String(http.StatusOK, "success")
			})

			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Fatalf("expected status code %d, got %d", tt.expectedStatus, w.Code)
			}

			if w.Body.String() != tt.expectedResponse {
				t.Fatalf("expected response body %s, got %s", tt.expectedResponse, w.Body.String())
			}
		})
	}
}
