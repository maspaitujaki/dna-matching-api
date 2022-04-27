package controllers

import (
	"dna-matching-api/database"
	"dna-matching-api/entity"
	"dna-matching-api/stringMatching"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
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
	var penyakit entity.Penyakit
	key := database.Connector.Where("nama = ?", pemeriksaan.Penyakit).First(&penyakit)

	if key.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// // kalo ada lanjut ke bawah

	//TODO : Pemeriksaan rantai
	var cek int = stringMatching.KmpMatch(pemeriksaan.Rantai, penyakit.Rantai)
	if cek == -1 {
		pemeriksaan.Hasil = false
		pemeriksaan.Tanggal = time.Now()
	} else {
		pemeriksaan.Hasil = true
		pemeriksaan.Tanggal = time.Now()
	}
	// Penambahan ke database
	if result := database.Connector.Create(&pemeriksaan); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pemeriksaan)
}
