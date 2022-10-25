package scheduleBeforeAfter

import (
	"autfinal/internal/models"
)

type Repository interface {
	CreateScheduleBeforeAfter(*models.ScheduleBeforeAfter) (*models.ScheduleBeforeAfter, error)
	GetSchedulesBeforeAfter(int) ([]*models.ScheduleBeforeAfter, error)
	GetScheduleBeforeAfter(int,int) (*models.ScheduleBeforeAfter, error)
	GetMentorIdFromScheduleBeforeAfterID(int) (int, error)
	UpdateScheduleBeforeAfter(*models.ScheduleBeforeAfter, int, int) (*models.ScheduleBeforeAfter, error)
	MakeFavouriteScheduleBeforeAfter(*models.ScheduleBeforeAfter, int) (*models.ScheduleBeforeAfter, error)
	DeleteScheduleBeforeAfter(int, int) (error)
}