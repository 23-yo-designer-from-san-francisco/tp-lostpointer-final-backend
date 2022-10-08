package delivery

import "github.com/gin-gonic/gin"

type ScheduleDelivery struct {
	cardUsecase schedule.Usecase
}

func NewScheduleDelivery(cardU schedule.Usecase) *ScheduleDelivery {
	return &ScheduleDelivery{
		cardUsecase: cardU,
	}
}

func(sD *ScheduleDelivery) CreateSchedule(c *gin.Context) {
	
}