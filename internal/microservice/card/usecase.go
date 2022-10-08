package card

import (
	"autfinal/internal/models"
)

type Usecase interface {
	CreateCard(card *models.Card) (*models.Card, error)
	GetCards() (*[]models.Card, error)
}