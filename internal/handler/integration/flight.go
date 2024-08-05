package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/dilyara4949/flight-booking-api/internal/config"
	"github.com/dilyara4949/flight-booking-api/internal/database/postgres"
	"github.com/dilyara4949/flight-booking-api/internal/handler"
	"github.com/dilyara4949/flight-booking-api/internal/handler/request"
	"github.com/dilyara4949/flight-booking-api/internal/repository"
	"github.com/dilyara4949/flight-booking-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func initDB(cfg config.Config) (*gorm.DB, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := postgres.Connect(ctx, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %v", err)
	}
	return database, nil
}

func TestUpdateFlightHandler(t *testing.T) {
	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("error getting config: %v", err)
	}

	database, err := initDB(cfg)
	if err != nil {
		t.Fatalf("init database failed: %v", err)
	}

	flightRepo := repository.NewFlightRepository(database)
	flightService := service.NewFlightService(flightRepo)

	tests := map[string]struct {
		flightID        string
		updateReq       request.Flight
		expectedStatus  int
		expectedMessage string
		expectedFlight  domain.Flight
	}{
		"OK flight update": {
			flightID: "e01b59c6-3eeb-4c79-b69f-24f0aa9d3516",
			updateReq: request.Flight{
				Destination: "Almaty",
			},
			expectedStatus:  http.StatusOK,
			expectedMessage: "",
			expectedFlight: domain.Flight{
				ID: uuid.MustParse("5a57c98d-87a0-436b-a016-634622efbf4e"),
				// Set other expected fields
			},
		},
		"invalid flight ID format": {
			flightID:  "invalid-id",
			updateReq: request.Flight{
				// Set fields to update
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "id format is not correct",
		},
		"non-existent flight": {
			flightID:  "5a57c98d-87a0-436b-436b-034622efbf4e",
			updateReq: request.Flight{
				// Set fields to update
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "flight not found",
		},
		"binding error": {
			flightID:  "5a57c98d-87a0-436b-a016-634622efbf4e",
			updateReq: request.Flight{
				// Leave fields empty or invalid to trigger binding error
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "error at binding request body",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.New()

			router.Use(func(c *gin.Context) {
				c.Set("flight_id", tt.flightID)
				c.Next()
			})

			router.PUT("/flights/:flightId", handler.UpdateFlightHandler(flightService))

			body, _ := json.Marshal(tt.updateReq)
			req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/flights/%s", tt.flightID), bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("flight_id", tt.flightID)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, w.Code)
			}

			if tt.expectedMessage != "" {
				var resp map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				if err != nil {
					t.Fatalf("error unmarshaling response: %v", err)
				}
				if resp["error"] != tt.expectedMessage {
					t.Errorf("expected message %v, got %v", tt.expectedMessage, resp["error"])
				}
			} else if !reflect.DeepEqual(tt.expectedFlight, domain.Flight{}) {
				var resp domain.Flight
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				if err != nil {
					t.Fatalf("error unmarshaling response: %v", err)
				}
				if resp.ID != tt.expectedFlight.ID {
					t.Errorf("expected flight %v, got %v", tt.expectedFlight, resp)
				}
			}
		})
	}
}
