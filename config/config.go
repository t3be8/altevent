package config

import (
	"os"
	"strconv"

	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port         int16
	DBPort       int16
	Host         string
	User         string
	Password     string
	DBName       string
	AWSKeyId     string
	AWSAccessKey string
	AWSRegion    string
}

func InitConfig() *AppConfig {
	app := GetConfig()
	if app == nil {
		log.Fatal("Cannot init config")
		return nil
	}
	return app
}

func GetConfig() *AppConfig {
	var res AppConfig

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal("Cannot open config file")
	// 	return nil
	// }

	portconv, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Warn(err)
	}
	res.Port = int16(portconv)
	conv, _ := strconv.Atoi(os.Getenv("DBPORT"))
	res.DBPort = int16(conv)
	res.Host = os.Getenv("DBHOST")
	res.User = os.Getenv("DBUSER")
	res.Password = os.Getenv("DBPASSWORD")
	res.DBName = os.Getenv("DBNAME")
	res.AWSKeyId = os.Getenv("AWS_KEY_ID")
	res.AWSAccessKey = os.Getenv("AWS_ACCESS_KEY")
	res.AWSRegion = os.Getenv("AWS_REGION")

	return &res
}
