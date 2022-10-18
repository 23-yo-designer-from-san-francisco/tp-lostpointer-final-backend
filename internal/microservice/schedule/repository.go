package schedule

import (
	"autfinal/internal/models"
)

type Repository interface {
	CreateScheduleDay(*models.ScheduleDay) (*models.ScheduleDay, error)
	GetSchedulesDay(int) ([]*models.ScheduleDay, error)
	GetScheduleDay(int,int) (*models.ScheduleDay, error)
	GetMentorIdFromScheduleDayID(int) (int, error)
	UpdateScheduleDay(*models.ScheduleDay, int, int) (*models.ScheduleDay, error)
	MakeFavouriteScheduleDay(*models.ScheduleDay, int) (*models.ScheduleDay, error)
	DeleteScheduleDay(int, int) (error)

	CreateScheduleLesson(*models.ScheduleLesson) (*models.ScheduleLesson, error)
	GetMentorIdFromScheduleLessonID(int) (int, error)
	GetSchedulesLesson(int) ([]*models.ScheduleLesson, error)
	GetScheduleLesson(int,int) (*models.ScheduleLesson, error)
	UpdateScheduleLesson(*models.ScheduleLesson, int, int) (*models.ScheduleLesson, error)
	MakeFavouriteScheduleLesson(*models.ScheduleLesson, int) (*models.ScheduleLesson, error)
	DeleteScheduleLesson(int, int) (error)
}