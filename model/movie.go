package model

import "github.com/jinzhu/gorm"

type Movie struct {
	gorm.Model
	Movie_id int
	Name string
	En_name string
	Category string
	Cover_url string
	Area string
	Release_date int
	Info string
	Duration int
}
