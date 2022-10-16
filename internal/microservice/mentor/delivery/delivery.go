package delivery

import (
	"autfinal/internal/microservice/mentor"
	"autfinal/internal/models"
	"net/http"
	"strconv"
	log "autfinal/pkg/logger"

	"github.com/gin-gonic/gin"
)

const logMessage = "microservice:mentor:delivery:"

type MentorDelivery struct {
	mentorUsecase mentor.Usecase
}

func NewMentorDelivery(mentorU mentor.Usecase) *MentorDelivery {
	return &MentorDelivery{
		mentorUsecase: mentorU,
	}
}

func (mD *MentorDelivery) CreateMentor(c *gin.Context) {
	message := logMessage + "CreateMentor:"
	log.Debug(message + "started")

	requestMentor := &models.Mentor{}
	err := c.ShouldBindJSON(requestMentor)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultMentor, err := mD.mentorUsecase.CreateMentor(requestMentor)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultMentor,
	}

	c.JSON(http.StatusOK, response)
}

func (mD *MentorDelivery) UpdateMentor(c *gin.Context) {
	message := logMessage + "UpdateMentor:"
	log.Debug(message + "started")

	requestMentor := &models.Mentor{}
	err := c.ShouldBindJSON(requestMentor)
	if err != nil {
		c.Error(err)
		return
	}

	resultMentor, err := mD.mentorUsecase.UpdateMentor(requestMentor)
	if err != nil {
		log.Error(message + "err = ", err)
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultMentor,
	}
	c.JSON(http.StatusOK,response)
}

func (mD *MentorDelivery) GetMentor(c *gin.Context) {
	message := logMessage + "GetMentor:"
	log.Debug(message + "started")

	idStr := c.Param("mentor_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultMentor, err := mD.mentorUsecase.GetMentor(id)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultMentor,
	}
	c.JSON(http.StatusOK,response)
}

func (mD *MentorDelivery) GetMentors(c *gin.Context) {
	message := logMessage + "GetMentors:"
	log.Debug(message + "started")
	resultMentors, err := mD.mentorUsecase.GetMentors()
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultMentors,
	}
	c.JSON(http.StatusOK,response)
}