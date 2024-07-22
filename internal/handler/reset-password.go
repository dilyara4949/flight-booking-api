package handler

import (
	"log/slog"
	"net/http"

	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ResetPasswordHandler(userService UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.ResetPassword

		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error{Error: "error at binding request body"})
			return
		}

		if req.NewPassword == "" {
			c.JSON(http.StatusBadRequest, response.Error{Error: response.ErrEmptyRequestFields.Error()})
			return
		}

		userID, err := uuid.Parse(c.GetString("user_id"))
		if err != nil {
			slog.Error("user id format is not correct at jwt", "error", err.Error())
			c.JSON(http.StatusBadRequest, response.Error{Error: "user token set incorrectly"})
			return
		}

		err = userService.ResetPassword(c, userID, req.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, "password reset successful")
	}
}
