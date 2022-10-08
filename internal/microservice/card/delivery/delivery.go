package delivery

import (
	"autfinal/internal/microservice/card"
	"autfinal/internal/models"
	"autfinal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CardDelivery struct {
	cardUsecase card.Usecase
}

func NewCardDelivery(cardU card.Usecase) *CardDelivery {
	return &CardDelivery{
		cardUsecase: cardU,
	}
}

func(cD *CardDelivery) CreateCard(c *gin.Context) {
	card := c.MustGet("card").(models.Card)

	imgUrl, err := utils.SaveImageFromRequest(c,"image")
	if err != nil {
		c.Error(err)
		return
	}
	if err == nil {
		card.ImgUrl = imgUrl
	}

	resultCard, err := cD.cardUsecase.CreateCard(&card)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, &resultCard)
}

func (cD *CardDelivery) GetCards(c *gin.Context) {
	resultCards, err := cD.cardUsecase.GetCards()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, &resultCards)
}