package mentor

import (
	"autfinal/internal/models"
)

type Repository interface {
	CreateMentor(*models.Mentor) (*models.Mentor, error)
	GetMentor(int) (*models.Mentor, error)
	GetMentors() (*models.Mentors, error)
	UpdateMentor(*models.Mentor) (*models.Mentor, error)
	GetMentorByEmail(string) (*models.Mentor, error)
}