package usecase

import (
	"autfinal/internal/microservice/mentor"
	"autfinal/internal/models"
	"autfinal/internal/utils"
	log "autfinal/pkg/logger"
	"errors"
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

func (mU *mentorUsecase) DeleteMentor(id int) (error) {
	//Check with cookie
	return mU.mentorRepository.DeleteMentor(id)
}

func (mU *mentorUsecase) GetPersonalImages(mentor_id int) ([]*models.PersonalImage, error) {
	personalImages, err := mU.mentorRepository.GetPersonalImages(mentor_id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for _, image := range personalImages {
		image.ImgUrl = utils.MakeImageName(image.ImgUUID)
	}
	return personalImages, nil
}