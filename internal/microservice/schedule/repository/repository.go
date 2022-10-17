package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"

	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:schedule:repository:"

const (
	createScheduleQuery = `insert into "schedule_day" (name, day, child_id) values ($1, $2, $3) 
		returning id, name, day, favourite;`
	getSchedulesDay = `select id, name, day, favourite, child_id from "schedule_day" where child_id = $1;`
	getScheduleDay = `select id, name, day, favourite, child_id from "schedule_day" where child_id = $1 and id = $2;`
	getMentorIDFromScheduleIDQuery = `select mentor.id from "mentor" join "child" on child.mentor_id = mentor.id
		join "schedule_day" on schedule_day.child_id = child.id where schedule_day.id = $1;`
	updateScheduleDay = `update "schedule_day" set name = $1, day = $2 where child_id = $3 and id = $4
		returning id, name, day, favourite, child_id;`
	removeFavouriteScheduleDay = `update "schedule_day" set favourite = false where favourite = true;`
	makeFavouriteScheduleDay = `update "schedule_day" set favourite = true where child_id = $1 and id = $2
		returning id, name, day, favourite, child_id;`
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

func (schR *scheduleRepository) UpdateScheduleDay(schedule *models.ScheduleDay, childID, scheduleID int) (*models.ScheduleDay, error) {
	message := logMessage + "UpdateScheduleDay:"
	log.Debug(message + "started")

	resultSchedule := models.ScheduleDay{}
	err := schR.db.QueryRowx(updateScheduleDay, &schedule.Name, &schedule.Day, &childID, &scheduleID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return &resultSchedule, nil
}

func (schR *scheduleRepository) MakeFavouriteScheduleDay(schedule *models.ScheduleDay, childID int) (*models.ScheduleDay, error) {
	message := logMessage + "MakeFavouriteScheduleDay:"
	log.Debug(message + "started")

	tx, err := schR.db.Beginx()
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	_, err = tx.Exec(removeFavouriteScheduleDay)
	if err != nil {
		log.Error(message + "err = ", err)
		tx.Rollback()
		return nil, err
	}

	resultSchedule := models.ScheduleDay{} 
	err = tx.QueryRowx(makeFavouriteScheduleDay, childID, schedule.ID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &resultSchedule, nil
}