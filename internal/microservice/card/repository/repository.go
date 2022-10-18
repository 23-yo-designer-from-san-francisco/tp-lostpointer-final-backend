package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:card:repository:"

const (
	createCardQuery = `insert into "card_day" (name, imguuid, startTime, endTime, orderPlace, schedule_id) values ($1, $2, $3, $4, (select COUNT(id) + 1 from "card_day" where schedule_id = $5), $5) 
		returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	createCardWOEndTimeQuery = `insert into "card_day" (name, imguuid, startTime, orderPlace, schedule_id) values ($1, $2, $3, (select COUNT(id) + 1 from "card_day" where schedule_id = $4), $4) 
		returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	createCardWOStartTimeQuery = `insert into "card_day" (name, imguuid, orderPlace, schedule_id) values ($1, $2, (select COUNT(id) + 1 from "card_day" where schedule_id = $3), $3) 
		returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	createCardOnlyImgQuery = `insert into "card_day" (imguuid, orderPlace, schedule_id) values ($1, (select COUNT(id) + 1 from "card_day" where schedule_id = $2), $2) 
		returning id, done, imguuid, orderPlace, schedule_id;`

	getCardsQuery = `select id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id from "card_day" where schedule_id = $1 order by orderPlace;`
	getCardQuery  = `select id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id from "card_day" where schedule_id = $1 and id = $2;`

	updateCardQuery = `update "card_day" set name = $1, done = $2, imguuid = $3, startTime = $4, endTime = $5 where schedule_id = $6 and id = $7 returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
	updateCardOrder = `update "card_day" set orderPlace = $1 where schedule_id = $2 and id = $3 returning id, name, done, imguuid, startTime, endTime, orderPlace, schedule_id;`
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{
		db: db,
	}
}

func (cR *CardRepository) CreateCardDay(CardDay *models.CardDay, imguuid string, schedule_id int) (*models.CardDay, error) {
	message := logMessage + "CreateCardDay:"
	log.Debug(message + "started")

	var resultCard models.CardDay
	row := cR.db.QueryRowx(createCardQuery, &CardDay.Name, &imguuid, &CardDay.StartTime, &CardDay.EndTime, &schedule_id)
	// if CardDay.EndTime == nil {
	// 	row = cR.db.QueryRowx(createCardWOEndTimeQuery, &CardDay.Name, &imguuid, &CardDay.StartTime, &schedule_id)
	// }
	// if CardDay.StartTime == nil {
	// 	row = cR.db.QueryRowx(createCardWOStartTimeQuery, &CardDay.Name, &imguuid, &schedule_id)
	// }
	// if CardDay.Name == nil {
	// 	row = cR.db.QueryRowx(createCardOnlyImgQuery, &imguuid, &schedule_id)
	// }

	err := row.StructScan(&resultCard)
	if err != nil {
		return nil, err
	}
	return &resultCard, nil
}

func (cR *CardRepository) GetCardsDay(scheduleID int) ([]*models.CardDay, error) {
	message := logMessage + "GetCardsDay:"
	log.Debug(message + "started")

	cards := []*models.CardDay{}
	err := cR.db.Select(&cards, getCardsQuery, scheduleID)
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
	err := cR.db.Get(&card, getCardQuery, scheduleID, cardID)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return &card, nil
}

func (cR *CardRepository) UpdateCardDay(card *models.CardDay, scheduleID, cardID int) (*models.CardDay, error) {
	message := logMessage + "UpdateCardDay:"
	log.Debug(message + "started")

	resultCard := models.CardDay{}
	err := cR.db.QueryRowx(updateCardQuery, &card.Name, &card.Done, &card.ImgUUID, &card.StartTime, &card.EndTime, &scheduleID, &cardID).StructScan(&resultCard)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return card, nil
}

func (cR *CardRepository) UpdateCardsOrder(cards []*models.CardDay, schedule_id int) error {
	message := logMessage + "UpdateCardsOrder:"
	log.Debug(message + "started")

	tx, err := cR.db.Begin()
	if err != nil {
		log.Error(message+"err = ", err)
		return err
	}
	for _, card := range cards {
		_, err := tx.Exec(updateCardOrder, card.Order, schedule_id, card.ID)
		if err != nil {
			log.Error(message+"err = ", err)
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
