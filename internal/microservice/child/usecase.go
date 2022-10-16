package child

import (
	"autfinal/internal/models"
)

type Usecase interface {
	CreateChild(*models.Child) (*models.Child, error)
	GetChild(int) (*models.Child, error)
	GetChilds() ([]*models.Child, error)
}