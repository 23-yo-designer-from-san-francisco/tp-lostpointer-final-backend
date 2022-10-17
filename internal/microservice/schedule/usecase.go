package schedule

import (
	"autfinal/internal/models"
)

type Usecase interface {
	CreateScheduleDay(*models.ScheduleDay) (*models.ScheduleDay, error)
	GetSchedulesDay(int) ([]*models.ScheduleDay, error)
	GetScheduleDay(int,int) (*models.ScheduleDay, error)
	UpdateScheduleDay(*models.ScheduleDay, int, int) (*models.ScheduleDay, error)
	MakeFavouriteScheduleDay(*models.ScheduleDay, int) (*models.ScheduleDay, error)
}