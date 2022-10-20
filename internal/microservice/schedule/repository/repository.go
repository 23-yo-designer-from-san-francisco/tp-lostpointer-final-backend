package repository

import (
	"autfinal/internal/models"
	log "autfinal/pkg/logger"

	"github.com/jmoiron/sqlx"
)

const logMessage = "microservice:schedule:repository:"

const (
	createScheduleDayQuery = `insert into "schedule_day" (name, day, child_id) values ($1, $2, $3) 
		returning id, name, day, favourite;`

	getSchedulesDayQuery = `select id, name, day, favourite, child_id, cards_count from "schedule_day" where child_id = $1 and deletedAt is null;`

	getScheduleDayQuery = `select id, name, day, favourite, child_id from "schedule_day" where child_id = $1 and id = $2;`

	getMentorIDFromScheduleDayIDQuery = `select mentor.id from "mentor" join "child" on child.mentor_id = mentor.id
		join "schedule_day" on schedule_day.child_id = child.id where schedule_day.id = $1;`

	updateScheduleDayQuery = `update "schedule_day" set name = $1, day = $2, updatedAt = now() where child_id = $3 and id = $4
		returning id, name, day, favourite, child_id;`

	removeFavouriteScheduleDayQuery = `update "schedule_day" set favourite = false where favourite = true;`

	makeFavouriteScheduleDayQuery = `update "schedule_day" set favourite = true where child_id = $1 and id = $2
		returning id, name, day, favourite, child_id;`

	safeDeleteScheduleDayQuery = `update "schedule_day" set deletedAt = now() where child_id = $1 and id = $2;`


	createScheduleLessonQuery = `insert into "schedule_lesson" (name, duration, child_id) values ($1, $2, $3) 
		returning id, name, duration, favourite, cards_count;`

	getMentorIDFromScheduleLessonIDQuery = `select mentor.id from "mentor" join "child" on child.mentor_id = mentor.id
		join "schedule_lesson" on schedule_lesson.child_id = child.id where schedule_lesson.id = $1;`

	getSchedulesLessonQuery = `select id, name, duration, favourite, child_id, cards_count from "schedule_lesson" where child_id = $1 and deletedAt is null;`

	getScheduleLessonQuery = `select id, name, duration, favourite, child_id from "schedule_lesson" where child_id = $1 and id = $2;`

	updateScheduleLessonQuery = `update "schedule_lesson" set name = $1, duration = $2, updatedAt = now() where child_id = $3 and id = $4
		returning id, name, duration, favourite, child_id;`

	removeFavouriteScheduleLessonQuery = `update "schedule_lesson" set favourite = false where favourite = true;`
	
	makeFavouriteScheduleLessonQuery = `update "schedule_lesson" set favourite = true where child_id = $1 and id = $2
		returning id, name, duration, favourite, child_id;`
	
	safeDeleteScheduleLessonQuery = `update "schedule_lesson" set deletedAt = now() where child_id = $1 and id = $2;`
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
	message := logMessage + "CreateScheduleDay:"
	log.Debug(message + "started")
	var resultSchedule models.ScheduleDay
	
	err := schR.db.QueryRowx(createScheduleDayQuery, &schedule.Name, &schedule.Day, &schedule.Child_ID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return &resultSchedule, nil
}

func(schR *scheduleRepository) CreateScheduleLesson(schedule *models.ScheduleLesson) (*models.ScheduleLesson, error) {
	message := logMessage + "CreateScheduleLesson:"
	log.Debug(message + "started")
	var resultSchedule models.ScheduleLesson
	
	err := schR.db.QueryRowx(createScheduleLessonQuery, &schedule.Name, &schedule.Duration, &schedule.Child_ID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return &resultSchedule, nil
}

func (schR *scheduleRepository) GetSchedulesDay(childID int) ([]*models.ScheduleDay, error) {
	message := logMessage + "GetSchedulesDay:"
	log.Debug(message + "started")

	resultSchedules := []*models.ScheduleDay{}
	err := schR.db.Select(&resultSchedules, getSchedulesDayQuery, childID)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return resultSchedules, nil
}

func (schR *scheduleRepository) GetSchedulesLesson(childID int) ([]*models.ScheduleLesson, error) {
	message := logMessage + "GetSchedulesLesson:"
	log.Debug(message + "started")

	resultSchedules := []*models.ScheduleLesson{}
	err := schR.db.Select(&resultSchedules, getSchedulesLessonQuery, childID)
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
	err := schR.db.Get(&resultSchedule, getScheduleDayQuery, childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return &resultSchedule, nil
}

func (schR *scheduleRepository) GetScheduleLesson(childID, scheduleID int) (*models.ScheduleLesson, error) {
	message := logMessage + "GetScheduleLesson:"
	log.Debug(message + "started")

	resultSchedule := models.ScheduleLesson{}
	err := schR.db.Get(&resultSchedule, getScheduleLessonQuery, childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}
	return &resultSchedule, nil
}

func (schR *scheduleRepository) GetMentorIdFromScheduleDayID(schedule_id int) (int, error) {
	message := logMessage + "GetMentorIdFromScheduleDayID:"
	log.Debug(message + "started")

	var mentorID int
	err := schR.db.Get(&mentorID, getMentorIDFromScheduleDayIDQuery, schedule_id)
	if err != nil {
		log.Error(message + "err = ", err)
		return 0, err
	}
	return mentorID, nil
}

func (schR *scheduleRepository) GetMentorIdFromScheduleLessonID(schedule_id int) (int, error) {
	message := logMessage + "GetMentorIdFromScheduleLessonID:"
	log.Debug(message + "started")

	var mentorID int
	err := schR.db.Get(&mentorID, getMentorIDFromScheduleLessonIDQuery, schedule_id)
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
	err := schR.db.QueryRowx(updateScheduleDayQuery, &schedule.Name, &schedule.Day, &childID, &scheduleID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	return &resultSchedule, nil
}

func (schR *scheduleRepository) UpdateScheduleLesson(schedule *models.ScheduleLesson, childID, scheduleID int) (*models.ScheduleLesson, error) {
	message := logMessage + "UpdateScheduleLesson:"
	log.Debug(message + "started")

	resultSchedule := models.ScheduleLesson{}
	err := schR.db.QueryRowx(updateScheduleLessonQuery, &schedule.Name, &schedule.Duration, &childID, &scheduleID).StructScan(&resultSchedule)
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

	_, err = tx.Exec(removeFavouriteScheduleDayQuery)
	if err != nil {
		log.Error(message + "err = ", err)
		tx.Rollback()
		return nil, err
	}

	resultSchedule := models.ScheduleDay{} 
	err = tx.QueryRowx(makeFavouriteScheduleDayQuery, childID, schedule.ID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &resultSchedule, nil
}

func (schR *scheduleRepository) MakeFavouriteScheduleLesson(schedule *models.ScheduleLesson, childID int) (*models.ScheduleLesson, error) {
	message := logMessage + "MakeFavouriteScheduleDay:"
	log.Debug(message + "started")

	tx, err := schR.db.Beginx()
	if err != nil {
		log.Error(message + "err = ", err)
		return nil, err
	}

	_, err = tx.Exec(removeFavouriteScheduleLessonQuery)
	if err != nil {
		log.Error(message + "err = ", err)
		tx.Rollback()
		return nil, err
	}

	resultSchedule := models.ScheduleLesson{} 
	err = tx.QueryRowx(makeFavouriteScheduleLessonQuery, childID, schedule.ID).StructScan(&resultSchedule)
	if err != nil {
		log.Error(message + "err = ", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &resultSchedule, nil
}

func (schR *scheduleRepository) DeleteScheduleDay(childID, scheduleID int) error {
	message := logMessage + "DeleteScheduleDay:"
	log.Debug(message + "started")

	_, err := schR.db.Exec(safeDeleteScheduleDayQuery, childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return err
	}

	return nil
}

func (schR *scheduleRepository) DeleteScheduleLesson(childID, scheduleID int) error {
	message := logMessage + "DeleteScheduleLesson:"
	log.Debug(message + "started")

	_, err := schR.db.Exec(safeDeleteScheduleLessonQuery, childID, scheduleID)
	if err != nil {
		log.Error(message + "err = ", err)
		return err
	}

	return nil
}