package models

type Card struct {
	ID int `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`
	Done bool `json:"done,omitempty" db:"done"`
	ImgUrl string `json:"imgUrl,omitempty" db:"imgurl"`
	StartTime string `json:"startTime,omitempty" db:"starttime"`
	EndTime string `json:"endTime,omitempty" db:"endtime"`
}