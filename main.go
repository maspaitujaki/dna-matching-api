package main

import (
	"dna-matching-api/controllers"
	"dna-matching-api/database"
	"dna-matching-api/entity"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "https://awokwok-dna-matching.herokuapp.com"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/penyakit/create", controllers.CreatePenyakit).Methods("POST")
	router.HandleFunc("/penyakit/get", controllers.GetAllPenyakit).Methods("GET")
	router.HandleFunc("/penyakit/get/{nama}", controllers.GetPenyakitByNama).Methods("GET")
	router.HandleFunc("/penyakit/update/{nama}", controllers.UpdatePenyakitByNama).Methods("PUT")
	router.HandleFunc("/penyakit/delete/{nama}", controllers.DeletePenyakitByNama).Methods("DELETE")

	router.HandleFunc("/pemeriksaan/get", controllers.GetAllPemeriksaan).Methods("GET")
	router.HandleFunc("/pemeriksaan/get/what", controllers.GetPemeriksaanByPenyakit).Methods("GET")
	router.HandleFunc("/pemeriksaan/get/when", controllers.GetPemeriksaanByTanggal).Methods("GET")
	router.HandleFunc("/pemeriksaan/get/whenwhat", controllers.GetPemeriksaanByPenyakitAndTanggal).Methods("GET")
	router.HandleFunc("/pemeriksaan/create", controllers.CreatePemeriksaan).Methods("POST")
	router.HandleFunc("/pemeriksaan/delete/{id}", controllers.DeletePemeriksaan).Methods("DELETE")
}

func initDB() {
	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	panic("Error loading .env file")
	// }
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
