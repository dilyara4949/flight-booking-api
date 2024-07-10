package handler

import (
	service2 "github.com/dilyara4949/flight-booking-api/internal/service"
	"log/slog"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/gin-gonic/gin"
)

type signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type signupResponse struct {
	AccessToken string
	User        domain.User
}

func SignupHandler(service service2.AuthService, cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request signup

		err := c.ShouldBind(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "incorrect request body"})
			return
		}

		if request.Password == "" || request.Email == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "fields cannot be empty"})
			return
		}

		user := domain.User{Email: request.Email}

		err = service.CreateUser(c, &user, request.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		token, err := service.CreateAccessToken(c, user, cfg.JWTTokenSecret, cfg.AccessTokenExpire)
		if err != nil {
			slog.Error("signup: error at creating access token,", err)

			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "create access token error"})
			return
		}

		response := signupResponse{
			AccessToken: token,
			User:        user,
		}
		c.JSON(http.StatusOK, response)
	}
}
