package transformers

import (
	"api-alemao/dto/responses"
	"api-alemao/models"
)

func ListPhraseResponse(phrases []models.Phrases) []responses.PhraseResponse {

	var result []responses.PhraseResponse

	for _, frase := range phrases {
		result = append(result, responses.PhraseResponse{
			ID:          frase.ID,
			UserID:      frase.UserID,
			Original:    frase.Original,
			Translation: frase.Translation,
			Language:    frase.Language,
		})
	}
	return result
}

func UpdatedPhrase(phrase models.Phrases) responses.PhraseResponse {
	result := responses.PhraseResponse{
		UserID:      phrase.UserID,
		ID:          phrase.ID,
		Original:    phrase.Original,
		Translation: phrase.Translation,
		Language:    phrase.Language,
	}

	return result
}
