package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "35.240.191.214"
	dbname   = "fillgoods-lab"
	password = "IkDD43cIamwavO7s"
	port     = "5432"
	user     = "postgres"
)

var DB *gorm.DB

func InitialSqliteDatabase() error {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		host, port, user, dbname, password)

	DB, err = gorm.Open("postgres", psqlInfo)

	autoMigrate()

	return err
}

func autoMigrate() {
	DB.AutoMigrate(&User{})
}
