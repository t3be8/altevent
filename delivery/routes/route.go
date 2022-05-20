package routes

import (
	"altevent/delivery/controllers/comment"
	event "altevent/delivery/controllers/events"
	"altevent/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc user.IUserController, ec event.IEventController, cc comment.ICommentController) {
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
	apiGroup.POST("/events", ec.InsertEvent(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))
	// apiGroup.POST("/events/:id/join", ec.InsertEvent())
	apiGroup.GET("/events", ec.SelectEvent())
	apiGroup.GET("/events?title=:title", ec.SearchEventContains())
	// apiGroup.GET("/events?byme", ec.SelectEvent())
	apiGroup.GET("/events/:id", ec.GetEventById(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))
	apiGroup.PUT("/events/:id", ec.UpdateEvent(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))
	apiGroup.DELETE("/events/:id", ec.DeleteEvent(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTEVEN")}))

	// Comment Routes
	apiGroup.POST("/events/:id/comments", cc.PostComment(), middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")}))
	apiGroup.GET("/events/:id/comments", cc.SelectAllInEvent())
	apiGroup.PUT("/comments/:id", cc.Update(), middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")}))
	apiGroup.DELETE("/comments/:id", cc.Delete(), middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("ALTEVEN")}))

}
