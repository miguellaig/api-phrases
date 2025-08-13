package routes

import (
	"api-alemao/handlers"
	"api-alemao/middleware"

	_ "api-alemao/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Requests(e *echo.Echo, userhandler *handlers.UserHandler, phrasehandler *handlers.PhraseHandler) error {

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/register", userhandler.RegisterUser)
	e.POST("/login", userhandler.LoginUserHandler)

	r := e.Group("/login")
	r.Use(middleware.ValidationMiddleware)
	r.POST("", phrasehandler.RegisterPhraseHandler)
	r.GET("", phrasehandler.ListPhrasesHandler)
	r.PUT("/:id", phrasehandler.UpdatePhrasesHandler)
	r.DELETE("/:id", phrasehandler.DeletePhraseHandler)

	return nil
}
