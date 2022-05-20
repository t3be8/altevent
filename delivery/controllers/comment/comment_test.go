package comment

import (
	"altevent/delivery/middlewares"
	"altevent/entity"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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
		token, _ = middlewares.CreateToken(99, "Nobar EUFA Champions", "Nobar final liga champions")
	})
}

func TestPostComment(t *testing.T) {
	t.Run("Success posting comment", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"comment": "Pertamax",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events/:id/comments")
		context.SetParamNames("id")
		context.SetParamValues("1")

		commentController := New(&mockCommentRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(commentController.PostComment())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Post Comment!", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})

	t.Run("Error Bind Data", func(t *testing.T) {
		e := echo.New()
		// rb, _ := json.Marshal(map[string]interface{}{
		// 	"comment": "Pertamax",
		// })
		rb := "wrong bind"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events/:id/comments")
		context.SetParamNames("id")
		context.SetParamValues("1")

		commentController := New(&errorMockCommentRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(commentController.PostComment())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error at validation", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events/:id/comments")
		context.SetParamNames("id")
		context.SetParamValues("1")

		commentController := New(&errorMockCommentRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(commentController.PostComment())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Validate Error", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error DB Insert Data", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"comment": "Pertamax",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events/:id/comments")
		context.SetParamNames("id")
		context.SetParamValues("1")

		commentController := New(&errorMockCommentRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(commentController.PostComment())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})
}

func TestSelectAllInEvent(t *testing.T) {
	t.Run("Success Select All in event", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events/:id/comments")
		context.SetParamNames("id")
		context.SetParamValues("99")

		comment := New(&mockCommentRepo{}, validator.New())
		comment.SelectAllInEvent()(context)

		// type response struct {
		// 	Code    int
		// 	Message string
		// 	Status  bool
		// 	Data    []entity.Event
		// }

		var resp []entity.Comment

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, resp[0].Comment, "Pertamax!")
	})
	t.Run("Error select all in event", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events/:id/comments")
		context.SetParamNames("id")
		context.SetParamValues("2")

		commentController := New(&errorMockCommentRepo{}, validator.New())
		commentController.SelectAllInEvent()

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

func TestUpdateComment(t *testing.T) {
	t.Run("Success Update Data", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Comment": "Pertalite Diamankan",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/comments/:id")
		context.SetParamNames("id")
		context.SetParamValues("99")
		eventController := New(&mockCommentRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(eventController.Update())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Updated", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})

	t.Run("Error data not found", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/comments/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("99")

		event := New(&errorMockCommentRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(event.Update())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Data not found", resp.Message)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)

	})
}

func TestDeleteComment(t *testing.T) {
	t.Run("Success Delete Data", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/comments/:id")
		context.SetParamNames("id")
		context.SetParamValues("99")
		eventController := New(&mockCommentRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(eventController.Delete())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Deleted", result.Message)
		assert.True(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error data not found", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/comments/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("99")

		event := New(&errorMockCommentRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(event.Delete())(ctx)

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

type mockCommentRepo struct{}

func (mcr *mockCommentRepo) CreateComment(comment entity.Comment) (entity.Comment, error) {
	return entity.Comment{Comment: "Pertamax!"}, nil
}
func (mcr *mockCommentRepo) SelectAllComment(evid uint) ([]entity.Comment, error) {
	return []entity.Comment{{Comment: "Pertamax!"}}, nil
}
func (mcr *mockCommentRepo) UpdateComment(id, user_id uint, UpdateComment entity.Comment) (entity.Comment, error) {
	return entity.Comment{Comment: "Pertamax!"}, nil
}
func (mcr *mockCommentRepo) DeleteComment(id, user_id uint) error {
	return nil
}

type errorMockCommentRepo struct{}

func (emcr *errorMockCommentRepo) CreateComment(comment entity.Comment) (entity.Comment, error) {
	return entity.Comment{}, errors.New("data access error")
}
func (emcr *errorMockCommentRepo) SelectAllComment(evid uint) ([]entity.Comment, error) {
	return []entity.Comment{}, errors.New("data access error")
}
func (emcr *errorMockCommentRepo) UpdateComment(id, user_id uint, UpdateComment entity.Comment) (entity.Comment, error) {
	return entity.Comment{}, errors.New("data access error")
}
func (emcr *errorMockCommentRepo) DeleteComment(id, user_id uint) error {
	return errors.New("data access error")
}
