package controllers

import (
	"dna-matching-api/database"
	"dna-matching-api/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetAllPemeriksaan returns all pemeriksaan
func GetAllPemeriksaan(w http.ResponseWriter, r *http.Request) {
	var pemeriksaans []entity.Pemeriksaan
	database.Connector.Find(&pemeriksaans)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pemeriksaans)
}

// CreatePemeriksaan creates pemeriksaan
func CreatePemeriksaan(w http.ResponseWriter, r *http.Request) {
	// TODO : Implement CreatePemeriksaan dengan KMP dan Boyer-Moore
	requestBody, _ := ioutil.ReadAll(r.Body)
	var pemeriksaan entity.Pemeriksaan
	json.Unmarshal(requestBody, &pemeriksaan)

	// TODO: Cek penyakit ada nggak
	// pemeriksaan.Penyakit
	// // kalo ada lanjut ke bawah

	//TODO : Pemeriksaan rantai

	// TODO : Assign Hasil dan tanggal()
	// pemeriksaan.Hasil =
	// pemeriksaan.Tanggal = time.Now()

	// Penambahan ke database
	if result := database.Connector.Create(&pemeriksaan); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pemeriksaan)
}
