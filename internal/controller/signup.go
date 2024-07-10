package controller

import (
	"github.com/dilyara4949/flight-booking-api/internal/service"
	"log/slog"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service service.AuthService
	Config  config.Config
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SignupResponse struct {
	AccessToken string
	User        domain.User
}

func NewAuthController(service service.AuthService, cfg config.Config) *AuthController {
	return &AuthController{Service: service, Config: cfg}
}

func (controller *AuthController) Signup(c *gin.Context) {
	var request domain.Signup

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

	err = controller.Service.CreateUser(c, &user, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	token, err := controller.Service.CreateAccessToken(c, user, controller.Config.JWTTokenSecret, controller.Config.AccessTokenExpire)
	if err != nil {
		slog.Error("signup: error at creating access token,", err)

		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "create access token error"})
		return
	}

	response := SignupResponse{
		AccessToken: token,
		User:        user,
	}
	c.JSON(http.StatusOK, response)
}
