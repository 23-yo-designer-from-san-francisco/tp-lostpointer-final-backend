package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

const (
	createCardQuery = `insert into "card_day" (name, imgurl, startTime, endTime, orderPlace, schedule_id) values ($1, $2, $3, $4, (select COUNT(id) + 1 from "card_day" where schedule_id = $5), $5) 
		returning id, name, done, imgurl, startTime, endTime, orderPlace, schedule_id;`
	createCardWOEndTimeQuery = `insert into "card_day" (name, imgurl, startTime, orderPlace, schedule_id) values ($1, $2, $3, (select COUNT(id) + 1 from "card_day" where schedule_id = $4), $4) 
		returning id, name, done, imgurl, startTime, endTime, orderPlace, schedule_id;`
	createCardWOStartTimeQuery = `insert into "card_day" (name, imgurl, orderPlace, schedule_id) values ($1, $2, (select COUNT(id) + 1 from "card_day" where schedule_id = $3), $3) 
		returning id, name, done, imgurl, startTime, endTime, orderPlace, schedule_id;`
	createCardOnlyImgQuery = `insert into "card_day" (imgUrl, orderPlace, schedule_id) values ($1, (select COUNT(id) + 1 from "card_day" where schedule_id = $2), $2) 
		returning id, done, imgurl, orderPlace, schedule_id;`
	
	getCardsQuery = `select id, name, done, imgurl, startTime, endTime, orderPlace, schedule_id from "card_day" where schedule_id = $1;`
	getCardQuery = `select id, name, done, imgurl, startTime, endTime, orderPlace, schedule_id from "card_day" where schedule_id = $1 and id = $2;`

	updateCardQuery = `update "card_day" set name = $1, done = $2, imgurl = $3, startTime = $4, endTime = $5 where schedule_id = $6 and id = $7;`
	updateCardOrder = `update "card_day" set orderPlace = $1 where schedule_id = $2 and id = $3`
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{
		db: db,
	}
}

func (cR *CardRepository) CreateCardDay(CardDay *models.CardDay, imgUrl string, schedule_id int) (*models.CardDay, error) {
	var resultCard models.CardDay

	row := cR.db.QueryRowx(createCardQuery, &CardDay.Name, &imgUrl, &CardDay.StartTime, &CardDay.EndTime, &schedule_id)
	if CardDay.EndTime == nil {
		row = cR.db.QueryRowx(createCardWOEndTimeQuery, &CardDay.Name, &imgUrl, &CardDay.StartTime, &schedule_id)
	}
	if CardDay.StartTime == nil {
		row = cR.db.QueryRowx(createCardWOStartTimeQuery, &CardDay.Name, &imgUrl, &schedule_id)
	}
	if CardDay.Name == nil {
		row = cR.db.QueryRowx(createCardOnlyImgQuery, &imgUrl, &schedule_id)
	}

	err := row.StructScan(&resultCard)
	if err != nil {
		return nil, err
	}
	return &resultCard, nil
}

func (cR *CardRepository) GetCardsDay(schedule_id int) (*[]models.CardDay, error) {
	cards := []models.CardDay{}
	err := cR.db.Select(&cards,getCardsQuery, schedule_id)
	if err != nil {
		return nil, err
	}
	return &cards, nil
}

func (cR *CardRepository) GetCardDay(scheduleID, cardID int) (*models.CardDay, error) {
	card := models.CardDay{}
	err := cR.db.Get(&card, getCardQuery, scheduleID, cardID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &card, nil
}

func (cR *CardRepository) UpdateCardDay(card *models.CardDay, scheduleID, cardID int) (*models.CardDay, error) {
	_, err := cR.db.Exec(updateCardQuery, &card.Name, &card.Done, &card.ImgUrl, &card.StartTime, &card.EndTime, &card.Schedule_ID, &card.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return card, nil
}

func (cR *CardRepository) UpdateCardsOrder(cards *models.CardsDay, schedule_id int) (*models.CardsDay, error) {
	tx, err := cR.db.Begin()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for _, card := range cards.Cards {
		_, err := tx.Exec(updateCardOrder, card.Order, schedule_id, card.ID)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return cards, nil
}