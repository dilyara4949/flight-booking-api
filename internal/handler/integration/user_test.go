//go:build integration
// +build integration

package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/database/postgres"
	"github.com/dilyara4949/flight-booking-api/internal/handler"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/dilyara4949/flight-booking-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initDB(cfg config.Config) (*gorm.DB, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := postgres.Connect(ctx, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %v", err)
	}
	return database, nil
}

func TestUpdateUserHandler(t *testing.T) {
	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("error getting config: %v", err)
	}

	database, err := initDB(cfg)
	if err != nil {
		t.Fatalf("init database failed: %v", err)
	}

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)

	tests := map[string]struct {
		userID          string
		role            string
		updateReq       request.UpdateUser
		expectedStatus  int
		expectedMessage string
		expectedUser    *response.User
	}{
		"OK user role": {
			userID: "5a57c98d-87a0-436b-a016-634622efbf4e",
			role:   "user",
			updateReq: request.UpdateUser{
				Email: "newuser@mail.ru",
				Phone: "88888888888",
			},
			expectedStatus:  http.StatusOK,
			expectedMessage: "",
			expectedUser: &response.User{
				ID:    uuid.MustParse("5a57c98d-87a0-436b-a016-634622efbf4e"),
				Email: "newuser@mail.ru",
				Phone: "88888888888",
			},
		},
		"invalid id": {
			userID: "5a57c98d-87a0-436b-634622efbf4e",
			role:   "user",
			updateReq: request.UpdateUser{
				Email: "user@mail.ru",
				Phone: "88888888888",
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "id format is not correct",
		},
		"non-existent user": {
			userID: "5a57c98d-87a0-436b-436b-034622efbf4e",
			role:   "user",
			updateReq: request.UpdateUser{
				Email: "qwerty@mail.ru",
				Phone: "88888888888",
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "user not found",
		},
		"change role without access": {
			userID: "5a57c98d-87a0-436b-a016-634622efbf4e",
			role:   "user",
			updateReq: request.UpdateUser{
				Email: "newuser@mail.ru",
				Phone: "88888888888",
				Role:  "admin",
			},
			expectedStatus:  http.StatusForbidden,
			expectedMessage: "access denied: not possible to change role",
		},
		"change role with access": {
			userID: "5c40fa54-cf37-4e7c-8ebb-d883c7e31f96",
			role:   "admin",
			updateReq: request.UpdateUser{
				Email: "newadmin@mail.ru",
				Phone: "88888888888",
				Role:  "admin",
			},
			expectedStatus:  http.StatusOK,
			expectedMessage: "",
			expectedUser: &response.User{
				ID:    uuid.MustParse("5c40fa54-cf37-4e7c-8ebb-d883c7e31f96"),
				Email: "newadmin@mail.ru",
				Phone: "88888888888",
			},
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

				if tt.userID != "" {
					c.Set("user_id", tt.userID)
				}
				c.Next()
			})

			router.PUT("/users/:userId", handler.UpdateUserHandler(userService))

			body, _ := json.Marshal(tt.updateReq)
			req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/users/%s", tt.userID), bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("user_id", tt.userID)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, w.Code)
			}

			if tt.expectedMessage != "" {
				var resp map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				if err != nil {
					t.Fatalf("error unmarshaling response: %v", err)
				}
				if resp["error"] != tt.expectedMessage {
					t.Errorf("expected message %v, got %v", tt.expectedMessage, resp["error"])
				}
			} else if tt.expectedUser != nil {
				var resp response.User
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				if err != nil {
					t.Fatalf("error unmarshaling response: %v", err)
				}
				if resp.ID != tt.expectedUser.ID || resp.Email != tt.expectedUser.Email || resp.Phone != tt.expectedUser.Phone {
					t.Errorf("expected user %v, got %v", tt.expectedUser, resp)
				}
			}
		})
	}
}
