package usecase

type Usecase interface {
	CreateSchedule(card *models.Card) (*models.Card, error)
}