package entity

import "time"

type Pemeriksaan struct {
	Id       int       `json:"id"`
	Nama     string    `json:"nama"`
	Penyakit string    `json:"penyakit"`
	Tanggal  time.Time `json:"tanggal"`
	Rantai   string    `json:"rantai"`
	Hasil    bool      `json:"hasil"`
}
