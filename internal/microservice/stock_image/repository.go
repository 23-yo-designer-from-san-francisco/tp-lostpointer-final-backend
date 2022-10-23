package image
import (
	"autfinal/internal/models"
)
type Repository interface {
	GetStockImagesByCategory(category string) ([]*models.StockImage, error)
	GetStockImages() ([]*models.StockImage, error)
	CreateStockImage(stockImage *models.StockImage) (*models.StockImage, error)
	SearchStockImages(searchPhrase string) ([]*models.StockImage, error)
}