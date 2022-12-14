package delivery

import (
	"autfinal/internal/microservice/child"
	"autfinal/internal/models"
	log "autfinal/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const logMessage = "microservice:child:delivery:"

type ChildDelivery struct {
	childUsecase child.Usecase
}

func NewChildDelivery(childU child.Usecase) *ChildDelivery {
	return &ChildDelivery{
		childUsecase: childU,
	}
}

func (cD *ChildDelivery) CreateChild(c *gin.Context) {
	requestChild := &models.Child{}
	err := c.ShouldBindJSON(requestChild)
	if err != nil {
		c.Error(err)
		return
	}

	resultMentor, err := cD.childUsecase.CreateChild(requestChild)
	if err != nil {
		c.Error(err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultMentor,
	}

	c.JSON(http.StatusOK, response)
}

func (cD *ChildDelivery) GetChild(c *gin.Context) {
	message := logMessage + "GetChild:"
	log.Debug(message + "started")

	idStr := c.Param("child_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	resultChild, err := cD.childUsecase.GetChild(id)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultChild,
	}
	c.JSON(http.StatusOK,response)
}

func (cD *ChildDelivery) GetChilds(c *gin.Context) {
	message := logMessage + "GetChilds:"
	log.Debug(message + "started")

	resultChilds, err := cD.childUsecase.GetChilds()
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultChilds,
	}
	c.JSON(http.StatusOK,response)
}

func (cD *ChildDelivery) UpdateChild(c *gin.Context) {
	message := logMessage + "UpdateChild:"
	log.Debug(message + "started")

	idStr := c.Param("child_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestChild := &models.Child{}
	err = c.ShouldBindJSON(requestChild)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	requestChild.ID = id

	resultChild, err := cD.childUsecase.UpdateChild(requestChild)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: resultChild,
	}
	c.JSON(http.StatusOK, response)
}

func (cD *ChildDelivery) DeleteChild(c *gin.Context) {
	message := logMessage + "DeleteChild:"
	log.Debug(message + "started")

	idStr := c.Param("child_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	err = cD.childUsecase.DeleteChild(id)
	if err != nil {
		log.Error(message + "err = ", err)
		return
	}

	response := &models.Response{
		Status: http.StatusOK,
		Response: "OK",
	}
	c.JSON(http.StatusOK, response)
}