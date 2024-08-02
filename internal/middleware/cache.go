package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/dilyara4949/flight-booking-api/internal/handler/auth"
	"github.com/dilyara4949/flight-booking-api/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

func Cache(cache *redis.Client, ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !(c.Request.Method == http.MethodGet || c.Request.Method == http.MethodPut) {
			c.Next()
			return
		}

		var id string
		switch {
		case strings.Contains(c.Request.URL.Path, "/flights/"):
			flightID := c.Param("flightId")
			if flightID == "" {
				c.Next()
				return
			}
			id = "flight-" + flightID
		case strings.Contains(c.Request.URL.Path, "/users/"):
			userID := c.Param("userId")
			if userID == "" {
				c.Next()
				return
			}

			if !auth.AccessCheck(*c, c.GetString("user_id"), "userId") {
				c.AbortWithStatusJSON(http.StatusForbidden, response.Error{Error: "access deniedo"})
				return
			}

			id = "user-" + userID
		default:
			c.Next()
			return
		}

		if c.Request.Method == http.MethodGet {
			res, err := cache.Get(c, id).Result()
			if err == nil {
				slog.Info("Cache hit for key", "id", id)
				var jsonResponse json.RawMessage
				err = json.Unmarshal([]byte(res), &jsonResponse)
				if err != nil {
					slog.Error("Failed to unmarshal cached response", "error", err)
					c.Next()
					return
				}
				c.AbortWithStatusJSON(http.StatusOK, jsonResponse)
				return
			}
		} else {
			slog.Info("Cache invalidation on update", "id", id)
		}

		slog.Info("Cache miss for key", "id", id)

		writer := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		c.Next()

		if c.Writer.Status() == http.StatusOK {
			cache.Set(c, id, writer.body.String(), ttl)
		}
	}
}

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
