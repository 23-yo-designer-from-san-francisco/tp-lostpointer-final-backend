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

func (schU *scheduleUsecase) CreateScheduleLesson(schedule *models.ScheduleLesson) (*models.ScheduleLesson, error) {
	return schU.scheduleRepository.CreateScheduleLesson(schedule)
}

func (schU *scheduleUsecase) GetSchedulesDay(childID int) ([]*models.ScheduleDay, error) {
	return schU.scheduleRepository.GetSchedulesDay(childID)
}

func (schU *scheduleUsecase) GetSchedulesLesson(childID int) ([]*models.ScheduleLesson, error) {
	return schU.scheduleRepository.GetSchedulesLesson(childID)
}

func (schU *scheduleUsecase) GetScheduleDay(childID, scheduleID int) (*models.ScheduleDay, error) {
	return schU.scheduleRepository.GetScheduleDay(childID, scheduleID)
}

func (schU *scheduleUsecase) GetScheduleLesson(childID, scheduleID int) (*models.ScheduleLesson, error) {
	return schU.scheduleRepository.GetScheduleLesson(childID, scheduleID)
}

func (schU *scheduleUsecase) UpdateScheduleDay(schedule *models.ScheduleDay, childID, scheduleID int) (*models.ScheduleDay, error) {
	return schU.scheduleRepository.UpdateScheduleDay(schedule, childID, scheduleID)
}

func (schU *scheduleUsecase) UpdateScheduleLesson(schedule *models.ScheduleLesson, childID, scheduleID int) (*models.ScheduleLesson, error) {
	return schU.scheduleRepository.UpdateScheduleLesson(schedule, childID, scheduleID)
}

func (schU *scheduleUsecase) MakeFavouriteScheduleDay(schedule *models.ScheduleDay, childID int) (*models.ScheduleDay, error) {
	return schU.scheduleRepository.MakeFavouriteScheduleDay(schedule, childID)
}

func (schU *scheduleUsecase) MakeFavouriteScheduleLesson(schedule *models.ScheduleLesson, childID int) (*models.ScheduleLesson, error) {
	return schU.scheduleRepository.MakeFavouriteScheduleLesson(schedule, childID)
}

func (schU *scheduleUsecase) DeleteScheduleDay(childID, scheduleID int) (error) {
	return schU.scheduleRepository.DeleteScheduleDay(childID, scheduleID)
}

func (schU *scheduleUsecase) DeleteScheduleLesson(childID, scheduleID int) (error) {
	return schU.scheduleRepository.DeleteScheduleLesson(childID, scheduleID)
}
