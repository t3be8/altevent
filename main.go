package main

import (
	"altevent/config"
	commentController "altevent/delivery/controllers/comment"
	eventController "altevent/delivery/controllers/events"
	userController "altevent/delivery/controllers/user"
	"altevent/delivery/routes"
	CommentRepo "altevent/repository/comment"
	eventRepo "altevent/repository/events"
	userRepo "altevent/repository/user"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	config.AutoMigrate(db)

	e := echo.New()

	repoUser := userRepo.New(db)
	repoEvents := eventRepo.New(db)
	CommentRepo := CommentRepo.New(db)

	controllerUser := userController.New(repoUser, validator.New())
	controllerEvent := eventController.New(repoEvents, validator.New())
	controllerComment := commentController.New(CommentRepo, validator.New())

	routes.RegisterPath(e, controllerUser, controllerEvent, controllerComment)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
