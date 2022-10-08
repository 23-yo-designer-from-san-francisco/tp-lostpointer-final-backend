package repository

import (
	"autfinal/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	createCardQuery = `insert into "card" (name, imgurl, startTime, endTime) values ($1, $2, $3, $4) returning id, name, done, imgurl, startTime, endTime`
	getCardsQuery = `select id, name, done, imgurl, startTime, endTime from "card"`
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{
		db: db,
	}
}

func (cR *CardRepository) CreateCard(card *models.Card) (*models.Card, error) {
	var resultCard models.Card
	err := cR.db.QueryRowx(createCardQuery, &card.Name, &card.ImgUrl, &card.StartTime, &card.EndTime).StructScan(&resultCard)
	if err != nil {
		return nil, err
	}
	return &resultCard, nil
}

func (cR *CardRepository) GetCards() (*[]models.Card, error) {
	cards := []models.Card{}
	err := cR.db.Select(&cards,getCardsQuery)
	if err != nil {
		return nil, err
	}
	return &cards, nil
}