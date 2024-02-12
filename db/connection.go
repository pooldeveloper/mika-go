package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=roundhouse.proxy.rlwy.net user=postgres password=-1C1eBD1bced4EFd-a-6ae2bF4-dG-db dbname=mika_db port=30558"

var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		println("DB conectada")
	}
}
