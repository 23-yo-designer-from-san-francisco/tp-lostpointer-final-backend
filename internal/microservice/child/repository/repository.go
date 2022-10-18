package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"

	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:child:repository:"

const (
	createChildQuery = `insert into "child" (name, mentor_id) values ($1, $2) 
		returning id, name;`
	createChildWithDOBQuery = `insert into "child" (name, date_of_birth, mentor_id) values ($1, $2, $3) 
		returning id, name, date_of_birth;`
	getChildQuery = `select id, name, date_of_birth from "child" where id = $1 and mentor_id = $2`
	getChildsQuery = `select id, name, date_of_birth from "child" where mentor_id = $1 and deletedAt is null`
	updateChildQuery = `update "child" set name = $1, date_of_birth = $2, updatedAt = now() where mentor_id = $3 and id = $4 
		returning id, name, date_of_birth, mentor_id;`
	safeDeleteChildQuery = `update "child" set deletedAt = now() where id = $1`
)

type childRepository struct {
	db *sqlx.DB
}

func NewChildRepository(db *sqlx.DB) *childRepository {
	return &childRepository{
		db: db,
	}
}

func(cR *childRepository) CreateChild(child *models.Child) (*models.Child, error) {
	var resultChild models.Child
	var row *sqlx.Row

	//row = cR.db.QueryRowx(createChildQuery, &child.Name, &child.Mentor_ID)
	row = cR.db.QueryRowx(createChildWithDOBQuery, &child.Name, child.DateOfBirth, &child.Mentor_ID)
	
	err := row.StructScan(&resultChild)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &resultChild, nil
}

func(cR *childRepository) GetChild(id, mentorID int) (*models.Child, error) {
	message := logMessage + "GetChild:"
	log.Debug(message + "started")

	var resultChild models.Child
	err := cR.db.Get(&resultChild, getChildQuery, id, mentorID)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return &resultChild, nil
}

func(cR *childRepository) GetChilds(mentorID int) ([]*models.Child, error) {
	message := logMessage + "GetChilds:"
	log.Debug(message + "started")

	resultChilds := []*models.Child{}
	err := cR.db.Select(&resultChilds, getChildsQuery, mentorID)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return resultChilds, nil
}

func(cR *childRepository) UpdateChild(child *models.Child) (*models.Child, error) {
	message := logMessage + "UpdateChild:"
	log.Debug(message + "started")

	resultChild := models.Child{}

	err := cR.db.QueryRowx(updateChildQuery, child.Name, child.DateOfBirth, child.Mentor_ID, child.ID).StructScan(&resultChild)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return &resultChild, nil
}

func (cR *childRepository) DeleteChild(id int) error {
	message := logMessage + "DeleteChild:"
	log.Debug(message + "started")

	_, err := cR.db.Exec(safeDeleteChildQuery, id)
	if err != nil {
		log.Error(message + "err = ", err)
		return err
	}
	return nil
}