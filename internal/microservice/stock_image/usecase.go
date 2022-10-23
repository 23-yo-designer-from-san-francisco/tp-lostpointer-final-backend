package image

import (
	"autfinal/internal/models"
)

type Usecase interface {
	GetStockImages(category string) ([]*models.StockImage, error)
	CreateStockImage(stockImage *models.StockImage) (*models.StockImage, error)
	SearchStockImages(searchPhrase string) ([]*models.StockImage, error)
}