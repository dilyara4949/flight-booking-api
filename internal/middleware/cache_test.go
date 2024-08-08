package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetCacheKey(t *testing.T) {
	tests := []struct {
		name       string
		role       string
		url        string
		paramKey   string
		paramValue string
		page       string
		pageSize   string
		expected   string
	}{
		{
			name:       "Flight by ID",
			url:        "/flights/123",
			paramKey:   "flightId",
			paramValue: "123",
			expected:   "flight-123",
		},
		{
			name:     "Flights list",
			url:      "/flights/",
			page:     "1",
			pageSize: "10",
			expected: "flights-page1-size10",
		},
		{
			name:       "User by ID",
			url:        "/users/456",
			paramKey:   "userId",
			role:       "admin",
			paramValue: "456",
			expected:   "user-456",
		},
		{
			name:     "Users list",
			url:      "/users/",
			role:     "admin",
			page:     "1",
			pageSize: "10",
			expected: "users-page1-size10",
		},
		{
			name:     "User does not have access",
			url:      "/users/",
			role:     "user",
			page:     "1",
			pageSize: "10",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)

			if tt.page != "" {
				q := req.URL.Query()
				q.Add("page", tt.page)
				req.URL.RawQuery = q.Encode()
			}

			if tt.pageSize != "" {
				q := req.URL.Query()
				q.Add("page_size", tt.pageSize)
				req.URL.RawQuery = q.Encode()
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			if tt.paramValue != "" {
				c.Params = gin.Params{gin.Param{Key: tt.paramKey, Value: tt.paramValue}}
			}

			if tt.role != "" {
				c.Set(UserRoleKey, tt.role)
			}

			if actual := getCacheKey(c); tt.expected != actual {
				t.Fatalf("expected: %s, got: %s", tt.expected, actual)
			}
		})
	}
}
