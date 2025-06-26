package client

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	dsn := "host=localhost user=postgres password=sctpwd dbname=bank_cont port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Conneciton to database failed")
	}

	if err = DB.AutoMigrate(&Account{}, &ClientInfo{}, &Transaction{}); err != nil {
		log.Fatal("migration failed")
	}

}
