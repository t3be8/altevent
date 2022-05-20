package comment

import "github.com/labstack/echo/v4"

type ICommentController interface {
	PostComment() echo.HandlerFunc
	SelectAllInEvent() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
