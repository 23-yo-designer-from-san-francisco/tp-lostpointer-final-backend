package card

import (
	"autfinal/internal/models"
)

type Repository interface {
	CreateCardDay(*models.CardDay, string, int) (*models.CardDay, error)
	GetCardsDay(int) ([]*models.CardDay, error)
	GetCardDay(int, int) (*models.CardDay, error)
	UpdateCardDay(*models.CardDay, int, int) (*models.CardDay, error)
	UpdateCardsDayOrder([]*models.CardDay, int) error
	DeleteCardDay(int, int) (error)

	CreateCardLesson(*models.CardLesson, string, int) (*models.CardLesson, error)
	GetCardsLesson(int) ([]*models.CardLesson, error)
	GetCardLesson(int, int) (*models.CardLesson, error)
	UpdateCardLesson(*models.CardLesson, int, int) (*models.CardLesson, error)
	UpdateCardsLessonOrder([]*models.CardLesson, int) error
	DeleteCardLesson(int, int) (error)
}
