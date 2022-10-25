package models

type Response struct {
	Status int `json:"status"`
	Response interface{} `json:"response"`
}

type Mentor struct {
	ID int `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`
	Surname string `json:"surname,omitempty" db:"surname"`
	Email string `json:"email,omitempty" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
}

type Child struct {
	ID int `json:"id,omitempty" db:"id"`
	Mentor_ID int `json:"mentor_id,omitempty" db:"mentor_id"`
	Name string `json:"name,omitempty" db:"name"`
	DateOfBirth *string `json:"dateOfBirth,omitempty" db:"date_of_birth"`
}

type ScheduleDay struct {
	ID int `json:"id,omitempty" db:"id"`
	Child_ID int `json:"child_id,omitempty" db:"child_id"`
	Name string `json:"name,omitempty" db:"name"`
	Day string `json:"day,omitempty" db:"day"`
	Favourite *bool `json:"favourite,omitempty" db:"favourite"`
	Cards int `json:"cards,omitempty" db:"cards_count"`
}

type ScheduleLesson struct {
	ID int `json:"id,omitempty" db:"id"`
	Child_ID int `json:"child_id,omitempty" db:"child_id"`
	Name string `json:"name,omitempty" db:"name"`
	Duration *int `json:"duration" db:"duration"`
	Favourite *bool `json:"favourite,omitempty" db:"favourite"`
	Cards int `json:"cards,omitempty" db:"cards_count"`
}

type ScheduleBeforeAfter struct {
	ID int `json:"id,omitempty" db:"id"`
	Child_ID int `json:"child_id,omitempty" db:"child_id"`
	Name string `json:"name,omitempty" db:"name"`
	Favourite *bool `json:"favourite,omitempty" db:"favourite"`
	Cards int `json:"cards,omitempty" db:"cards_count"`
}

type CardDay struct {
	ID int `json:"id,omitempty" db:"id"`
	Schedule_ID int `json:"schedule_id,omitempty" db:"schedule_id"`
	Name *string `json:"name,omitempty" db:"name"`
	Done *bool `json:"done,omitempty" db:"done"`
	ImgUrl string `json:"imgUrl,omitempty" db:"-"`
	ImgUUID string `json:"-" db:"imguuid"`
	Order int `json:"orderPlace,omitempty" db:"orderplace"`
	StartTime *string `json:"startTime,omitempty" db:"starttime"`
	EndTime *string `json:"endTime,omitempty" db:"endtime"`
}

type CardLesson struct {
	ID int `json:"id,omitempty" db:"id"`
	Schedule_ID int `json:"schedule_id,omitempty" db:"schedule_id"`
	Name *string `json:"name,omitempty" db:"name"`
	Done *bool `json:"done,omitempty" db:"done"`
	ImgUrl string `json:"imgUrl,omitempty" db:"-"`
	ImgUUID string `json:"-" db:"imguuid"`
	Order int `json:"orderPlace,omitempty" db:"orderplace"`
	Duration *int `json:"duration" db:"duration"`
}

type CardBeforeAfter struct {
	ID int `json:"id,omitempty" db:"id"`
	Schedule_ID int `json:"schedule_id,omitempty" db:"schedule_id"`
	Name *string `json:"name,omitempty" db:"name"`
	Done *bool `json:"done,omitempty" db:"done"`
	ImgUrl string `json:"imgUrl,omitempty" db:"-"`
	ImgUUID string `json:"-" db:"imguuid"`
	Order int `json:"orderPlace,omitempty" db:"orderplace"`
}

type PersonalImage struct {
	ID int `json:"id,omitempty" db:"id"`
	Mentor_ID int `json:"mentor_id,omitempty" db:"mentor_id"`
	ImgUrl string `json:"imgUrl,omitempty" db:"-"`
	ImgUUID string `json:"-" db:"imguuid"`
}

type StockImage struct {
	ID int `json:"id,omitempty" db:"id"`
	Category string `json:"category,omitempty" category:"id"`
	Names []string `json:"names,omitempty" db:"-"` 
	Name string `json:"name,omitempty" db:"name"`
	ImgUrl string `json:"imgUrl,omitempty" db:"-"`
	ImgUUID string `json:"-" db:"imguuid"`
}

