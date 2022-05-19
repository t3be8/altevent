package main

import (
	"altevent/config"
	eventController "altevent/delivery/controllers/events"
	userController "altevent/delivery/controllers/user"
	"altevent/delivery/routes"
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

	controllerUser := userController.New(repoUser, validator.New())
	controllerEvent := eventController.New(repoEvents, validator.New())

	routes.RegisterPath(e, controllerUser, controllerEvent)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
