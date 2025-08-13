package handlers

import (
	"api-alemao/dto/requests"
	"api-alemao/dto/responses"
	"api-alemao/services"
	"api-alemao/transformers"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PhraseHandler struct {
	service *services.PhraseService
}

func NewPhraseHandler(s *services.PhraseService) *PhraseHandler {
	return &PhraseHandler{service: s}
}

func (p *PhraseHandler) RegisterPhraseHandler(c echo.Context) error {
	userID := c.Get("user").(uint)

	var phrase requests.PhrasesRequest

	if err := c.Bind(&phrase); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Falha ao ler o json (register phrase)"})
	}

	if _, err := p.service.BuscarPhraseExistente(userID, phrase); err == nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Frase já cadastrada (register phrase)"})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Falha ao analisar frase (register phrase)"})
	}

	newPhrase, err := p.service.CreatePhrase(phrase)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Falha ao registrar frase (register phrase)"})
	}

	result := responses.PhraseResponse{
		ID:          newPhrase.ID,
		UserID:      newPhrase.UserID,
		Original:    newPhrase.Original,
		Translation: newPhrase.Translation,
		Language:    newPhrase.Language,
	}

	return c.JSON(http.StatusCreated, responses.APIResponse{Data: result})

}

func (p *PhraseHandler) ListPhrasesHandler(c echo.Context) error {
	userID := c.Get("user").(uint)

	lang := strings.TrimSpace(c.QueryParam("lang"))
	query := c.QueryParam("query")

	existingPhrases, err := p.service.FiltrarPhraseWithOptionalQuerys(userID, lang, query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Falha ao listar frases (list phrase)"})
	}
	if len(existingPhrases) == 0 {
		return c.JSON(http.StatusOK, responses.APIResponse{Data: []responses.PhraseResponse{}})
	}

	result := transformers.ListPhraseResponse(existingPhrases)

	return c.JSON(http.StatusOK, responses.APIResponse{Data: result})
}

func (p *PhraseHandler) UpdatePhrasesHandler(c echo.Context) error {
	userID := c.Get("user").(uint)

	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "ID inválido (update phrases)"})
	}

	var updatePhrase requests.PhrasesRequest

	if err := c.Bind(&updatePhrase); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Falha ao ler o json (update phrase)"})
	}

	existingPhrase, err := p.service.FindPhraseByUserIDAndIDparam(userID, idParam)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Frase não encontrada (update phrase)"})
		}
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Erro ao buscar frase (update phrase)"})
	}

	updatedPhrase, err := p.service.UpdatePhrase(existingPhrase, updatePhrase)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Falha ao tentar atualizar frase (update phrase)"})
	}

	result := transformers.UpdatedPhrase(updatedPhrase)

	return c.JSON(http.StatusOK, responses.APIResponse{Data: result})
}

func (p *PhraseHandler) DeletePhraseHandler(c echo.Context) error {
	userID := c.Get("user").(uint)

	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "ID inválido (delete phrase)"})
	}

	existingPhrase, err := p.service.FindPhraseByUserIDAndIDparam(userID, idParam)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Frase não encontrada (delete phrase)"})
		}
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Erro ao buscar frase (delete phrase)"})
	}

	result := p.service.DeletePhrase(existingPhrase)
	if result != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Erro ao tentar deletar frase (delete phrase)"})
	}

	return c.NoContent(http.StatusNoContent)
}
