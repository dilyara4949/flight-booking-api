package handler

import (
	"context"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, signup request.Signup, password string) (domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ValidateUser(ctx context.Context, signin request.Signin) (domain.User, error)
	ResetPassword(ctx context.Context, req request.ResetPassword, requirePasswordReset bool) error
}

func DeleteUserHandler(userService UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !AccessCheck(*c, c.GetString("user_id"), "userId") {
			c.JSON(http.StatusForbidden, response.Error{Error: "access denied"})
			return
		}

		userID, err := uuid.Parse(c.Param("userId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "id format is not correct"})
			return
		}

		err = userService.DeleteUser(c, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, response.Error{Error: "user not found"})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
