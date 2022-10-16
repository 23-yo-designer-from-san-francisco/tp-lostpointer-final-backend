package delivery

import (
	"autfinal/internal/microservice/card"
	"autfinal/internal/models"
	"autfinal/internal/utils"
	log "autfinal/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type CardDelivery struct {
	cardUsecase card.Usecase
}

func NewCardDelivery(cardU card.Usecase) *CardDelivery {
	return &CardDelivery{
		cardUsecase: cardU,
	}
}

func(cD *CardDelivery) CreateCardDay(c *gin.Context) {
	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr) 
	if err != nil {
		log.Error(err)
		return
	}

	cardJsonStr := c.Request.FormValue("card")
	var card = &models.CardDay{}
	if cardJsonStr != "" {
		json.Unmarshal([]byte(cardJsonStr), &card)
	}

	imgUrl, err := utils.SaveImageFromRequest(c,"image")
	if err != nil {
		log.Error(err)
		return
	}
	if err == nil {
		card.ImgUrl = imgUrl
	}

	resultCard, err := cD.cardUsecase.CreateCardDay(card, imgUrl, scheduleID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultCard,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) GetCardsDay(c *gin.Context) {
	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr) 
	if err != nil {
		log.Error(err)
		return
	}

	resultCards, err := cD.cardUsecase.GetCardsDay(scheduleID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: &models.CardsDay{
			Cards: *resultCards,
		},
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) GetCardDay(c *gin.Context) {
	cardIdStr := c.Param("card_id")
	cardID, err := strconv.Atoi(cardIdStr)
	if err != nil {
		log.Error()
		return 
	}

	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error()
		return 
	}

	resultCard, err := cD.cardUsecase.GetCardDay(scheduleID, cardID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultCard,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) UpdateCardDay(c *gin.Context) {
	cardIdStr := c.Param("card_id")
	cardID, err := strconv.Atoi(cardIdStr)
	if err != nil {
		log.Error()
		return 
	}

	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error()
		return 
	}

	cardJsonStr := c.Request.FormValue("card")
	var card = &models.CardDay{}
	if cardJsonStr != "" {
		json.Unmarshal([]byte(cardJsonStr), &card)
	}

	imgUrl, err := utils.SaveImageFromRequest(c,"image")
	if err != nil {
		c.Error(err)
		return
	}
	if err == nil {
		card.ImgUrl = imgUrl
	}

	resultCard, err := cD.cardUsecase.UpdateCardDay(card, scheduleID, cardID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultCard,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) UpdateCardsOrder(c *gin.Context) {
	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error(err)
		return 
	}

	var cards = &models.CardsDay{}
	err = c.ShouldBindJSON(cards)
	if err != nil {
		log.Error(err)
		return
	}

	resultCards, err := cD.cardUsecase.UpdateCardsOrder(cards, scheduleID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultCards,
	}

	c.JSON(http.StatusOK, &response)
}