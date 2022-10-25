package delivery

import (
	"autfinal/internal/microservice/card_before_after"
	"autfinal/internal/models"
	"autfinal/internal/utils"
	log "autfinal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strconv"
)

const logMessage = "microservice:card_before_after:delivery:"

type CardBeforeAfterDelivery struct {
	cardUsecase cardBeforeAfter.Usecase
}

func NewCardBeforeAfterDelivery(cardU cardBeforeAfter.Usecase) *CardBeforeAfterDelivery {
	return &CardBeforeAfterDelivery{
		cardUsecase: cardU,
	}
}

func (cbaD *CardBeforeAfterDelivery) CreateCardBeforeAfter(c *gin.Context) {
	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error(err)
		return
	}

	cardJsonStr := c.Request.FormValue("card")
	var card = &models.CardBeforeAfter{}
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

	resultCard, err := cbaD.cardUsecase.CreateCardBeforeAfter(card, scheduleID)
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

func (cbaD *CardBeforeAfterDelivery) GetCardsBeforeAfter(c *gin.Context) {
	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error(err)
		return
	}

	resultCards, err := cbaD.cardUsecase.GetCardsBeforeAfter(scheduleID)
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

func (cbaD *CardBeforeAfterDelivery) GetCardBeforeAfter(c *gin.Context) {
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

	resultCard, err := cbaD.cardUsecase.GetCardBeforeAfter(scheduleID, cardID)
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

func (cbaD *CardBeforeAfterDelivery) UpdateCardBeforeAfter(c *gin.Context) {
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
	var card = &models.CardBeforeAfter{}
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

	resultCard, err := cbaD.cardUsecase.UpdateCardBeforeAfter(card, scheduleID, cardID)
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

func (cbaD *CardBeforeAfterDelivery) UpdateCardsBeforeAfterOrder(c *gin.Context) {
	scheduleIdStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		log.Error(err)
		return
	}

	var cards []*models.CardBeforeAfter
	err = c.ShouldBindJSON(cards)
	if err != nil {
		log.Error(err)
		return
	}

	resultCards, err := cbaD.cardUsecase.UpdateCardsBeforeAfterOrder(cards, scheduleID)
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

func (cbaD *CardBeforeAfterDelivery) DeleteCardBeforeAfter(c *gin.Context) {
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

	err = cbaD.cardUsecase.DeleteCardBeforeAfter(scheduleID, cardID)
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