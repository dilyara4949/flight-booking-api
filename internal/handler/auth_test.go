package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
)

const (
	userRole = "user"
)

type authServiceMock struct {
	token string
	err   error
}

type userServiceMock struct {
	user domain.User
	err  error
}

func (s authServiceMock) CreateAccessToken(ctx context.Context, user domain.User, secret string, expiry int) (string, error) {
	if s.err != nil {
		return "", s.err
	}
	return s.token, nil
}

func (s userServiceMock) CreateUser(ctx context.Context, signup request.Signup, password string) (domain.User, error) {
	if s.err != nil {
		return domain.User{}, s.err
	}
	return s.user, nil
}

func (s userServiceMock) ValidateUser(ctx context.Context, signin request.Signin) (domain.User, error) {
	if s.err != nil {
		return domain.User{}, s.err
	}
	return s.user, nil
}

func (s userServiceMock) ResetPassword(ctx context.Context, req request.ResetPassword, requirePasswordReset bool) error {
	return nil
}

func (s userServiceMock) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return nil
}
func (s userServiceMock) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (s userServiceMock) UpdateUser(ctx context.Context, req request.UpdateUser, userID uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (s userServiceMock) GetUsers(ctx context.Context, page int, pageSize int) ([]domain.User, error) {
	return nil, nil
}

func TestSignupHandler(t *testing.T) {
	tests := map[string]struct {
		body         string
		userService  userServiceMock
		authService  authServiceMock
		expectedCode int
		expectedBody string
	}{
		"OK": {
			body: `{
				"email": "test@example.com",
				"password": "password",
				"role": "user"
			}`,
			userService: userServiceMock{
				user: domain.User{
					ID:    uuid.New(),
					Email: "test@example.com",
					Role:  "user",
				},
			},
			authService: authServiceMock{
				token: "token123",
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"access_token":"token123","user":{"id":"%s","email":"test@example.com","phone":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}}`,
		},
		"invalid request body": {
			body:         `{}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"fields cannot be empty"}`,
		},
		"error creating user": {
			body: `{
				"email": "test@example.com",
				"password": "password"
			}`,
			userService: userServiceMock{
				err: errors.New("error creating user"),
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"error creating user"}`,
		},
		"Error creating token": {
			body: `{
				"email": "test@example.com",
				"password": "password"
			}`,
			userService: userServiceMock{
				user: domain.User{
					ID:    uuid.New(),
					Email: "test@example.com",
				},
			},
			authService: authServiceMock{
				err: errors.New("error creating token"),
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"create access token error"}`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			r := gin.New()

			cfg := config.Config{
				JWTTokenSecret:    "secret",
				AccessTokenExpire: 3600,
			}

			r.POST("/signup", SignupHandler(tt.authService, tt.userService, cfg))

			req, err := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("couldn't create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedCode {
				t.Errorf("expected status %d, got %d", tt.expectedCode, resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("couldn't read response body: %v", err)
			}
			defer resp.Body.Close()

			if tt.expectedCode == http.StatusOK {
				expectedBody := strings.Replace(tt.expectedBody, "%s", tt.userService.user.ID.String(), 1)
				var expected, actual map[string]interface{}
				if err = json.Unmarshal([]byte(expectedBody), &expected); err != nil {
					t.Fatalf("couldn't unmarshal expected response: %v", err)
				}
				if err = json.Unmarshal(body, &actual); err != nil {
					t.Fatalf("couldn't unmarshal actual response: %v", err)
				}

				if !reflect.DeepEqual(expected, actual) {
					t.Errorf("expected body %v, got %v", expected, actual)
				}
			} else {
				if string(body) != tt.expectedBody {
					t.Errorf("expected body %s, got %s", tt.expectedBody, body)
				}
			}
		})
	}
}

func TestSigninHandler(t *testing.T) {
	tests := map[string]struct {
		body         string
		userService  userServiceMock
		authService  authServiceMock
		expectedCode int
		expectedBody string
	}{
		"OK": {
			body: `{
				"email": "test@example.com",
				"password": "password"
			}`,
			userService: userServiceMock{
				user: domain.User{
					ID:    uuid.New(),
					Email: "test@example.com",
				},
			},
			authService: authServiceMock{
				token: "token123",
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"access_token":"token123"}`,
		},
		"Invalid request body": {
			body:         `{}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"fields cannot be empty"}`,
		},
		"Error validating user": {
			body: `{
				"email": "test@example.com",
				"password": "password"
			}`,
			userService: userServiceMock{
				err: errors.New("error validating user"),
			},
			expectedCode: http.StatusNotFound,
			expectedBody: `{"error":"invalid email or password"}`,
		},
		"Error creating token": {
			body: `{
				"email": "test@example.com",
				"password": "password"
			}`,
			userService: userServiceMock{
				user: domain.User{
					ID:    uuid.New(),
					Email: "test@example.com",
					Role:  "user",
				},
			},
			authService: authServiceMock{
				err: errors.New("error creating token"),
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"create access token error"}`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			r := gin.New()

			cfg := config.Config{
				JWTTokenSecret:    "secret",
				AccessTokenExpire: 3600,
			}

			r.POST("/signin", SigninHandler(tt.authService, tt.userService, cfg))

			req, err := http.NewRequest(http.MethodPost, "/signin", strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("couldn't create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedCode {
				t.Errorf("expected status %d, got %d", tt.expectedCode, resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("couldn't read response body: %v", err)
			}
			defer resp.Body.Close()

			if tt.expectedCode == http.StatusOK {
				var expected, actual map[string]interface{}
				if err = json.Unmarshal([]byte(tt.expectedBody), &expected); err != nil {
					t.Fatalf("couldn't unmarshal expected response: %v", err)
				}
				if err = json.Unmarshal(body, &actual); err != nil {
					t.Fatalf("couldn't unmarshal actual response: %v", err)
				}

				if !reflect.DeepEqual(expected, actual) {
					t.Errorf("expected body %v, got %v", expected, actual)
				}
			} else {
				if string(body) != tt.expectedBody {
					t.Errorf("expected body %s, got %s", tt.expectedBody, body)
				}
			}
		})
	}
}

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
		"expectedContextID is empty": {
			role:              userRole,
			expectedContextID: "",
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
