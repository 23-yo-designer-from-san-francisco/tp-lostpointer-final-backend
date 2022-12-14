package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"

	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:mentor:repository:"

const (
	createMentorQuery = `insert into "mentor" (name,surname,email,password) values ($1, $2, $3, $4) 
		returning id, name, surname, email;`
	updateMentorQuery = `update "mentor" set name = $1, surname = $2, email = $3, updatedAt = now() where id = $4;`
	getMentorQuery = `select id, name, surname, email from "mentor" where id = $1 and deletedAt is null;`
	getMentorsQuery = `select id, name, surname, email from "mentor" where deletedAt is null;`
	getMentorByEmailQuery = `select id, name, surname, email from "mentor" where email = $1;`
	safeDeleteMentorQuery = `update "mentor" set deletedAt = now() where id = $1;`
	getPersonalImagesQuery = `select id, mentor_id, imguuid from "personal_image" where mentor_id = $1;`
)

type mentorRepository struct {
	db *sqlx.DB
}

func NewMentorRepository(db *sqlx.DB) *mentorRepository {
	return &mentorRepository{
		db: db,
	}
}

func (mR *mentorRepository) CreateMentor(mentor *models.Mentor) (*models.Mentor, error) {
	message := logMessage + "CreateMentor:"
	log.Debug(message + "started")

	var resultMentor models.Mentor
	err := mR.db.QueryRowx(createMentorQuery, &mentor.Name, &mentor.Surname, &mentor.Email, &mentor.Password).StructScan(&resultMentor)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &resultMentor, nil
}

func (mR *mentorRepository) UpdateMentor(mentor *models.Mentor) (*models.Mentor, error) {
	_, err := mR.db.Exec(updateMentorQuery, &mentor.Name, &mentor.Surname, &mentor.Email, &mentor.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return mentor, nil
}

func (mR *mentorRepository) GetMentor(id int) (*models.Mentor, error) {
	message := logMessage + "GetMentor:"
	log.Debug(message + "started")

	var resultMentor models.Mentor
	err := mR.db.QueryRowx(getMentorQuery, id).StructScan(&resultMentor)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return &resultMentor, nil
}

func (mR *mentorRepository) GetMentors() ([]*models.Mentor, error) {
	message := logMessage + "GetMentors:"
	log.Debug(message + "started")

	mentors := []*models.Mentor{}
	err := mR.db.Select(&mentors, getMentorsQuery)
	if err != nil {
		log.Error(message+"err = ", err)
		return nil, err
	}
	return mentors, nil
}

func (mR *mentorRepository) GetMentorByEmail(email string) (*models.Mentor, error) {
	message := logMessage + "GetMentorByEmail:"
	log.Debug(message + "started")

	var resultMentor models.Mentor
	err := mR.db.Get(&resultMentor, getMentorByEmailQuery, email)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return &resultMentor, nil
}

func (mR *mentorRepository) DeleteMentor(id int) (error) {
	message := logMessage + "DeleteMentor:"
	log.Debug(message + "started")

	_, err := mR.db.Exec(safeDeleteMentorQuery, id)
	if err != nil {
		log.Error(message + "err = ", err)
		return err
	}
	return nil
}

func (mR* mentorRepository) GetPersonalImages(mentor_id int) ([]*models.PersonalImage, error) {
	message := logMessage + "GetPersonalImages:"
	log.Debug(message + "started")

	personalImages := []*models.PersonalImage{}
	err := mR.db.Select(&personalImages, getPersonalImagesQuery, mentor_id)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return personalImages, nil
}