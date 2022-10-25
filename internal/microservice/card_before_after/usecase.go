package cardBeforeAfter

import (
	"autfinal/internal/models"
)

type Usecase interface {
	CreateCardBeforeAfter(card *models.CardBeforeAfter, schedule_id int) (*models.CardBeforeAfter, error)
	GetCardsBeforeAfter(int) ([]*models.CardBeforeAfter, error)
	GetCardBeforeAfter(int, int) (*models.CardBeforeAfter, error)
	UpdateCardBeforeAfter(*models.CardBeforeAfter, int, int) (*models.CardBeforeAfter, error)
	UpdateCardsBeforeAfterOrder([]*models.CardBeforeAfter, int) ([]*models.CardBeforeAfter, error)
	DeleteCardBeforeAfter(int, int) (error)
}