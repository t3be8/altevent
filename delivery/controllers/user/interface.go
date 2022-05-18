package user

import "github.com/labstack/echo/v4"

type IUserController interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Show() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
