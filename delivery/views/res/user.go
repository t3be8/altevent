package res

import (
	"net/http"
)

type LoginResponse struct {
	Data  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func RegisterSuccess(data UserResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil register user baru",
		"status":  true,
		"data":    data,
	}
}
