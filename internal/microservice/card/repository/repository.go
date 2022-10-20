package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:card:repository:"

const (
	createCardDayQuery = `insert into "card_day" (name, imguuid, startTime, endTime, orderPlace, schedule_id) values ($1, $2, $3, $4, (select cards_count from "schedule_day" where "id" = $5) + 1, $5) 
		returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	// createCardWOEndTimeQuery = `insert into "card_day" (name, imguuid, startTime, orderPlace, schedule_id) values ($1, $2, $3, (select COUNT(id) + 1 from "card_day" where schedule_id = $4), $4) 
	// 	returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	// createCardWOStartTimeQuery = `insert into "card_day" (name, imguuid, orderPlace, schedule_id) values ($1, $2, (select COUNT(id) + 1 from "card_day" where schedule_id = $3), $3) 
	// 	returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	// createCardOnlyImgQuery = `insert into "card_day" (imguuid, orderPlace, schedule_id) values ($1, (select COUNT(id) + 1 from "card_day" where schedule_id = $2), $2) 
	// 	returning id, done, imguuid, orderPlace, schedule_id;`

	getCardsDayQuery = `select id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id from "card_day" where schedule_id = $1 and deletedAt is null order by orderPlace;`
	getCardDayQuery  = `select id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id from "card_day" where schedule_id = $1 and id = $2;`

	updateCardDayQuery = `update "card_day" set name = $1, done = $2, imguuid = $3, startTime = $4, endTime = $5 where schedule_id = $6 and id = $7 
		returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	updateCardDayWOImgQuery = `update "card_day" set name = $1, done = $2, startTime = $3, endTime = $4 where schedule_id = $5 and id = $6 
		returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	updateCardDayOrder = `update "card_day" set orderPlace = $1 where schedule_id = $2 and id = $3 
		returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	safeDeleteCardDay = `update "card_day" set deletedAt = now() where schedule_id = $1 and id = $2 returning orderPlace;`
	changeOrderCardsDayAfterDelete = `update "card_day" set orderPlace = (orderPlace - 1) where orderPlace > $1;`

	createCardLessonQuery = `insert into "card_lesson" (name, imguuid, duration, orderPlace, schedule_id) values ($1, $2, $3, (select cards_count from "schedule_lesson" where "id" = $4) + 1, $4) 
		returning id, name, done, imguuid, duration, orderPlace, schedule_id;`
	getCardsLessonQuery = `select id, name, done, imguuid, duration, orderPlace, schedule_id from "card_lesson" where schedule_id = $1 and deletedAt is null order by orderPlace;`
	getCardLessonQuery = `select id, name, done, imguuid, duration, orderPlace, schedule_id from "card_lesson" where schedule_id = $1 and id = $2;`
	updateCardLessonQuery = `update "card_lesson" set name = $1, done = $2, imguuid = $3, duration = $4 where schedule_id = $5 and id = $6 
		returning id, name, done, imguuid, duration, orderPlace, schedule_id;`
	updateCardLessonWOImgQuery = `update "card_lesson" set name = $1, done = $2, duration = $3 where schedule_id = $4 and id = $5
		returning id, name, done, imguuid, duration, orderPlace, schedule_id;`
	updateCardLessonOrder = `update "card_lesson" set orderPlace = $1 where schedule_id = $2 and id = $3 
		returning id, name, done, imguuid, duration, orderPlace, schedule_id;`
	safeDeleteCardLesson = `update "card_lesson" set deletedAt = now() where schedule_id = $1 and id = $2 returning orderPlace;`
	changeOrderCardsLessonAfterDelete = `update "card_lesson" set orderPlace = (orderPlace - 1) where orderPlace > $1;`

	savePersonalImageQuery = `insert into "personal_image" (imguuid, mentor_id) values ($1, $2);`
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{
		db: db,
	}
}

func (cR *CardRepository) CreateCardDay(CardDay *models.CardDay, mentor_id int) (*models.CardDay, error) {
	message := logMessage + "CreateCardDay:"
	log.Debug(message + "started")

	var resultCard models.CardDay
	tx, err :=  cR.db.Beginx()
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	err = tx.QueryRowx(createCardDayQuery, &CardDay.Name, &CardDay.ImgUUID, &CardDay.StartTime, &CardDay.EndTime, &CardDay.Schedule_ID).StructScan(&resultCard)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(savePersonalImageQuery, &CardDay.ImgUUID, &mentor_id)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &resultCard, nil
}

func (cR *CardRepository) CreateCardLesson(CardLesson *models.CardLesson, mentor_id int) (*models.CardLesson, error) {
	message := logMessage + "CreateCardLesson:"
	log.Debug(message + "started")

	var resultCard models.CardLesson
	tx, err := cR.db.Beginx()
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}

	err = tx.QueryRowx(createCardLessonQuery, &CardLesson.Name, &CardLesson.ImgUUID, &CardLesson.Duration, &CardLesson.Schedule_ID).StructScan(&resultCard)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(savePersonalImageQuery, &CardLesson.ImgUUID, &mentor_id)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &resultCard, nil
}

func (cR *CardRepository) GetCardsDay(scheduleID int) ([]*models.CardDay, error) {
	message := logMessage + "GetCardsDay:"
	log.Debug(message + "started")

	cards := []*models.CardDay{}
	err := cR.db.Select(&cards, getCardsDayQuery, scheduleID)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}

	return cards, nil
}

func (cR *CardRepository) GetCardsLesson(scheduleID int) ([]*models.CardLesson, error) {
	message := logMessage + "GetCardsLesson:"
	log.Debug(message + "started")

	cards := []*models.CardLesson{}
	err := cR.db.Select(&cards, getCardsLessonQuery, scheduleID)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}

	return cards, nil
}

func (cR *CardRepository) GetCardDay(scheduleID, cardID int) (*models.CardDay, error) {
	message := logMessage + "GetCardDay:"
	log.Debug(message + "started")

	card := models.CardDay{}
	err := cR.db.Get(&card, getCardDayQuery, scheduleID, cardID)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return &card, nil
}

func (cR *CardRepository) GetCardLesson(scheduleID, cardID int) (*models.CardLesson, error) {
	message := logMessage + "GetCardLesson:"
	log.Debug(message + "started")

	card := models.CardLesson{}
	err := cR.db.Get(&card, getCardLessonQuery, scheduleID, cardID)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return &card, nil
}

func (cR *CardRepository) UpdateCardDay(card *models.CardDay, scheduleID, cardID int) (*models.CardDay, error) {
	message := logMessage + "UpdateCardDay:"
	log.Debug(message + "started")
	
	var err error
	resultCard := models.CardDay{}
	if card.ImgUUID == "" {
		err = cR.db.QueryRowx(updateCardDayWOImgQuery, &card.Name, &card.Done, &card.StartTime, &card.EndTime, &scheduleID, &cardID).StructScan(&resultCard)
	} else {
		err = cR.db.QueryRowx(updateCardDayQuery, &card.Name, &card.Done, &card.ImgUUID, &card.StartTime, &card.EndTime, &scheduleID, &cardID).StructScan(&resultCard)
	}
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return &resultCard, nil
}

func (cR *CardRepository) UpdateCardLesson(card *models.CardLesson, scheduleID, cardID int) (*models.CardLesson, error) {
	message := logMessage + "UpdateCardLesson:"
	log.Debug(message + "started")

	var err error
	resultCard := models.CardLesson{}
	if card.ImgUUID == "" {
		err = cR.db.QueryRowx(updateCardLessonWOImgQuery, &card.Name, &card.Done, &card.Duration, &scheduleID, &cardID).StructScan(&resultCard)	
	} else {
		err = cR.db.QueryRowx(updateCardLessonQuery, &card.Name, &card.Done, &card.ImgUUID, &card.Duration, &scheduleID, &cardID).StructScan(&resultCard)
	}
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return &resultCard, nil
}

func (cR *CardRepository) UpdateCardsDayOrder(cards []*models.CardDay, schedule_id int) error {
	message := logMessage + "UpdateCardsDayOrder:"
	log.Debug(message + "started")

	tx, err := cR.db.Begin()
	if err != nil {
		log.Error(message+"err = ", err)
		return err
	}
	for _, card := range cards {
		_, err := tx.Exec(updateCardDayOrder, card.Order, schedule_id, card.ID)
		if err != nil {
			log.Error(message+"err = ", err)
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (cR *CardRepository) UpdateCardsLessonOrder(cards []*models.CardLesson, schedule_id int) error {
	message := logMessage + "UpdateCardsLessonOrder:"
	log.Debug(message + "started")

	tx, err := cR.db.Begin()
	if err != nil {
		log.Error(message+"err = ", err)
		return err
	}
	for _, card := range cards {
		_, err := tx.Exec(updateCardLessonOrder, card.Order, schedule_id, card.ID)
		if err != nil {
			log.Error(message+"err = ", err)
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (cR *CardRepository) DeleteCardDay(scheduleID, cardID int) (error) {
	message := logMessage + "DeleteCardDay:"
	log.Debug(message + "started")

	tx, err := cR.db.Beginx()
	if err != nil {
		log.Error(message+"err = ", err)
		return err
	}
	var deletedOrderPlace int
	err = tx.QueryRowx(safeDeleteCardDay, &scheduleID, &cardID).Scan(&deletedOrderPlace)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(changeOrderCardsDayAfterDelete, &deletedOrderPlace) 
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (cR *CardRepository) DeleteCardLesson(scheduleID, cardID int) (error) {
	message := logMessage + "DeleteCardLesson:"
	log.Debug(message + "started")

	tx, err := cR.db.Beginx()
	if err != nil {
		log.Error(message+"err = ", err)
		return err
	}
	var deletedOrderPlace int
	err = tx.QueryRowx(safeDeleteCardLesson, &scheduleID, &cardID).Scan(&deletedOrderPlace)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(changeOrderCardsLessonAfterDelete, &deletedOrderPlace) 
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}