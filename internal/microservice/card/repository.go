package card

import (
	"autfinal/internal/models"
)

type Repository interface {
	CreateCard(card *models.Card) (*models.Card, error)
	GetCards() (*[]models.Card, error)
}
