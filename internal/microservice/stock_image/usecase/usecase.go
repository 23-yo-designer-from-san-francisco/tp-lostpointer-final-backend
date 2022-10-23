package usecase

import (
	"autfinal/internal/microservice/stock_image"
	"autfinal/internal/models"
	"autfinal/internal/utils"
	log "autfinal/pkg/logger"
	"errors"
)

const logMessage = "microservice:image:usecase:"

type imageUsecase struct {
	imageRepository image.Repository
}

func NewImageUsecase(imageR image.Repository) *imageUsecase {
	return &imageUsecase{
		imageRepository: imageR,
	}
}

func (iU *imageUsecase) GetStockImages(category string) ([]*models.StockImage, error) {
	message := logMessage + "GetStockImages:"
	log.Debug(message + "started")
	
	var resultImages []*models.StockImage
	var err error

	switch category {
		case "routine", "hygiene", "tasks", "leisure", "social": 
			resultImages, err = iU.imageRepository.GetStockImagesByCategory(category)
			if err != nil {
				log.Error(message + "err = ", err)
				return nil, err
			}
		default:
			resultImages, err = iU.imageRepository.GetStockImages()
			if err != nil {
				log.Error(message + "err = ", err)
				return nil, err
			}
	}
	return utils.ChangeImgUuidToImgUrlAll(resultImages), nil
}

func (iU *imageUsecase) CreateStockImage(stockImage *models.StockImage) (*models.StockImage, error) {
	message := logMessage + "GetStockImages:"
	log.Debug(message + "started")

	if stockImage.Category == "" {
		return nil, errors.New("no category")
	}
	return iU.imageRepository.CreateStockImage(stockImage)
}

func (iU *imageUsecase) SearchStockImages(searchPhrase string) ([]*models.StockImage, error) {
	message := logMessage + "SearchStockImages:"
	log.Debug(message + "started")
	
	resultImages, err :=  iU.imageRepository.SearchStockImages(searchPhrase)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return utils.ChangeImgUuidToImgUrlAll(resultImages), nil
}