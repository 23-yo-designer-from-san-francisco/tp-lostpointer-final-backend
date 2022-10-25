package usecase

import (
	"autfinal/internal/microservice/schedule_before_after"
	"autfinal/internal/models"
)

type scheduleBeforeAfterUsecase struct {
	scheduleBeforeAfterRepository scheduleBeforeAfter.Repository
}

func NewScheduleBeforeAfterUsecase(scheduleR scheduleBeforeAfter.Repository) *scheduleBeforeAfterUsecase {
	return &scheduleBeforeAfterUsecase{
		scheduleBeforeAfterRepository: scheduleR,
	}
}

func (schbaU *scheduleBeforeAfterUsecase) CreateScheduleBeforeAfter(schedule *models.ScheduleBeforeAfter) (*models.ScheduleBeforeAfter, error) {
	return schbaU.scheduleBeforeAfterRepository.CreateScheduleBeforeAfter(schedule)
}

func (schbaU *scheduleBeforeAfterUsecase) GetSchedulesBeforeAfter(childID int) ([]*models.ScheduleBeforeAfter, error) {
	return schbaU.scheduleBeforeAfterRepository.GetSchedulesBeforeAfter(childID)
}

func (schbaU *scheduleBeforeAfterUsecase) GetScheduleBeforeAfter(childID, scheduleID int) (*models.ScheduleBeforeAfter, error) {
	return schbaU.scheduleBeforeAfterRepository.GetScheduleBeforeAfter(childID, scheduleID)
}

func (schbaU *scheduleBeforeAfterUsecase) UpdateScheduleBeforeAfter(schedule *models.ScheduleBeforeAfter, childID, scheduleID int) (*models.ScheduleBeforeAfter, error) {
	return schbaU.scheduleBeforeAfterRepository.UpdateScheduleBeforeAfter(schedule, childID, scheduleID)
}

func (schbaU *scheduleBeforeAfterUsecase) MakeFavouriteScheduleBeforeAfter(schedule *models.ScheduleBeforeAfter, childID int) (*models.ScheduleBeforeAfter, error) {
	return schbaU.scheduleBeforeAfterRepository.MakeFavouriteScheduleBeforeAfter(schedule, childID)
}

func (schbaU *scheduleBeforeAfterUsecase) DeleteScheduleBeforeAfter(childID, scheduleID int) (error) {
	return schbaU.scheduleBeforeAfterRepository.DeleteScheduleBeforeAfter(childID, scheduleID)
}