package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserService interface {
	CreateUser(ctx context.Context, signup request.Signup, password string) (domain.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

func UpdateUserHandler(userService UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.UpdateUser

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "error at binding request body"})
			return
		}

		if req.Email == "" || req.Phone == "" {
			c.JSON(http.StatusBadRequest, response.Error{Error: "fields cannot be empty"})
			return
		}

		userID, err := uuid.Parse(c.Param("userId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "id format is not correct"})
			return
		}

	}
}
