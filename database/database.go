package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func DatabaseInit() {
	var err error

	DSN := "host=localhost user=postgres password=password dbname=db_tayoh_rewrite port=54321 sslmode=disable TimeZone=Asia/Jakarta"

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database")
}