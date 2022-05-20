package req

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Fullname string `json:"fullname"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserReq struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email"`
	Phone    string `json:"hp"`
	Password string `json:"password"`
}
