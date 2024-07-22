package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthService interface {
	CreateAccessToken(ctx context.Context, user domain.User, secret string, expiry int) (accessToken string, err error)
}

type UserService interface {
	CreateUser(ctx context.Context, signup request.Signup, password string) (domain.User, error)
	ResetPassword(ctx context.Context, userID uuid.UUID, newPassword string) error
}

func SignupHandler(authService AuthService, userService UserService, cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.Signup

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "incorrect req body"})

			return
		}

		if req.Password == "" || req.Email == "" || req.Role == "" {
			c.JSON(http.StatusBadRequest, response.Error{Error: "fields cannot be empty"})

			return
		}

		user, err := userService.CreateUser(c, req, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})

			return
		}

		token, err := authService.CreateAccessToken(c, user, cfg.JWTTokenSecret, cfg.AccessTokenExpire)
		if err != nil {
			slog.Error("signup: error at creating access token,", err)

			c.JSON(http.StatusInternalServerError, response.Error{Error: "create access token error"})

			return
		}

		resp := response.Signup{
			AccessToken: token,
			User:        domainUserToResponse(user),
		}
		c.JSON(http.StatusOK, resp)
	}
}

func domainUserToResponse(user domain.User) response.User {
	return response.User{
		ID:        user.ID,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
