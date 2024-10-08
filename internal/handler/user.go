package handler

import (
	"context"
	"github.com/dilyara4949/flight-booking-api/internal/handler/auth"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response/pagination"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, signup request.Signup, password string) (domain.User, error)
	UpdateUser(ctx context.Context, req request.UpdateUser, userID uuid.UUID) (domain.User, error)
	ResetPassword(ctx context.Context, req request.ResetPassword, requirePasswordReset bool) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ValidateUser(ctx context.Context, signin request.Signin) (domain.User, error)
	Get(ctx context.Context, id uuid.UUID) (domain.User, error)
	GetUsers(ctx context.Context, page, pageSize int) ([]domain.User, error)
}

const userIDParamKey = "userId"

func GetUserHandler(service UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !auth.AccessCheck(c, c.GetString("user_id"), userIDParamKey) {
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

func GetUsersHandler(service UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, pageSize := pagination.GetPageInfo(c)

		users, err := service.GetUsers(c, page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, domainUsersToResponse(users))
	}
}

func UpdateUserHandler(userService UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !auth.AccessCheck(c, c.GetString("user_id"), userIDParamKey) {
			c.JSON(http.StatusForbidden, response.Error{Error: "access denied"})
			return
		}

		var req request.UpdateUser

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "error at binding request body"})
			return
		}

		if req.Role != "" {
			if !auth.AccessCheck(c, "", userIDParamKey) {
				c.JSON(http.StatusForbidden, response.Error{Error: "access denied: not possible to change role"})
				return
			}
		}

		userID, err := uuid.Parse(c.Param(userIDParamKey))
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "id format is not correct"})
			return
		}

		user, err := userService.UpdateUser(c, req, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
			return
		}

		resp := domainUserToResponse(user)
		c.JSON(http.StatusOK, resp)
	}
}


func DeleteUserHandler(userService UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !auth.AccessCheck(c, c.GetString("user_id"), "userId") {
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

func domainUserToResponse(user domain.User) response.User {
	return response.User{
		ID:        user.ID,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func domainUsersToResponse(users []domain.User) []response.User {
	res := make([]response.User, 0)

	for _, user := range users {
		res = append(res, response.User{
			ID:        user.ID,
			Email:     user.Email,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return res
}
