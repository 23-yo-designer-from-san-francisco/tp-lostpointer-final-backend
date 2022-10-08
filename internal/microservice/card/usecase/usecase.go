package usecase

import (
	"autfinal/internal/microservice/card"
	"autfinal/internal/models"
)

type CardUsecase struct {
	cardRepo card.Repository
}

func NewCardUsecase(cardR card.Repository) *CardUsecase {
	return &CardUsecase{
		cardRepo: cardR,
	}
}

func(cU *CardUsecase) CreateCard(card *models.Card) (*models.Card, error) {
	return cU.cardRepo.CreateCard(card)
}

func(cU *CardUsecase) GetCards() (*[]models.Card, error) {
	return cU.cardRepo.GetCards()
}