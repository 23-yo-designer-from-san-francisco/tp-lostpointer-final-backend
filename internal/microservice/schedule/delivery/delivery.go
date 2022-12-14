package delivery

import (
	"autfinal/internal/microservice/schedule"
	"autfinal/internal/models"
	"net/http"
	"strconv"

	log "autfinal/pkg/logger"

	"github.com/gin-gonic/gin"
)

const logMessage = "microservice:schedule:delivery:"

type ScheduleDelivery struct {
	scheduleUsecase schedule.Usecase
}

func NewScheduleDelivery(scheduleU schedule.Usecase) *ScheduleDelivery {
	return &ScheduleDelivery{
		scheduleUsecase: scheduleU,
	}
}

func(schD *ScheduleDelivery) CreateScheduleDay(c *gin.Context) {
	message := logMessage + "CreateScheduleDay:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule := &models.ScheduleDay{}
	err = c.ShouldBindJSON(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule.Child_ID = childID

	resultSchedule, err := schD.scheduleUsecase.CreateScheduleDay(requestSchedule)
	if err != nil {
		c.Error(err)
		return 
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedule,
	}

	c.JSON(http.StatusOK,response)
}

func(schD *ScheduleDelivery) CreateScheduleLesson(c *gin.Context) {
	message := logMessage + "CreateScheduleLesson:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule := &models.ScheduleLesson{}
	err = c.ShouldBindJSON(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule.Child_ID = childID

	resultSchedule, err := schD.scheduleUsecase.CreateScheduleLesson(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return 
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedule,
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) GetSchedulesDay (c *gin.Context) {
	message := logMessage + "GetSchedulesDay:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultSchedules, err := schD.scheduleUsecase.GetSchedulesDay(childID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedules,
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) GetSchedulesLesson (c *gin.Context) {
	message := logMessage + "GetSchedulesLesson:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultSchedules, err := schD.scheduleUsecase.GetSchedulesLesson(childID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedules,
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) GetScheduleDay(c *gin.Context) {
	message := logMessage + "GetScheduleDay:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	scheduleIDStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultSchedule, err := schD.scheduleUsecase.GetScheduleDay(childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedule,
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) GetScheduleLesson(c *gin.Context) {
	message := logMessage + "GetScheduleLesson:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	scheduleIDStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultSchedule, err := schD.scheduleUsecase.GetScheduleLesson(childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedule,
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) UpdateScheduleDay(c *gin.Context) {
	message := logMessage + "UpdateScheduleDay:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	scheduleIDStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule := &models.ScheduleDay{}
	err = c.ShouldBindJSON(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultSchedule, err := schD.scheduleUsecase.UpdateScheduleDay(requestSchedule, childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedule,
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) UpdateScheduleLesson(c *gin.Context) {
	message := logMessage + "UpdateScheduleLesson:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	scheduleIDStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule := &models.ScheduleLesson{}
	err = c.ShouldBindJSON(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultSchedule, err := schD.scheduleUsecase.UpdateScheduleLesson(requestSchedule, childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedule,
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) MakeFavouriteScheduleDay(c *gin.Context) {
	message := logMessage + "MakeFavouriteScheduleDay:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	scheduleIDStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule := &models.ScheduleDay{}
	err = c.ShouldBindJSON(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule.ID = scheduleID

	resultSchedule, err := schD.scheduleUsecase.MakeFavouriteScheduleDay(requestSchedule, childID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedule,
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) MakeFavouriteScheduleLesson(c *gin.Context) {
	message := logMessage + "MakeFavouriteScheduleLesson:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	scheduleIDStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule := &models.ScheduleLesson{}
	err = c.ShouldBindJSON(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule.ID = scheduleID

	resultSchedule, err := schD.scheduleUsecase.MakeFavouriteScheduleLesson(requestSchedule, childID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultSchedule,
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) DeleteScheduleDay(c *gin.Context) {
	message := logMessage + "DeleteScheduleDay:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	scheduleIDStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	err = schD.scheduleUsecase.DeleteScheduleDay(childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: "OK",
	}

	c.JSON(http.StatusOK,response)
}

func (schD *ScheduleDelivery) DeleteScheduleLesson(c *gin.Context) {
	message := logMessage + "DeleteScheduleLesson:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	scheduleIDStr := c.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	err = schD.scheduleUsecase.DeleteScheduleLesson(childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: "OK",
	}

	c.JSON(http.StatusOK,response)
}