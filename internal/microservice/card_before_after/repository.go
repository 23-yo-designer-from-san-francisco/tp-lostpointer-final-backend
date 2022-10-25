package cardBeforeAfter

import (
	"autfinal/internal/models"
)

type Repository interface {
	CreateCardBeforeAfter(card *models.CardBeforeAfter, mentor_id int) (*models.CardBeforeAfter, error)
	GetCardsBeforeAfter(int) ([]*models.CardBeforeAfter, error)
	GetCardBeforeAfter(int, int) (*models.CardBeforeAfter, error)
	UpdateCardBeforeAfter(*models.CardBeforeAfter, int, int) (*models.CardBeforeAfter, error)
	UpdateCardsBeforeAfterOrder([]*models.CardBeforeAfter, int) error
	DeleteCardBeforeAfter(int, int) (error)
}