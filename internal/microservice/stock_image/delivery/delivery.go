package delivery

import (
	"autfinal/internal/microservice/stock_image"
	"autfinal/internal/models"
	log "autfinal/pkg/logger"
	"net/http"
	"autfinal/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

const logMessage = "microservice:image:delivery:"

type ImageDelivery struct {
	imageUsecase image.Usecase
}

func NewImageDelivery(imageU image.Usecase) *ImageDelivery {
	return &ImageDelivery{
		imageUsecase: imageU,
	}
}

func (iD *ImageDelivery) GetStockImages(c *gin.Context) {
	message := logMessage + "GetStockImages:"
	log.Debug(message + "started")

	category := c.DefaultQuery("category","")
	resultImages, err := iD.imageUsecase.GetStockImages(category)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}
	response := &models.Response{
		Status: http.StatusOK,
		Response: resultImages,
	}
	c.JSON(http.StatusOK, response)
}

func (iD *ImageDelivery) CreateStockImage(c *gin.Context) {
	message := logMessage + "CreateStockImage:"
	log.Debug(message + "started")

	category := c.DefaultQuery("category","")
	stockImageStr := c.Request.FormValue("stock")
	log.Debug(stockImageStr)
	var stockImage = &models.StockImage{}
	if stockImageStr != "" {
		json.Unmarshal([]byte(stockImageStr), &stockImage)
	}
	log.Debug(stockImage.Names)
	
	imgUUID, err := utils.SaveImageFromRequest(c, "image")
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}
	log.Debug(imgUUID)

	stockImage.ImgUUID = imgUUID;
	stockImage.Category = category

	resultImage, err := iD.imageUsecase.CreateStockImage(stockImage)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}
	response := &models.Response{
		Status: http.StatusOK,
		Response: resultImage,
	}
	c.JSON(http.StatusOK, response)
}

func (iD *ImageDelivery) SearchStockImages(c *gin.Context) {
	message := logMessage + "CreateStockImage:"
	log.Debug(message + "started")

	searchPhrase := c.DefaultQuery("phrase","")

	resultImages, err := iD.imageUsecase.SearchStockImages(searchPhrase)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}
	response := &models.Response{
		Status: http.StatusOK,
		Response: resultImages,
	}
	c.JSON(http.StatusOK, response)
}