package schedule

import (
	"autfinal/internal/models"
)

type Repository interface {
	CreateScheduleDay(*models.ScheduleDay) (*models.ScheduleDay, error)
	GetSchedulesDay(int) ([]*models.ScheduleDay, error)
	GetScheduleDay(int,int) (*models.ScheduleDay, error)
	GetMentorIdFromScheduleID(int) (int, error)
}