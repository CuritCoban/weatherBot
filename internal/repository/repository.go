package repository

import "gorm.io/gorm"

type DataBase struct {
	Postgres *gorm.DB
}
