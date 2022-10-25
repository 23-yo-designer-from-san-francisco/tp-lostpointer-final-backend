package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"

	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:schedule:repository:"

const (
	createScheduleBeforeAfterQuery = `insert into "schedule_before_after" (name, child_id) values ($1, $2) 
		returning id, name, favourite;`

	getSchedulesBeforeAfterQuery = `select id, name, favourite, child_id, cards_count from "schedule_before_after" where child_id = $1 and deletedAt is null;`

	getScheduleBeforeAfterQuery = `select id, name, favourite, child_id from "schedule_before_after" where child_id = $1 and id = $2;`

	getMentorIDFromScheduleBeforeAfterIDQuery = `select mentor.id from "mentor" join "child" on child.mentor_id = mentor.id
		join "schedule_before_after" on schedule_before_after.child_id = child.id where schedule_before_after.id = $1;`

	updateScheduleBeforeAfterQuery = `update "schedule_before_after" set name = $1, updatedAt = now() where child_id = $2 and id = $3
		returning id, name, favourite, child_id;`

	removeFavouriteScheduleBeforeAfterQuery = `update "schedule_before_after" set favourite = false where favourite = true;`

	makeFavouriteScheduleBeforeAfterQuery = `update "schedule_before_after" set favourite = true where child_id = $1 and id = $2
		returning id, name, favourite, child_id;`

	safeDeleteScheduleBeforeAfterQuery = `update "schedule_before_after" set deletedAt = now() where child_id = $1 and id = $2;`
)

type scheduleRepository struct {
	db *sqlx.DB
}

func NewScheduleRepository(db *sqlx.DB) *scheduleRepository {
	return &scheduleRepository{
		db: db,
	}
}

func(schR *scheduleRepository) CreateScheduleBeforeAfter(schedule *models.ScheduleBeforeAfter) (*models.ScheduleBeforeAfter, error) {
	message := logMessage + "CreateScheduleBeforeAfter:"
	log.Debug(message + "started")
	var resultSchedule models.ScheduleBeforeAfter
	
	err := schR.db.QueryRowx(createScheduleBeforeAfterQuery, &schedule.Name, &schedule.Child_ID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return &resultSchedule, nil
}

func (schR *scheduleRepository) GetSchedulesBeforeAfter(childID int) ([]*models.ScheduleBeforeAfter, error) {
	message := logMessage + "GetSchedulesBeforeAfter:"
	log.Debug(message + "started")

	resultSchedules := []*models.ScheduleBeforeAfter{}
	err := schR.db.Select(&resultSchedules, getSchedulesBeforeAfterQuery, childID)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return resultSchedules, nil
}

func (schR *scheduleRepository) GetMentorIdFromScheduleBeforeAfterID(schedule_id int) (int, error) {
	message := logMessage + "GetMentorIdFromScheduleBeforeAfterID:"
	log.Debug(message + "started")

	var mentorID int
	err := schR.db.Get(&mentorID, getMentorIDFromScheduleBeforeAfterIDQuery, schedule_id)
	if err != nil {
		log.Error(message + "err = ", err)
		return 0, err
	}
	return mentorID, nil
}

func (schR *scheduleRepository) GetScheduleBeforeAfter(childID, scheduleID int) (*models.ScheduleBeforeAfter, error) {
	message := logMessage + "GetScheduleBeforeAfter:"
	log.Debug(message + "started")

	resultSchedule := models.ScheduleBeforeAfter{}
	err := schR.db.Get(&resultSchedule, getScheduleBeforeAfterQuery, childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return &resultSchedule, nil
}

func (schR *scheduleRepository) MakeFavouriteScheduleBeforeAfter(schedule *models.ScheduleBeforeAfter, childID int) (*models.ScheduleBeforeAfter, error) {
	message := logMessage + "MakeFavouriteScheduleBeforeAfter:"
	log.Debug(message + "started")

	tx, err := schR.db.Beginx()
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	_, err = tx.Exec(removeFavouriteScheduleBeforeAfterQuery)
	if err != nil {
		log.Error(message + "err = ", err)
		tx.Rollback()
		return nil, err
	}

	resultSchedule := models.ScheduleBeforeAfter{} 
	err = tx.QueryRowx(makeFavouriteScheduleBeforeAfterQuery, childID, schedule.ID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &resultSchedule, nil
}

func (schR *scheduleRepository) UpdateScheduleBeforeAfter(schedule *models.ScheduleBeforeAfter, childID, scheduleID int) (*models.ScheduleBeforeAfter, error) {
	message := logMessage + "UpdateScheduleeforeAfter:"
	log.Debug(message + "started")

	resultSchedule := models.ScheduleBeforeAfter{}
	err := schR.db.QueryRowx(updateScheduleBeforeAfterQuery, &schedule.Name, &childID, &scheduleID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return &resultSchedule, nil
}


func (schR *scheduleRepository) DeleteScheduleBeforeAfter(childID, scheduleID int) error {
	message := logMessage + "DeleteScheduleBeforeAfter:"
	log.Debug(message + "started")

	_, err := schR.db.Exec(safeDeleteScheduleBeforeAfterQuery, childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return err
	}

	return nil
}

