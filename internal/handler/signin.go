package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/handler/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthService interface {
	CreateAccessToken(ctx context.Context, user domain.User, secret string, expiry int) (accessToken string, err error)
}

type UserService interface {
	//CreateUser(ctx context.Context, signup request.Signup, password string) (domain.User, error)
	GetUser
}

func SigninHandler() gin.HandlerFunc {
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

	}
}
