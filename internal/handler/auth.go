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
	AdminRole = "admin"
	UserRole  = "user"
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

		//respByte, err := json.Marshal(resp)
		//if err != nil {
		//	panic(err)
		//}
		//
		//err = producer.Produce(&kafka.Message{
		//	TopicPartition: kafka.TopicPartition{Topic: &cfg.Kafka.EmailPush, Partition: kafka.PartitionAny},
		//	Value:          respByte,
		//}, nil)

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

		if user.RequirePasswordReset {
			c.JSON(http.StatusForbidden, response.Error{Error: "access denied: reset password required"})

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

func AccessCheck(req *gin.Context, expectedContextID, expectedIDKey string) bool {
	role, exists := req.Get(middleware.UserRoleKey)
	if !exists {
		return false
	}

	userRole, ok := role.(string)
	if !ok {
		return false
	}

	userID := req.Param(expectedIDKey)
	if userRole == AdminRole || expectedContextID == userID && expectedContextID != "" {
		return true
	}

	return false
}

func ResetPasswordHandler(userService UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.ResetPassword

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "error at binding request body"})
			return
		}

		if req.NewPassword == "" || req.OldPassword == "" || req.Email == "" {
			c.JSON(http.StatusBadRequest, response.Error{Error: response.ErrEmptyRequestFields.Error()})
			return
		}

		err = userService.ResetPassword(c, req, false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, "password reset successful")
	}
}
