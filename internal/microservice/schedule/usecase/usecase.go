package usecase

import (
	"autfinal/internal/microservice/schedule"
	"autfinal/internal/models"
)

type scheduleUsecase struct {
	scheduleRepository schedule.Repository
}

func NewScheduleUsecase(scheduleR schedule.Repository) *scheduleUsecase {
	return &scheduleUsecase{
		scheduleRepository: scheduleR,
	}
}

func (schU *scheduleUsecase) CreateScheduleDay(schedule *models.ScheduleDay) (*models.ScheduleDay, error) {
	return schU.scheduleRepository.CreateScheduleDay(schedule)
}

func (schU *scheduleUsecase) GetSchedulesDay(childID int) ([]*models.ScheduleDay, error) {
	return schU.scheduleRepository.GetSchedulesDay(childID)
}

func (schU *scheduleUsecase) GetScheduleDay(childID, scheduleID int) (*models.ScheduleDay, error) {
	return schU.scheduleRepository.GetScheduleDay(childID, scheduleID)
}