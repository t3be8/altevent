package event

import "github.com/labstack/echo/v4"

type IEventController interface {
	InsertEvent() echo.HandlerFunc
	SelectEvent() echo.HandlerFunc
	DeleteEvent() echo.HandlerFunc
	UpdateEvent() echo.HandlerFunc
	GetEventById() echo.HandlerFunc
}
