package card

import (
	"autfinal/internal/models"
)

type Usecase interface {
	CreateCardDay(card *models.CardDay, schedule_id int) (*models.CardDay, error)
	GetCardsDay(int) ([]*models.CardDay, error)
	GetCardDay(int, int) (*models.CardDay, error)
	UpdateCardDay(*models.CardDay, int, int) (*models.CardDay, error)
	UpdateCardsDayOrder([]*models.CardDay, int) ([]*models.CardDay, error)
	DeleteCardDay(int, int) (error)

	CreateCardLesson(card *models.CardLesson, schedule_id int) (*models.CardLesson, error)
	GetCardsLesson(int) ([]*models.CardLesson, error)
	GetCardLesson(int, int) (*models.CardLesson, error)
	UpdateCardLesson(*models.CardLesson, int, int) (*models.CardLesson, error)
	UpdateCardsLessonOrder([]*models.CardLesson, int) ([]*models.CardLesson, error)
	DeleteCardLesson(int, int) (error)
}
