package models

import "time"

type Temperatura struct {
	ID          uint      `gorm: "primarykey; autoIncrement:true" `
	Temperatura string    `json: "temperatura" `
	Data        time.Time `json: "data"`
}
