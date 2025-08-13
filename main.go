package main

import (
	"api-alemao/database"
	"api-alemao/handlers"
	"api-alemao/routes"
	"api-alemao/services"
	"log"

	_ "api-alemao/docs"

	"github.com/labstack/echo/v4"
)

func main() {

	UserService := services.NewUserService(database.DB)
	UserHandler := handlers.NewUserHandler(UserService)
	PhraseService := services.NewPhraseService(database.DB)
	PhraseHandler := handlers.NewPhraseHandler(PhraseService)

	e := echo.New()

	database.Connect()

	routes.Requests(e, UserHandler, PhraseHandler)

	log.Fatal(e.Start(":8080"))
}
