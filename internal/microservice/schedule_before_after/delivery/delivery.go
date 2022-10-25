package delivery

import (
	"autfinal/internal/microservice/schedule_before_after"
	"autfinal/internal/models"
	"net/http"
	"strconv"

	log "autfinal/pkg/logger"

	"github.com/gin-gonic/gin"
)

const logMessage = "microservice:schedule_before_after:delivery:"

type ScheduleBeforeAfterDelivery struct {
	scheduleUsecase scheduleBeforeAfter.Usecase
}

func NewScheduleBeforeAfterDelivery(scheduleU scheduleBeforeAfter.Usecase) *ScheduleBeforeAfterDelivery {
	return &ScheduleBeforeAfterDelivery{
		scheduleUsecase: scheduleU,
	}
}

func (schbaD *ScheduleBeforeAfterDelivery) CreateScheduleBeforeAfter(c *gin.Context) {
	message := logMessage + "CreateScheduleBeforeAfter:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule := &models.ScheduleBeforeAfter{}
	err = c.ShouldBindJSON(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule.Child_ID = childID

	resultSchedule, err := schbaD.scheduleUsecase.CreateScheduleBeforeAfter(requestSchedule)
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

func (schbaD *ScheduleBeforeAfterDelivery) GetSchedulesBeforeAfter (c *gin.Context) {
	message := logMessage + "GetSchedulesBeforeAfter:"
	log.Debug(message + "started")

	childIDStr := c.Param("child_id")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultSchedules, err := schbaD.scheduleUsecase.GetSchedulesBeforeAfter(childID)
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

func (schbaD *ScheduleBeforeAfterDelivery) GetScheduleBeforeAfter(c *gin.Context) {
	message := logMessage + "GetScheduleBeforeAfter:"
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

	resultSchedule, err := schbaD.scheduleUsecase.GetScheduleBeforeAfter(childID, scheduleID)
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

func (schbaD *ScheduleBeforeAfterDelivery) UpdateScheduleBeforeAfter(c *gin.Context) {
	message := logMessage + "UpdateScheduleBeforeAfter:"
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

	requestSchedule := &models.ScheduleBeforeAfter{}
	err = c.ShouldBindJSON(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultSchedule, err := schbaD.scheduleUsecase.UpdateScheduleBeforeAfter(requestSchedule, childID, scheduleID)
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

func (schbaD *ScheduleBeforeAfterDelivery) MakeFavouriteScheduleBeforeAfter(c *gin.Context) {
	message := logMessage + "MakeFavouriteScheduleBeforeAfter:"
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

	requestSchedule := &models.ScheduleBeforeAfter{}
	err = c.ShouldBindJSON(requestSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestSchedule.ID = scheduleID

	resultSchedule, err := schbaD.scheduleUsecase.MakeFavouriteScheduleBeforeAfter(requestSchedule, childID)
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

func (schbaD *ScheduleBeforeAfterDelivery) DeleteScheduleBeforeAfter(c *gin.Context) {
	message := logMessage + "DeleteScheduleBeforeAfter:"
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

	err = schbaD.scheduleUsecase.DeleteScheduleBeforeAfter(childID, scheduleID)
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