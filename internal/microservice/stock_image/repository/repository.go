package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:card:repository:"

const (
	createStockImage = "insert into stock_image (category, imguuid) values ($1, $2) returning id, category, imguuid;"
	saveStockImageName = "insert into stock_image_name (name, image_id) values ($1, $2);"

	GetStockImages = "select id, category, imguuid from stock_image;"
	GetStockImagesByCategory = "select id, category, imguuid from stock_image where category = $1;"

	SearchStockImages = `select id, category, imguuid from stock_image 
		where id in (select distinct image_id from stock_image_name where name % $1);`
)

type ImageRepository struct {
	db *sqlx.DB
}

func NewImageRepository(db *sqlx.DB) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}

func(iR *ImageRepository) CreateStockImage(stockImage *models.StockImage) (*models.StockImage, error) {
	message := logMessage + "CreateStockImage:"
	log.Debug(message + "started")

	tx, err := iR.db.Beginx()
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	var resultStockImage models.StockImage
	err = tx.QueryRowx(createStockImage, &stockImage.Category, &stockImage.ImgUUID).StructScan(&resultStockImage)
	if err != nil {
		log.Error(message + "err = ", err)
		tx.Rollback()
		return nil, err
	}

	for _, name := range stockImage.Names {
		_, err := tx.Exec(saveStockImageName, &name, &resultStockImage.ID)
		if err != nil {
			log.Error(message + "err = ", err)
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	
	return &resultStockImage, nil
}

func (iR *ImageRepository) GetStockImages() ([]*models.StockImage, error) {
	message := logMessage + "GetStockImages:"
	log.Debug(message + "started")

	var resultImages []*models.StockImage
	err := iR.db.Select(&resultImages,GetStockImages)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return resultImages, nil
}

func (iR *ImageRepository) GetStockImagesByCategory(category string) ([]*models.StockImage, error) {
	message := logMessage + "GetStockImages:"
	log.Debug(message + "started")

	var resultImages []*models.StockImage
	err := iR.db.Select(&resultImages,GetStockImagesByCategory, category)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return resultImages, nil
}
 
func (iR *ImageRepository) SearchStockImages(searchPhrase string) ([]*models.StockImage, error) {
	message := logMessage + "SearchStockImages:"
	log.Debug(message + "started")

	var resultImages []*models.StockImage
	err := iR.db.Select(&resultImages, SearchStockImages, searchPhrase)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return resultImages, nil
}