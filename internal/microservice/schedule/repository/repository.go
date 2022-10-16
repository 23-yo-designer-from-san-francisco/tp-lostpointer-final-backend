package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"

	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:schedule:repository:"

const (
	createScheduleQuery = `insert into "schedule_day" (name, day, child_id) values ($1, $2, $3) 
		returning id, name, day;`
	getSchedulesDay = `select id, name, day, child_id from "schedule_day" where child_id = $1;`
	getScheduleDay = `select id, name, day, child_id from "schedule_day" where child_id = $1 and id = $2;`
	getMentorIDFromScheduleIDQuery = `select mentor.id from "mentor" join "child" on child.mentor_id = mentor.id
		join "schedule_day" on schedule_day.child_id = child.id where schedule_day.id = $1;`
)

type scheduleRepository struct {
	db *sqlx.DB
}

func NewScheduleRepository(db *sqlx.DB) *scheduleRepository {
	return &scheduleRepository{
		db: db,
	}
}

func(schR *scheduleRepository) CreateScheduleDay(schedule *models.ScheduleDay) (*models.ScheduleDay, error) {
	var resultSchedule models.ScheduleDay
	
	err := schR.db.QueryRowx(createScheduleQuery, &schedule.Name, &schedule.Day, &schedule.Child_ID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &resultSchedule, nil
}

func (schR *scheduleRepository) GetSchedulesDay(childID int) ([]*models.ScheduleDay, error) {
	message := logMessage + "GetSchedulesDay:"
	log.Debug(message + "started")

	resultSchedules := []*models.ScheduleDay{}
	err := schR.db.Select(&resultSchedules, getSchedulesDay, childID)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return resultSchedules, nil
}

func (schR *scheduleRepository) GetScheduleDay(childID, scheduleID int) (*models.ScheduleDay, error) {
	message := logMessage + "GetScheduleDay:"
	log.Debug(message + "started")

	resultSchedule := models.ScheduleDay{}
	err := schR.db.Get(&resultSchedule, getScheduleDay, childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return &resultSchedule, nil
}

func (schR *scheduleRepository) GetMentorIdFromScheduleID(schedule_id int) (int, error) {
	message := logMessage + "GetMentorIdFromScheduleID:"
	log.Debug(message + "started")

	var mentorID int
	err := schR.db.Get(&mentorID, getMentorIDFromScheduleIDQuery, schedule_id)
	if err != nil {
		log.Error(message + "err = ", err)
		return 0, err
	}
	return mentorID, nil

}