package card

import (
	"autfinal/internal/models"
)

type Repository interface {
	CreateCardDay(*models.CardDay, string, int) (*models.CardDay, error)
	GetCardsDay(int) (*models.CardsDay, error)
	GetCardDay(int, int) (*models.CardDay, error)
	UpdateCardDay(*models.CardDay, int, int) (*models.CardDay, error)
	UpdateCardsOrder(*models.CardsDay, int) error
}
