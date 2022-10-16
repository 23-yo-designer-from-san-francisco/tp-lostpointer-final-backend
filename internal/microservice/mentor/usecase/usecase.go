package usecase

import (
	"autfinal/internal/microservice/mentor"
	"autfinal/internal/models"
	"errors"
	log "autfinal/pkg/logger"
)

const logMessage = "microservice:mentor:usecase:"

type mentorUsecase struct {
	mentorRepository mentor.Repository
}

func NewMentorUsecase(mentorR mentor.Repository) *mentorUsecase {
	return &mentorUsecase{
		mentorRepository: mentorR,
	}
}

func (mU *mentorUsecase) CreateMentor(mentor *models.Mentor) (*models.Mentor, error) {
	message := logMessage + "CreateMentor:"
	log.Debug(message + "started")

	existingMentor, err := mU.mentorRepository.GetMentorByEmail(mentor.Email)
	if existingMentor == nil {
		log.Error(err)
		return mU.mentorRepository.CreateMentor(mentor)
	}
	return nil, errors.New("user with this email exists")
}

func (mU *mentorUsecase) UpdateMentor(mentor *models.Mentor) (*models.Mentor, error) {
	mentor.ID = 1
	return mU.mentorRepository.UpdateMentor(mentor)
}

func (mU *mentorUsecase) GetMentor(id int) (*models.Mentor, error) {
	return mU.mentorRepository.GetMentor(id)
}

func (mU *mentorUsecase) GetMentors() ([]*models.Mentor, error) {
	return mU.mentorRepository.GetMentors()
}