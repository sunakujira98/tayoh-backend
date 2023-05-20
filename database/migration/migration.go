package migration

import (
	"fmt"
	"log"
	"tayoh-backend/database"
	"tayoh-backend/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}