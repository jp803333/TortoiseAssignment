package model

import "time"

type Profile struct {
	Profileid   uint      `json:"profileid" gorm:"primary_key"`
	Name        string    `json:"name"`
	Dateofbirth time.Time `json:"dateofbirth"`
	Status      string    `json:"status"`
}
