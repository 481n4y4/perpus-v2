package models

type Book struct {
	Id           int    `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	Author       string `json:"author"`
	Publisher    string `json:"publisher"`
	Publish_date string `json:"publish_date"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
}
