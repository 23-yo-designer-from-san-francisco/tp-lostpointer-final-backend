package usecase

import (
	"autfinal/internal/microservice/card_before_after"
	"autfinal/internal/microservice/schedule_before_after"
	"autfinal/internal/models"
	"autfinal/internal/utils"

	log "autfinal/pkg/logger"
)

type CardBeforeAfterUsecase struct {
	scheduleRepo scheduleBeforeAfter.Repository
	cardRepo     cardBeforeAfter.Repository
}

func NewCardBeforeAfterUsecase(scheduleR scheduleBeforeAfter.Repository, cardR cardBeforeAfter.Repository) *CardBeforeAfterUsecase {
	return &CardBeforeAfterUsecase{
		scheduleRepo: scheduleR,
		cardRepo:     cardR,
	}
}

func (cbaU *CardBeforeAfterUsecase) CreateCardBeforeAfter(card *models.CardBeforeAfter, schedule_id int) (*models.CardBeforeAfter, error) {
	mentor_id, err := cbaU.scheduleRepo.GetMentorIdFromScheduleBeforeAfterID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie
	card.Schedule_ID = schedule_id
	resultCard, err := cbaU.cardRepo.CreateCardBeforeAfter(card, mentor_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cbaU *CardBeforeAfterUsecase) GetCardsBeforeAfter(schedule_id int) ([]*models.CardBeforeAfter, error) {
	_, err := cbaU.scheduleRepo.GetMentorIdFromScheduleBeforeAfterID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCards, err := cbaU.cardRepo.GetCardsBeforeAfter(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, resultCard := range resultCards {
		resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	}

	return resultCards, nil
}

func (cbaU *CardBeforeAfterUsecase) GetCardBeforeAfter(schedule_id, card_id int) (*models.CardBeforeAfter, error) {
	_, err := cbaU.scheduleRepo.GetMentorIdFromScheduleBeforeAfterID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCard, err := cbaU.cardRepo.GetCardBeforeAfter(schedule_id, card_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cbaU *CardBeforeAfterUsecase) UpdateCardBeforeAfter(card *models.CardBeforeAfter, schedule_id, card_id int) (*models.CardBeforeAfter, error) {
	_, err := cbaU.scheduleRepo.GetMentorIdFromScheduleBeforeAfterID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCard, err := cbaU.cardRepo.UpdateCardBeforeAfter(card, schedule_id, card_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cbaU *CardBeforeAfterUsecase) UpdateCardsBeforeAfterOrder(cards []*models.CardBeforeAfter, schedule_id int) ([]*models.CardBeforeAfter, error) {
	_, err := cbaU.scheduleRepo.GetMentorIdFromScheduleBeforeAfterID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie
	err = cbaU.cardRepo.UpdateCardsBeforeAfterOrder(cards, schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCards, err := cbaU.cardRepo.GetCardsBeforeAfter(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, resultCard := range resultCards {
		resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	}

	return resultCards, nil
}

func (cbaU *CardBeforeAfterUsecase) DeleteCardBeforeAfter(schedule_id, card_id int) error {
	_, err := cbaU.scheduleRepo.GetMentorIdFromScheduleBeforeAfterID(schedule_id)
	if err != nil {
		log.Error(err)
		return err
	}

	//check mentor_id with cookie

	err = cbaU.cardRepo.DeleteCardBeforeAfter(schedule_id, card_id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}