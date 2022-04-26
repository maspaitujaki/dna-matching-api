package entity

type Penyakit struct {
	Nama   string `json:"nama" gorm:"primary_key;not null"`
	Rantai string `json:"rantai"`
}
