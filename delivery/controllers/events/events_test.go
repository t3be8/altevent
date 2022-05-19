package event

import (
	"altevent/delivery/middlewares"
	"altevent/entity"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	token string
)

func TestUseTokenizer(t *testing.T) {
	t.Run("Set Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(99, "username", "user99@test.com")
	})
}

func TestGetEventById(t *testing.T) {
	t.Run("Success get event data", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/event/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("99")

		eventControl := New(&mockEventRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(eventControl.GetEventById())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "Success get data by ID", resp.Message)
		assert.True(t, resp.Status)
		assert.NotNil(t, resp.Data)
	})

	t.Run("Error data not found", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/event/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("2")

		event := New(&errorMockEventRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(event.GetEventById())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Data not found", resp.Message)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)

	})
}

type Response struct {
	Code    int
	Message string
	Status  bool
	Data    interface{}
}

// Dummies
type mockEventRepo struct{}

func (mer *mockEventRepo) Create(newEvent entity.Event) (entity.Event, error) {
	return entity.Event{Title: "Uefa champions", Description: "final"}, nil
}
func (mer *mockEventRepo) GetEvent() ([]entity.Event, error) {
	return []entity.Event{{Title: "Nobar EUFA Champions", Description: "Nobar final liga champions"}}, nil
}

func (mer *mockEventRepo) GetEventID(id uint) (entity.Event, error) {
	return entity.Event{Title: "Nobar EUFA Champions", Description: "Nobar finaa liga champions"}, nil
}

func (mer *mockEventRepo) UpdateEvent(id uint, update entity.Event) (entity.Event, error) {
	return entity.Event{Title: "Uefa champions", Description: "final"}, nil
}
func (mer *mockEventRepo) DeleteEvent(id uint) (entity.Event, error) {
	return entity.Event{Title: "Uefa champions", Description: "final"}, nil
}

type errorMockEventRepo struct{}

func (emur *errorMockEventRepo) Create(newEvent entity.Event) (entity.Event, error) {
	return entity.Event{}, errors.New("error while accessing data")
}
func (emur *errorMockEventRepo) GetEvent() ([]entity.Event, error) {
	return []entity.Event{}, errors.New("error while accessing data")
}
func (emur *errorMockEventRepo) GetEventID(id uint) (entity.Event, error) {
	return entity.Event{}, errors.New("error while accessing data")
}
func (emur *errorMockEventRepo) UpdateEvent(id uint, update entity.Event) (entity.Event, error) {
	return entity.Event{}, errors.New("error while accessing data")
}
func (emur *errorMockEventRepo) DeleteEvent(id uint) (entity.Event, error) {
	return entity.Event{}, errors.New("error while accessing data")
}