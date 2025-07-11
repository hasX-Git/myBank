package client

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			log.Println("Connected")
			break
		}
		log.Println("Reconnecting...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Conneciton to database failed")
	}

	if err = DB.AutoMigrate(&Account{}, &ClientInfo{}, &Transaction{}, &File{}); err != nil {
		log.Fatal("migration failed")
	}

}
