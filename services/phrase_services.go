package services

import (
	"api-alemao/dto/requests"
	"api-alemao/models"

	"gorm.io/gorm"
)

type PhraseService struct {
	db *gorm.DB
}

func NewPhraseService(db *gorm.DB) *PhraseService {
	return &PhraseService{db: db}
}

func (p *PhraseService) CreatePhrase(phrase requests.PhrasesRequest) (models.Phrases, error) {
	newPhrase := models.Phrases{
		Original:    phrase.Original,
		Translation: phrase.Translation,
		Language:    phrase.Language,
	}

	err := p.db.Create(&newPhrase).Error
	return newPhrase, err
}

func (p *PhraseService) BuscarPhraseExistente(userid uint, phrase requests.PhrasesRequest) (models.Phrases, error) {

	var existingPhrase models.Phrases

	err := p.db.Where("original = ? AND user_id = ?", phrase.Original, userid).First(&existingPhrase).Error
	return existingPhrase, err
}

func (p *PhraseService) FiltrarPhraseWithOptionalQuerys(userid uint, query, lang string) ([]models.Phrases, error) {

	var phrases []models.Phrases

	db := p.db.Where("user_id = ?", userid)

	if lang != "" {
		db = db.Where("language = ?", lang)
	}
	if query != "" {
		db = db.Where("original ILIKE ?", "%"+query+"%")
	}

	err := db.Find(&phrases).Error

	return phrases, err

}

func (p *PhraseService) FindPhraseByUserIDAndIDparam(userid uint, idparam int) (models.Phrases, error) {
	var existingPhrase models.Phrases

	err := p.db.Where("user_id = ? AND id = ?", userid, idparam).First(&existingPhrase).Error

	return existingPhrase, err

}

func (p *PhraseService) UpdatePhrase(phrase models.Phrases, updatePhrase requests.PhrasesRequest) (models.Phrases, error) {

	err := p.db.Model(&phrase).Updates(models.Phrases{
		Original:    updatePhrase.Original,
		Translation: updatePhrase.Translation,
		Language:    updatePhrase.Language,
	}).Error

	return phrase, err

}

func (p *PhraseService) DeletePhrase(phrase models.Phrases) error {

	return p.db.Delete(&phrase).Error
}
