package usecase

import (
	"autfinal/internal/microservice/child"
	"autfinal/internal/models"

	log "autfinal/pkg/logger"
)

const logMessage = "microservice:child:usecase:"

type childUsecase struct {
	childRepository child.Repository
}

func NewChildUsecase(childR child.Repository) *childUsecase {
	return &childUsecase{
		childRepository: childR,
	}
}

func (cU *childUsecase) CreateChild(child *models.Child) (*models.Child, error) {
	message := logMessage + "CreateChild:"
	log.Debug(message + "started")

	child.Mentor_ID = 1;
	return cU.childRepository.CreateChild(child)
}

func (cU *childUsecase) GetChild(id int) (*models.Child, error) {
	message := logMessage + "GetChild:"
	log.Debug(message + "started")

	mentorID := 1
	return cU.childRepository.GetChild(id, mentorID)
}

func (cU *childUsecase) GetChilds() ([]*models.Child, error) {
	message := logMessage + "GetChilds:"
	log.Debug(message + "started")

	mentorID := 1
	return cU.childRepository.GetChilds(mentorID)
}

func (cU *childUsecase) UpdateChild(child *models.Child) (*models.Child, error) {
	message := logMessage + "UpdateChild:"
	log.Debug(message + "started")

	child.Mentor_ID = 1;
	return cU.childRepository.UpdateChild(child)
}

func (cU *childUsecase) DeleteChild(id int) error {
	message := logMessage + "DeleteChild:"
	log.Debug(message + "started")

	return cU.childRepository.DeleteChild(id)
}