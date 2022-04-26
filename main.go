package main

import (
	"dna-matching-api/controllers"
	"dna-matching-api/database"
	"dna-matching-api/entity"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/penyakit/create", controllers.CreatePenyakit).Methods("POST")
	router.HandleFunc("/penyakit/get", controllers.GetAllPenyakit).Methods("GET")
	router.HandleFunc("/penyakit/get/{nama}", controllers.GetPenyakitByNama).Methods("GET")
	router.HandleFunc("/penyakit/update/{nama}", controllers.UpdatePenyakitByNama).Methods("PUT")
	router.HandleFunc("/penyakit/delete/{nama}", controllers.DeletePenyakitByNama).Methods("DELETE")

	router.HandleFunc("/pemeriksaan/get", controllers.GetAllPemeriksaan).Methods("GET")
	router.HandleFunc("/pemeriksaan/create", controllers.CreatePemeriksaan).Methods("POST")
}

func initDB() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error loading .env file")
	}
	config :=
		database.Config{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.MigratePenyakit(&entity.Penyakit{})
	database.MigratePemeriksaan(&entity.Pemeriksaan{})
}
