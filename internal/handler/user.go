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
	ResetPassword(ctx context.Context, req request.ResetPassword, requirePasswordReset bool) error
	ValidateUser(ctx context.Context, signin request.Signin) (domain.User, error)
	Get(ctx context.Context, id uuid.UUID) (domain.User, error)
}

const userIDParamKey = "userId"

func GetUserHandler(service UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !AccessCheck(*c, c.GetString("user_id"), userIDParamKey) {
			c.JSON(http.StatusForbidden, response.Error{Error: "access denied"})
			return
		}

		userID, err := uuid.Parse(c.Param(userIDParamKey))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "id is not correct"})
			return
		}

		user, err := service.Get(c, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, response.Error{Error: "user not found"})
			return
		}

		c.JSON(http.StatusOK, domainUserToResponse(user))
	}
}
