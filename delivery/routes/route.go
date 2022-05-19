package routes

import (
	event "altevent/delivery/controllers/events"
	"altevent/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc user.IUserController, ec event.IEventController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())

	apiGroup := e.Group("/api")

	// Auth route
	apiGroup.POST("/login", uc.Login())
	apiGroup.POST("/register", uc.Register())

	// Order ticket route
	// apiGroup.POST("/orders", oc.CreateOrder(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))
	// apiGroup.POST("/orders/{order_id}/cancel", oc.CancelOrder(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))
	// apiGroup.POST("/orders/{order_id}/payout", oc.Payment(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))

	// Event Route
	apiGroup.POST("/events", ec.InsertEvent())
	apiGroup.GET("/events", ec.SelectEvent())
	apiGroup.PUT("/events/{id}", ec.UpdateEvent())
	apiGroup.DELETE("/events/{id}", ec.DeleteEvent())
	apiGroup.GET("/events/{id}", ec.GetEventById())

}
