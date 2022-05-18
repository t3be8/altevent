package req

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdateUserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"`
	Phone    string `json:"hp"`
	Password string `json:"password"`
}
