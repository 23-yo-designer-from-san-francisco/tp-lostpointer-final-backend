package delivery

import (
	"autfinal/internal/microservice/card"
	"autfinal/internal/models"
	"autfinal/internal/utils"
	log "autfinal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strconv"
)

const logMessage = "microservice:card:delivery:"

type CardDelivery struct {
	cardUsecase card.Usecase
}

func NewCardDelivery(cardU card.Usecase) *CardDelivery {
	return &CardDelivery{
		cardUsecase: cardU,
	}
}

func (cD *CardDelivery) CreateCardDay(c *gin.Context) {
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

	imgUUID, err := utils.SaveImageFromRequest(c, "image")
	if err != nil {
		log.Error(err)
		return
	}

	if err == nil {
		card.ImgUUID = imgUUID
	}

	resultCard, err := cD.cardUsecase.CreateCardDay(card, imgUUID, scheduleID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
		Response: resultCard,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) CreateCardLesson(c *gin.Context) {
	message := logMessage + "CreateCardLesson:"
	log.Debug(message + "started")

	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error(err)
		return
	}

	cardJsonStr := c.Request.FormValue("card")
	var card = &models.CardLesson{}
	if cardJsonStr != "" {
		json.Unmarshal([]byte(cardJsonStr), &card)
	}

	imgUUID, err := utils.SaveImageFromRequest(c, "image")
	if err != nil {
		log.Error(err)
		return
	}

	if err == nil {
		card.ImgUUID = imgUUID
	}

	resultCard, err := cD.cardUsecase.CreateCardLesson(card, imgUUID, scheduleID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
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
		Status:   http.StatusOK,
		Response: resultCards,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) GetCardsLesson(c *gin.Context) {
	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error(err)
		return
	}

	resultCards, err := cD.cardUsecase.GetCardsLesson(scheduleID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
		Response: resultCards,
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
		log.Error()
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
		Response: resultCard,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) GetCardLesson(c *gin.Context) {
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

	resultCard, err := cD.cardUsecase.GetCardLesson(scheduleID, cardID)
	if err != nil {
		log.Error()
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
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
	
	imgUUID, err := utils.SaveImageFromRequest(c, "image")
	if err != nil {
		log.Error(err)
	}
	if err == nil {
		card.ImgUUID = imgUUID
	}

	resultCard, err := cD.cardUsecase.UpdateCardDay(card, scheduleID, cardID)
	if err != nil {
		log.Error(err)
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
		Response: resultCard,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) UpdateCardLesson(c *gin.Context) {
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
	var card = &models.CardLesson{}
	if cardJsonStr != "" {
		json.Unmarshal([]byte(cardJsonStr), &card)
	}

	imgUUID, err := utils.SaveImageFromRequest(c, "image")
	if err != nil {
		log.Error(err)
		return
	}
	if err == nil {
		card.ImgUUID = imgUUID
	}

	resultCard, err := cD.cardUsecase.UpdateCardLesson(card, scheduleID, cardID)
	if err != nil {
		log.Error(err)
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
		Response: resultCard,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) UpdateCardsDayOrder(c *gin.Context) {
	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error(err)
		return
	}

	var cards []*models.CardDay
	err = c.ShouldBindJSON(cards)
	if err != nil {
		log.Error(err)
		return
	}

	resultCards, err := cD.cardUsecase.UpdateCardsDayOrder(cards, scheduleID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
		Response: resultCards,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) UpdateCardsLessonOrder(c *gin.Context) {
	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error(err)
		return
	}

	var cards []*models.CardLesson
	err = c.ShouldBindJSON(cards)
	if err != nil {
		log.Error(err)
		return
	}

	resultCards, err := cD.cardUsecase.UpdateCardsLessonOrder(cards, scheduleID)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
		Response: resultCards,
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) DeleteCardDay(c *gin.Context) {
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

	err = cD.cardUsecase.DeleteCardDay(scheduleID, cardID)
	if err != nil {
		log.Error()
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
		Response: "OK",
	}

	c.JSON(http.StatusOK, &response)
}

func (cD *CardDelivery) DeleteCardLesson(c *gin.Context) {
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

	err = cD.cardUsecase.DeleteCardLesson(scheduleID, cardID)
	if err != nil {
		log.Error()
		return
	}

	response := &models.Response{
		Status:   http.StatusOK,
		Response: "OK",
	}

	c.JSON(http.StatusOK, &response)
}