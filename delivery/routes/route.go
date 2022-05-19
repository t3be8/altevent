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

	// Users route
	apiGroup.GET("/users/:id", uc.Show(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))
	apiGroup.PUT("/users/:id", uc.Update(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))
	apiGroup.DELETE("/users/:id", uc.Delete(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))
	apiGroup.GET("/users/:id/events", uc.ShowMyEvent(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))

	// Events route
	apiGroup.POST("/events", ec.InsertEvent())
	// apiGroup.POST("/events/:id/join", ec.InsertEvent())
	apiGroup.GET("/events", ec.SelectEvent())
	// apiGroup.GET("/events?title", ec.SelectEvent())
	// apiGroup.GET("/events?byme", ec.SelectEvent())
	apiGroup.GET("/events/:id", ec.GetEventById(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))
	apiGroup.PUT("/events/:id", ec.UpdateEvent(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))
	apiGroup.DELETE("/events/:id", ec.DeleteEvent(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))

	// Order ticket route
	// apiGroup.POST("/orders", oc.CreateOrder(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))
	// apiGroup.POST("/orders/{order_id}/cancel", oc.CancelOrder(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))
	// apiGroup.POST("/orders/{order_id}/payout", oc.Payment(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))

}
