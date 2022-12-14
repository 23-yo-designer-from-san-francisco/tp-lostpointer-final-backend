package child

import (
	"autfinal/internal/models"
)

type Repository interface {
	CreateChild(*models.Child) (*models.Child, error)
	GetChild(int, int) (*models.Child, error)
	GetChilds(int) ([]*models.Child, error)
	UpdateChild(*models.Child) (*models.Child, error)
	DeleteChild(int) (error)
}