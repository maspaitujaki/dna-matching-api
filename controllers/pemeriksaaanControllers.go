package controllers

import (
	"dna-matching-api/database"
	"dna-matching-api/entity"
	"dna-matching-api/stringMatching"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// GetAllPemeriksaan returns all pemeriksaan
func GetAllPemeriksaan(w http.ResponseWriter, r *http.Request) {
	var pemeriksaans []entity.Pemeriksaan
	database.Connector.Find(&pemeriksaans)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pemeriksaans)
}

func DeletePemeriksaan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var pemeriksaan entity.Pemeriksaan
	if result := database.Connector.Where("id = ?", key).Delete(&pemeriksaan); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pemeriksaan)
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
	var cek float32 = stringMatching.KmpMatch(pemeriksaan.Rantai, penyakit.Rantai)
	if cek < 1 {
		cek *= 100
		pemeriksaan.Hasil = false
		pemeriksaan.Prediksi = fmt.Sprintf("%.2f", cek) + "%"
		pemeriksaan.Tanggal = time.Now()
	} else {
		pemeriksaan.Hasil = true
		pemeriksaan.Prediksi = "100%"
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
