package controller

import (
	"log"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	Service domain.AuthService
	Config  config.Config
}

func NewAuthController(service domain.AuthService, cfg config.Config) *AuthController {
	return &AuthController{Service: service, Config: cfg}
}

func (controller *AuthController) Signup(c *gin.Context) {
	var request domain.Signup

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "incorrect request body"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "generate password error"})
		return
	}

	user := domain.User{Email: request.Email}

	err = controller.Service.CreateUser(c, &user, string(encryptedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	token, err := controller.Service.CreateAccessToken(user.ID, controller.Config.JWTTokenSecret, controller.Config.AccessTokenExpire)
	if err != nil {
		log.Printf("signup: error at creating acces token, %v", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "create access token error"})
	}

	response := domain.SignupResponse{
		AccessToken: token,
		User:        user,
	}
	c.JSON(http.StatusOK, response)
}
