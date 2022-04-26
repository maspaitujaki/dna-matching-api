package database

import (
	"dna-matching-api/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return err
	}
	log.Println("Connected to database")
	return nil
}

func MigratePenyakit(table *entity.Penyakit) {
	Connector.AutoMigrate(&table)
	log.Println("Migrated table penyakit")
}

func MigratePemeriksaan(table *entity.Pemeriksaan) {
	Connector.AutoMigrate(&table)
	log.Println("Migrated table pemeriksaan")
}
