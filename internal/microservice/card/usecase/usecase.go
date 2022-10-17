package usecase

import (
	"autfinal/internal/microservice/card"
	"autfinal/internal/microservice/schedule"
	"autfinal/internal/models"
	"autfinal/internal/utils"

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

func (cU *CardUsecase) CreateCardDay(card *models.CardDay, imgUUID string, schedule_id int) (*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCard, err := cU.cardRepo.CreateCardDay(card, imgUUID, schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cU *CardUsecase) GetCardsDay(schedule_id int) (*models.CardsDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCards, err := cU.cardRepo.GetCardsDay(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, resultCard := range resultCards.Cards{
		resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	}

	return resultCards, nil
}

func (cU *CardUsecase) GetCardDay(schedule_id, card_id int) (*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCard, err := cU.cardRepo.GetCardDay(schedule_id, card_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cU *CardUsecase) UpdateCardDay(card *models.CardDay, schedule_id, card_id int) (*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCard, err := cU.cardRepo.UpdateCardDay(card, schedule_id, card_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cU *CardUsecase) UpdateCardsOrder(cards *models.CardsDay, schedule_id int) (*models.CardsDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie
	err = cU.cardRepo.UpdateCardsOrder(cards, schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCards, err := cU.cardRepo.GetCardsDay(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, resultCard := range resultCards.Cards{
		resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	}

	return resultCards, nil
}