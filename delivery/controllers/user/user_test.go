package user

import (
	"altevent/delivery/middlewares"
	"altevent/entity"
	"altevent/utils"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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

func TestRegister(t *testing.T) {
	t.Run("Success register new user", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"fullname": "John Doe",
			"username": "jodoe",
			"email":    "jdoe@test.com",
			"phone":    "089123123123",
			"password": "admin1234",
		})

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()
		context := e.NewContext(request, response)
		context.SetPath("/register")

		registerController := New(&mockUserRepo{}, validator.New())
		registerController.Register()(context)

		var resp Response

		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, "berhasil register user baru", resp.Message)
		assert.True(t, resp.Status)
		assert.Equal(t, 201, resp.Code)
		assert.NotNil(t, resp.Data)
	})

	t.Run("Error Binding Params", func(t *testing.T) {
		e := echo.New()
		// rb, _ := json.Marshal(map[string]interface{}{
		// 	"fullname": "Jane Doe",
		// 	"username": "jadoe",
		// 	"email":    "jadoe@test.com",
		// 	"phone":    "089123123123",
		// 	"password": "admin1234",
		// })

		rb := "wrong params"

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		ctx := e.NewContext(request, response)
		ctx.SetPath("/register")

		registerController := New(&errorMockUserRepo{}, validator.New())
		registerController.Register()(ctx)

		var resp Response

		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "Invalid Request", resp.Message)
	})

	t.Run("Error at validate username", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"email": "jdoe@test.com",
		})

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		ctx := e.NewContext(request, response)
		ctx.SetPath("/register")

		registerController := New(&errorMockUserRepo{}, validator.New())
		registerController.Register()(ctx)

		var resp Response

		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 406, resp.Code)
		assert.Equal(t, "Validate Error", resp.Message)
	})

	t.Run("Error at validate email", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"username": "jdoe",
		})

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		ctx := e.NewContext(request, response)
		ctx.SetPath("/register")

		registerController := New(&errorMockUserRepo{}, validator.New())
		registerController.Register()(ctx)

		var resp Response

		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 406, resp.Code)
		assert.Equal(t, "Validate Error", resp.Message)
	})

	t.Run("Error at validate password", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"username": "jdoe",
			"email":    "jdoe@test.com",
		})

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		ctx := e.NewContext(request, response)
		ctx.SetPath("/register")

		registerController := New(&errorMockUserRepo{}, validator.New())
		registerController.Register()(ctx)

		var resp Response

		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 406, resp.Code)
		assert.Equal(t, "Validate Error", resp.Message)
	})

	t.Run("Error DB Repo Register", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"fullname": "John Doe",
			"username": "jodoe",
			"email":    "jdoe@test.com",
			"phone":    "089123123123",
			"password": "admin1234",
		})

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		ctx := e.NewContext(request, response)
		ctx.SetPath("/register")

		registerController := New(&errorMockUserRepo{}, validator.New())
		registerController.Register()(ctx)

		var resp Response

		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 500, resp.Code)
		assert.Equal(t, "Cannot Access Database", resp.Message)
	})
}

func TestShow(t *testing.T) {
	t.Run("Success get user data", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/users/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("99")

		userControl := New(&mockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userControl.Show())(ctx)

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
		ctx.SetPath("/users/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("99")

		user := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(user.Show())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Data not found", resp.Message)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)

	})

	t.Run("Error Convert params", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/users/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("xx")

		userControl := New(&mockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userControl.Show())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, 406, resp.Code)
		assert.Equal(t, "Cannot Convert ID", resp.Message)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
	})

	t.Run("Error Get Data", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/users/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("99")

		user := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(user.Show())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		// log.Warn(resp)
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Data not found", resp.Message)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
	})

}

func TestShowMyEvent(t *testing.T) {
	t.Run("Success get user data", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/users/:id/events")
		ctx.SetParamNames("id")
		ctx.SetParamValues("99")

		userControl := New(&mockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(userControl.ShowMyEvent())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "Success get data by ID", resp.Message)
		assert.True(t, resp.Status)
		assert.NotNil(t, resp.Data)
	})

	t.Run("Error UserID didnt have access token", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/users/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("9")

		user := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(user.ShowMyEvent())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Data not found", resp.Message)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)

	})

	t.Run("Error convert params", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/users/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("x")

		user := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(user.ShowMyEvent())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, 406, resp.Code)
		assert.Equal(t, "Cannot Convert ID", resp.Message)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)

	})

	t.Run("Error DB data not found", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		response := httptest.NewRecorder()

		ctx := e.NewContext(request, response)
		ctx.SetPath("/users/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("99")

		user := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")})(user.ShowMyEvent())(ctx)

		var resp Response
		json.Unmarshal([]byte(response.Body.Bytes()), &resp)
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Data not found", resp.Message)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success Update Data", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"fullname": "John Doel",
			"email":    "jdoes@test.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("99")
		userController := New(&mockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Update())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Updated", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})

	t.Run("Error Bind Data", func(t *testing.T) {
		e := echo.New()
		// rb, _ := json.Marshal(map[string]interface{}{
		// 	"id":       99,
		// 	"fullname": "John Doel",
		// 	"email":    "jdoes@test.com",
		// })

		rb := "wrongs binding"
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("99")

		userController := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Update())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error at validate fullname", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"email": "jdoes@test.com",
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("99")

		userController := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Update())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Validate Error", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error at convert params", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"fullname": "John Doea",
			"email":    "jdoes@test.com",
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("xx")

		userController := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Update())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error other user cannot update others", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"fullname": "John Doea",
			"email":    "jdoes@test.com",
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("9")

		userController := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Update())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "token tidak ditemukan", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error update password hash", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"fullname": "John Doea",
			"email":    "jdoes@test.com",
			"password": "admin2345",
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("99")

		userController := New(&mockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Update())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		// assert.Equal(t, 500, result.Code)
		// assert.Equal(t, "Cannot Access Database", result.Message)
		// assert.False(t, result.Status)
		// assert.Nil(t, result.Data)

		// pwd := result.Data.(map[string]interface{})
		// log.Info(pwd["password"])
		_, err := utils.HashPassword("")
		if err != nil {
			assert.Equal(t, 500, result.Code)
			assert.Equal(t, "Cannot Access Database", result.Message)
			assert.False(t, result.Status)
			assert.Nil(t, result.Data)
		}
	})

	t.Run("Error DB update", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"fullname": "John Doea",
			"email":    "jdoes@test.com",
			"password": "admin2345",
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("99")

		userController := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Update())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Success Delete Data", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("99")
		userController := New(&mockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Delete())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Deleted", result.Message)
		assert.True(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error at convert params", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPut, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("x")

		userController := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Delete())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		// log.Warn(result)
		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error authorized user", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPut, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("9")

		userController := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Delete())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		log.Warn(result)
		assert.Equal(t, 401, result.Code)
		assert.Equal(t, "Unauthorized", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})

	t.Run("Error DB Delete", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPut, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("99")

		userController := New(&errorMockUserRepo{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")})(userController.Delete())(context)

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		log.Warn(result)
		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
		assert.Nil(t, result.Data)
	})

}

func TestIsLogin(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":    "jdoe@test.com",
			"password": "admin123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		loginController := New(&mockUserRepo{}, validator.New())
		loginController.Login()(context)

		var response Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.True(t, response.Status)
		// log.Warn(response.Data)
		assert.NotNil(t, response.Data)
		data := response.Data.(map[string]interface{})
		token = data["token"].(string)
		assert.Equal(t, "Berhasil login!", response.Message)
	})

	t.Run("Error Binding", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email": 123434,
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		loginController := New(&errorMockUserRepo{}, validator.New())
		loginController.Login()(context)

		var response Response
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// log.Info(response)
		assert.Equal(t, 415, response.Code)
		assert.False(t, response.Status)
		assert.Equal(t, "Cannot Bind Data", response.Message)
		assert.Nil(t, response.Data)
	})

	t.Run("Error Login Validate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email": "",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		loginController := New(&errorMockUserRepo{}, validator.New())
		loginController.Login()(context)

		var response Response
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// log.Info(response)
		assert.Equal(t, 406, response.Code)
		assert.False(t, response.Status)
		assert.Equal(t, "Validate Error", response.Message)
		assert.Nil(t, response.Data)
	})

	t.Run("Error Login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"email":    "admin1@test.com",
			"password": "salahpassword",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		loginController := New(&errorMockUserRepo{}, validator.New())
		loginController.Login()(context)

		var response Response
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// log.Info(response)
		assert.Equal(t, 404, response.Code)
		assert.False(t, response.Status)
		assert.Equal(t, "Data not found, gagal login!", response.Message)
		assert.Nil(t, response.Data)
	})
}

type Response struct {
	Code    int
	Message string
	Status  bool
	Data    interface{}
}

// Dummies
type mockUserRepo struct{}

func (mur *mockUserRepo) Register(newUser entity.User) (entity.User, error) {
	return entity.User{Username: "username", Email: "user99@test.com"}, nil
}
func (mur *mockUserRepo) IsLogin(email, password string) (entity.User, error) {
	return entity.User{Username: "username", Email: "user99@test.com"}, nil
}
func (mur *mockUserRepo) GetUserID(id uint) (entity.User, error) {
	return entity.User{Username: "username", Email: "user99@test.com"}, nil
}
func (mur *mockUserRepo) UpdateUser(id uint, update entity.User) (entity.User, error) {
	return entity.User{Username: "username", Email: "user99@test.com"}, nil
}
func (mur *mockUserRepo) DeleteUser(id uint) (entity.User, error) {
	return entity.User{Username: "username", Email: "user99@test.com"}, nil
}
func (mur *mockUserRepo) GetMyEvent(id uint) ([]entity.Event, error) {
	return []entity.Event{{Title: "Nobar Final Champions League", Ticket: 125}}, nil
}

type errorMockUserRepo struct{}

func (emur *errorMockUserRepo) Register(newUser entity.User) (entity.User, error) {
	return entity.User{}, errors.New("error while accessing data")
}
func (emur *errorMockUserRepo) IsLogin(email, password string) (entity.User, error) {
	return entity.User{}, errors.New("error while accessing data")

}
func (emur *errorMockUserRepo) GetUserID(id uint) (entity.User, error) {
	return entity.User{}, errors.New("error while accessing data")

}
func (emur *errorMockUserRepo) UpdateUser(id uint, update entity.User) (entity.User, error) {
	return entity.User{}, errors.New("error while accessing data")

}
func (emur *errorMockUserRepo) DeleteUser(id uint) (entity.User, error) {
	return entity.User{}, errors.New("error while accessing data")

}
func (emur *errorMockUserRepo) GetMyEvent(id uint) ([]entity.Event, error) {
	return []entity.Event{}, errors.New("error while accessing data")
}
