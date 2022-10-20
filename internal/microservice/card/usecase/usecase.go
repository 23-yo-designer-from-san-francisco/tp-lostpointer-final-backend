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
	cardRepo     card.Repository
}

func NewCardUsecase(scheduleR schedule.Repository, cardR card.Repository) *CardUsecase {
	return &CardUsecase{
		scheduleRepo: scheduleR,
		cardRepo:     cardR,
	}
}

func (cU *CardUsecase) CreateCardDay(card *models.CardDay, schedule_id int) (*models.CardDay, error) {
	mentor_id, err := cU.scheduleRepo.GetMentorIdFromScheduleDayID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie
	card.Schedule_ID = schedule_id
	resultCard, err := cU.cardRepo.CreateCardDay(card, mentor_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cU *CardUsecase) CreateCardLesson(card *models.CardLesson, schedule_id int) (*models.CardLesson, error) {
	mentor_id, err := cU.scheduleRepo.GetMentorIdFromScheduleLessonID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie
	card.Schedule_ID = schedule_id
	resultCard, err := cU.cardRepo.CreateCardLesson(card, mentor_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cU *CardUsecase) GetCardsDay(schedule_id int) ([]*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleDayID(schedule_id)
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

	for _, resultCard := range resultCards {
		resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	}

	return resultCards, nil
}

func (cU *CardUsecase) GetCardsLesson(schedule_id int) ([]*models.CardLesson, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleLessonID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCards, err := cU.cardRepo.GetCardsLesson(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, resultCard := range resultCards {
		resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	}

	return resultCards, nil
}

func (cU *CardUsecase) GetCardDay(schedule_id, card_id int) (*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleDayID(schedule_id)
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

func (cU *CardUsecase) GetCardLesson(schedule_id, card_id int) (*models.CardLesson, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleLessonID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCard, err := cU.cardRepo.GetCardLesson(schedule_id, card_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cU *CardUsecase) UpdateCardDay(card *models.CardDay, schedule_id, card_id int) (*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleDayID(schedule_id)
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

func (cU *CardUsecase) UpdateCardLesson(card *models.CardLesson, schedule_id, card_id int) (*models.CardLesson, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleDayID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie

	resultCard, err := cU.cardRepo.UpdateCardLesson(card, schedule_id, card_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	return resultCard, nil
}

func (cU *CardUsecase) UpdateCardsDayOrder(cards []*models.CardDay, schedule_id int) ([]*models.CardDay, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleDayID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie
	err = cU.cardRepo.UpdateCardsDayOrder(cards, schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCards, err := cU.cardRepo.GetCardsDay(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, resultCard := range resultCards {
		resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	}

	return resultCards, nil
}

func (cU *CardUsecase) UpdateCardsLessonOrder(cards []*models.CardLesson, schedule_id int) ([]*models.CardLesson, error) {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleDayID(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//check mentor_id with cookie
	err = cU.cardRepo.UpdateCardsLessonOrder(cards, schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultCards, err := cU.cardRepo.GetCardsLesson(schedule_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, resultCard := range resultCards {
		resultCard.ImgUrl = utils.MakeImageName(resultCard.ImgUUID)
	}

	return resultCards, nil
}

func (cU *CardUsecase) DeleteCardDay(schedule_id, card_id int) error {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleDayID(schedule_id)
	if err != nil {
		log.Error(err)
		return err
	}

	//check mentor_id with cookie

	err = cU.cardRepo.DeleteCardDay(schedule_id, card_id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (cU *CardUsecase) DeleteCardLesson(schedule_id, card_id int) error {
	_, err := cU.scheduleRepo.GetMentorIdFromScheduleLessonID(schedule_id)
	if err != nil {
		log.Error(err)
		return err
	}

	//check mentor_id with cookie

	err = cU.cardRepo.DeleteCardLesson(schedule_id, card_id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}