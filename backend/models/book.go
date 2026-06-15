package models

import "time"

type Book struct {
	Id          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"publish_date" gorm:"type:date"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
}
