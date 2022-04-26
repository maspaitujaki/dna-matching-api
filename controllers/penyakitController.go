package controllers

import (
	"dna-matching-api/database"
	"dna-matching-api/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Get All penyakit
func GetAllPenyakit(w http.ResponseWriter, r *http.Request) {
	var penyakits []entity.Penyakit
	database.Connector.Find(&penyakits)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(penyakits)
}

//GetPenyakitByName return penyakit with specific name
func GetPenyakitByNama(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["nama"]

	var penyakit entity.Penyakit

	if result := database.Connector.Where("nama = ?", key).Find(&penyakit); result.Error != nil || penyakit.Nama == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(penyakit)
}

// CreatePenyakit creates penyakit
func CreatePenyakit(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var penyakit entity.Penyakit
	json.Unmarshal(requestBody, &penyakit)

	if result := database.Connector.Create(penyakit); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(penyakit)
}

//UpdatePenyakitByNama updates penyakit with respective nama
func UpdatePenyakitByNama(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var penyakit entity.Penyakit
	json.Unmarshal(requestBody, &penyakit)

	if result := database.Connector.Save(&penyakit); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(penyakit)
}

//DeletePenyakitByNama delete's penyakit with specific nama
func DeletePenyakitByNama(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["nama"]

	var penyakit entity.Penyakit
	database.Connector.Where("nama = ?", key).Delete(&penyakit)
	w.WriteHeader(http.StatusNoContent)
}
