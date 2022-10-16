package mentor

import (
	"autfinal/internal/models"
)

type Usecase interface {
	CreateMentor(*models.Mentor) (*models.Mentor, error)
	GetMentor(int) (*models.Mentor, error)
	GetMentors() ([]*models.Mentor, error)
	UpdateMentor(*models.Mentor) (*models.Mentor, error)
}