package repository

import (
	"autfinal/internal/models"
	"autfinal/internal/utils/queries"
	log "autfinal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:card:repository:"

const (
	createCardBeforeAfterQuery = `insert into "card_before_after" (name, imguuid, orderPlace, schedule_id) values ($1, $2, (select cards_count from "schedule_day" where "id" = $3) + 1, $3) 
		returning id, name, done, imguuid, orderPlace, schedule_id;`

	getCardsBeforeAfterQuery = `select id, name, done, imguuid, orderPlace, schedule_id from "card_before_after" where schedule_id = $1 and deletedAt is null order by orderPlace;`
	getCardBeforeAfterQuery  = `select id, name, done, imguuid, orderPlace, schedule_id from "card_before_after" where schedule_id = $1 and id = $2;`

	updateCardBeforeAfterQuery = `update "card_before_after" set name = $1, done = $2, imguuid = $3 where schedule_id = $4 and id = $5 
		returning id, name, done, imguuid, orderPlace, schedule_id;`
	updateCardBeforeAfterWOImgQuery = `update "card_before_after" set name = $1, done = $2, where schedule_id = $3 and id = $4 
		returning id, name, done, imguuid, orderPlace, schedule_id;`
	updateCardBeforeAfterOrder = `update "card_before_after" set orderPlace = $1 where schedule_id = $2 and id = $3 
		returning id, name, done, imguuid, orderPlace, schedule_id;`
	safeDeleteCardBeforeAfter              = `update "card_before_after" set deletedAt = now() where schedule_id = $1 and id = $2 returning orderPlace;`
	changeOrderCardsBeforeAfterAfterDelete = `update "card_before_after" set orderPlace = (orderPlace - 1) where orderPlace > $1;`
)

type CardBeforeAfterRepository struct {
	db *sqlx.DB
}

func NewBeforeAfterRepository(db *sqlx.DB) *CardBeforeAfterRepository {
	return &CardBeforeAfterRepository{
		db: db,
	}
}

func (cbaR *CardBeforeAfterRepository) CreateCardBeforeAfter(CardBeforeAfter *models.CardBeforeAfter, mentor_id int) (*models.CardBeforeAfter, error) {
	message := logMessage + "CreateCardBeforeAfter:"
	log.Debug(message + "started")

	var resultCard models.CardBeforeAfter
	tx, err := cbaR.db.Beginx()
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}

	err = tx.QueryRowx(createCardBeforeAfterQuery, &CardBeforeAfter.Name, &CardBeforeAfter.ImgUUID, &CardBeforeAfter.Schedule_ID).StructScan(&resultCard)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(queries.SavePersonalImageQuery, &CardBeforeAfter.ImgUUID, &mentor_id)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &resultCard, nil
}

func (cbaR *CardBeforeAfterRepository) GetCardsBeforeAfter(scheduleID int) ([]*models.CardBeforeAfter, error) {
	message := logMessage + "GetCardsBeforeAfter:"
	log.Debug(message + "started")

	cards := []*models.CardBeforeAfter{}
	err := cbaR.db.Select(&cards, getCardsBeforeAfterQuery, scheduleID)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}

	return cards, nil
}

func (cbaR *CardBeforeAfterRepository) GetCardBeforeAfter(scheduleID, cardID int) (*models.CardBeforeAfter, error) {
	message := logMessage + "GetCardBeforeAfter:"
	log.Debug(message + "started")

	card := models.CardBeforeAfter{}
	err := cbaR.db.Get(&card, getCardBeforeAfterQuery, scheduleID, cardID)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return &card, nil
}

func (cbaR *CardBeforeAfterRepository) UpdateCardBeforeAfter(card *models.CardBeforeAfter, scheduleID, cardID int) (*models.CardBeforeAfter, error) {
	message := logMessage + "UpdateCardBeforeAfter:"
	log.Debug(message + "started")

	var err error
	resultCard := models.CardBeforeAfter{}
	if card.ImgUUID == "" {
		err = cbaR.db.QueryRowx(updateCardBeforeAfterWOImgQuery, &card.Name, &card.Done, &scheduleID, &cardID).StructScan(&resultCard)
	} else {
		err = cbaR.db.QueryRowx(updateCardBeforeAfterQuery, &card.Name, &card.Done, &card.ImgUUID, &scheduleID, &cardID).StructScan(&resultCard)
	}
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return &resultCard, nil
}

func (cbaR *CardBeforeAfterRepository) UpdateCardsBeforeAfterOrder(cards []*models.CardBeforeAfter, schedule_id int) error {
	message := logMessage + "UpdateCardsBeforeAfterOrder:"
	log.Debug(message + "started")

	tx, err := cbaR.db.Begin()
	if err != nil {
		log.Error(message+"err = ", err)
		return err
	}
	for _, card := range cards {
		_, err := tx.Exec(updateCardBeforeAfterOrder, card.Order, schedule_id, card.ID)
		if err != nil {
			log.Error(message+"err = ", err)
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (cbaR *CardBeforeAfterRepository) DeleteCardBeforeAfter(scheduleID, cardID int) error {
	message := logMessage + "DeleteCardBeforeAfter:"
	log.Debug(message + "started")

	tx, err := cbaR.db.Beginx()
	if err != nil {
		log.Error(message+"err = ", err)
		return err
	}
	var deletedOrderPlace int
	err = tx.QueryRowx(safeDeleteCardBeforeAfter, &scheduleID, &cardID).Scan(&deletedOrderPlace)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(changeOrderCardsBeforeAfterAfterDelete, &deletedOrderPlace)
	if err != nil {
		log.Error(message+"err = ", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}