package handler

import (
	"context"
	"log/slog"
	"net/http"

	errs "github.com/dilyara4949/flight-booking-api/internal/repository/errors"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/dilyara4949/flight-booking-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	CreateAccessToken(ctx context.Context, user domain.User, secret string, expiry int) (accessToken string, err error)
}

const (
	adminRole = "admin"
)

func SignupHandler(authService AuthService, userService UserService, cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.Signup

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "incorrect req body"})

			return
		}

		if req.Password == "" || req.Email == "" {
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
			slog.Error("signup: error at creating access token,", "error", err.Error())

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

func SigninHandler(authService AuthService, userService UserService, cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.Signin

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "incorrect request body"})

			return
		}

		if req.Email == "" || req.Password == "" {
			c.JSON(http.StatusBadRequest, response.Error{Error: "fields cannot be empty"})

			return
		}

		user, err := userService.ValidateUser(c, req)
		if err != nil {
			c.JSON(http.StatusNotFound, response.Error{Error: errs.ErrInvalidEmailPassword.Error()})

			return
		}

		token, err := authService.CreateAccessToken(c, user, cfg.JWTTokenSecret, cfg.AccessTokenExpire)
		if err != nil {
			slog.Error("signin: error at creating access token,", "error", err.Error())

			c.JSON(http.StatusInternalServerError, response.Error{Error: "create access token error"})

			return
		}

		resp := response.Signin{
			AccessToken: token,
		}
		c.JSON(http.StatusOK, resp)
	}
}

func AccessCheck(req gin.Context, expectedContextID, expectedIDKey string) bool {
	role, exists := req.Get(middleware.UserRoleKey)
	if !exists {
		return false
	}

	userRole, ok := role.(string)
	if !ok {
		return false
	}

	userID := req.Param(expectedIDKey)
	if userRole == adminRole || expectedContextID == userID {
		return true
	}

	return false
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
