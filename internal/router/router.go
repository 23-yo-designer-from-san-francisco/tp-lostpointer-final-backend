package router

import (
	cardD "autfinal/internal/microservice/card/delivery"
	mentorD "autfinal/internal/microservice/mentor/delivery"
	childD "autfinal/internal/microservice/child/delivery"
	scheduleD "autfinal/internal/microservice/schedule/delivery"

	"github.com/gin-gonic/gin"
)

func MentorEndpoints(r *gin.RouterGroup, mD *mentorD.MentorDelivery) {
	r.POST("", mD.CreateMentor)
	r.POST("/profile", mD.UpdateMentor)
	r.GET("/profile/:mentor_id", mD.GetMentor)
	r.GET("", mD.GetMentors)
	r.POST("/")
}

func ChildEndpoints(r *gin.RouterGroup, cD *childD.ChildDelivery) {
	r.POST("", cD.CreateChild)
	r.GET("", cD.GetChilds)
	r.GET("/:child_id", cD.GetChild)
	r.POST("/:child_id", cD.UpdateChild)
}

func ScheduleEndpoints(r *gin.RouterGroup, schD *scheduleD.ScheduleDelivery) {
	r.POST("", schD.CreateScheduleDay)
	r.GET("", schD.GetSchedulesDay)
	r.GET("/:schedule_id", schD.GetScheduleDay)
	r.POST("/:schedule_id", schD.UpdateScheduleDay)
	r.POST("/:schedule_id/favourite", schD.MakeFavouriteScheduleDay)
}

func CardDayEndpoints(r *gin.RouterGroup, cD *cardD.CardDelivery) {
	r.POST("", cD.CreateCardDay)
	r.GET("", cD.GetCardsDay)
	r.GET("/:card_id", cD.GetCardDay)
	r.POST("/:card_id/info", cD.UpdateCardDay)
	r.POST("/order", cD.UpdateCardsOrder)
}