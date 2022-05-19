package event

import (
	"altevent/delivery/middlewares"
	"altevent/entity"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// "github.com/labstack/gommon/log"

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

func TestSelectEvent(t *testing.T) {
	t.Run("Success Select Event", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events")

		event := New(&mockEventRepo{}, validator.New())
		event.SelectEvent()(context)

		// type response struct {
		// 	Code    int
		// 	Message string
		// 	Status  bool
		// 	Data    []entity.Event
		// }

		var resp []entity.Event

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, resp[0].Title, "Nobar EUFA Champions")
	})
	t.Run("Error select product", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events")

		eventController := New(&errorMockEventRepo{}, validator.New())
		eventController.SelectEvent()

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entity.Event
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		// assert.Nil(t, resp.Data)
		// assert.False(t, resp.Status)

	})
}

func TestInsertEvent(t *testing.T) {
	t.Run("Success Insert Event", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"title":       "Nobar EUFA Champions",
			"description": "Nobar final liga champions",
			"rules":       "no alcohol",
			"organizer":   "cafe 123",
			"due_date":    "29 Mei 2022",
			"begin_at":    "00:30",
			"location":    "jakarta",
			"ticket":      4,
			"links":       "bitly.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderAuthorization, "Bearer"+token)
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events")

		event := New(&mockEventRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(event.InsertEvent())(context)

		var resp Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		//assert.Equal(t, "Nobar final liga champions", resp.Data.(map[string]interface{})["description"])
		assert.True(t, resp.Status)
	})
	t.Run("Error at validate title", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"title": "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events")

		event := New(&errorMockEventRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(event.InsertEvent())(context)

		var resp Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.False(t, resp.Status)
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

func (mer *mockEventRepo) InsertEvent(newEvent entity.Event) (entity.Event, error) {
	return newEvent, nil
}
func (mer *mockEventRepo) SelectEvent() ([]entity.Event, error) {
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

func (emur *errorMockEventRepo) InsertEvent(newEvent entity.Event) (entity.Event, error) {
	return entity.Event{}, errors.New("error while accessing data")
}
func (emur *errorMockEventRepo) SelectEvent() ([]entity.Event, error) {
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
