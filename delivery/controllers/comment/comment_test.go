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
