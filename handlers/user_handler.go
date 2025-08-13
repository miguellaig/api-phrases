package handlers

import (
	"api-alemao/database"
	"api-alemao/dto/requests"
	"api-alemao/dto/responses"
	"api-alemao/models"
	"api-alemao/services"
	"api-alemao/utils"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	services *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{services: s}
}


func (h *UserHandler) RegisterUser(c echo.Context) error {
	var user requests.UserRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Falha ao lero json (register user)"})
	}

	_, err := h.services.BuscarUsuarioPorEmail(user.Email)
	if err == nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Email já cadastrado (register user)"})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Falha ao verificar email (register user)"})
	}

	hash, err := utils.GenerateHash(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Falha ao gerar hash (register user)"})
	}

	user.Password = hash

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hash,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Falha ao registrar usuário (register user)"})
	}

	result := responses.UserResponse{
		UserID: newUser.ID,
		Name:   newUser.Name,
		Email:  newUser.Email,
	}

	return c.JSON(http.StatusCreated, responses.APIResponse{Data: result})
}

func (h *UserHandler) LoginUserHandler(c echo.Context) error {
	var user requests.UserRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Falha ao ler o json (login user)"})
	}

	userFromDB, err := h.services.BuscarUsuarioPorEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Email ou senhas incorretas (login user)"})
	}

	if err := utils.CompareHashAndPassoword(userFromDB.Password, user.Password); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Email ou senha incorretos (login user)"})
	}

	token, err := utils.GenerateToken(userFromDB.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Falha ao gerar token (login user)"})
	}

	return c.JSON(http.StatusOK, responses.TokenResponse{Token: token})
}
