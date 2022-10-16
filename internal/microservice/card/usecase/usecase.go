package usecase

import (
	"autfinal/internal/microservice/card"
	"autfinal/internal/microservice/schedule"
	"autfinal/internal/models"

	log "autfinal/pkg/logger"
)

type CardUsecase struct {
	scheduleRepo schedule.Repository
	cardRepo card.Repository
}

func NewCardUsecase(scheduleR schedule.Repository, cardR card.Repository) *CardUsecase {
	return &CardUsecase{
		scheduleRepo: scheduleR,
		cardRepo: cardR,
	}
}

func (cU *CardUsecase) CreateCardDay(Card *models.CardDay, imgUrl string, schedule_id int) (*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	return cU.cardRepo.CreateCardDay(Card, imgUrl, schedule_id)
}

func (cU *CardUsecase) GetCardsDay(schedule_id int) (*[]models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	return cU.cardRepo.GetCardsDay(schedule_id)
}

func (cU *CardUsecase) GetCardDay(schedule_id, card_id int) (*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	return cU.cardRepo.GetCardDay(schedule_id, card_id)
}

func (cU *CardUsecase) UpdateCardDay(card *models.CardDay, schedule_id, card_id int) (*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	return cU.cardRepo.UpdateCardDay(card, schedule_id, card_id)
}

func (cU *CardUsecase) UpdateCardsOrder(cards *models.CardsDay, schedule_id int) (*models.CardsDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie
	return cU.cardRepo.UpdateCardsOrder(cards, schedule_id)
}