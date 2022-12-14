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
	r.POST("/profile/:mentor_id/goodbye", mD.DeleteMentor)
	r.GET("/images", mD.GetPersonalImages)
}

func ChildEndpoints(r *gin.RouterGroup, cD *childD.ChildDelivery) {
	r.POST("", cD.CreateChild)
	r.GET("", cD.GetChilds)
	r.GET("/:child_id", cD.GetChild)
	r.POST("/:child_id", cD.UpdateChild)
	r.POST("/:child_id/goodbye", cD.DeleteChild)
}

func ScheduleDayEndpoints(r *gin.RouterGroup, schD *scheduleD.ScheduleDelivery) {
	r.POST("", schD.CreateScheduleDay)
	r.GET("", schD.GetSchedulesDay)
	r.GET("/:schedule_id", schD.GetScheduleDay)
	r.POST("/:schedule_id", schD.UpdateScheduleDay)
	r.POST("/:schedule_id/favourite", schD.MakeFavouriteScheduleDay)
	r.POST("/:schedule_id/goodbye", schD.DeleteScheduleDay)
}

func ScheduleLessonEndpoints(r *gin.RouterGroup, schD *scheduleD.ScheduleDelivery) {
	r.POST("", schD.CreateScheduleLesson)
	r.GET("", schD.GetSchedulesLesson)
	r.GET("/:schedule_id", schD.GetScheduleLesson)
	r.POST("/:schedule_id", schD.UpdateScheduleLesson)
	r.POST("/:schedule_id/favourite", schD.MakeFavouriteScheduleLesson)
	r.POST("/:schedule_id/goodbye", schD.DeleteScheduleDay)
}

func CardDayEndpoints(r *gin.RouterGroup, cD *cardD.CardDelivery) {
	r.POST("", cD.CreateCardDay)
	r.GET("", cD.GetCardsDay)
	r.GET("/:card_id", cD.GetCardDay)
	r.POST("/:card_id", cD.UpdateCardDay)
	r.POST("/order", cD.UpdateCardsDayOrder)
	r.POST("/card_id/goodbye", cD.DeleteCardDay)
}

func CardLessonEndpoints(r *gin.RouterGroup, cD *cardD.CardDelivery) {
	r.POST("", cD.CreateCardLesson)
	r.GET("", cD.GetCardsLesson)
	r.GET("/:card_id", cD.GetCardLesson)
	r.POST("/:card_id", cD.UpdateCardLesson)
	r.POST("/order", cD.UpdateCardsLessonOrder)
	r.POST("/:card_id/goodbye", cD.DeleteCardLesson)
}