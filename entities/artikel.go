package entities

import "time"

type Artikel struct {
	Id        uint
	Judul     string
	Isi       string
	Foto      string
	Alamat    string
	Maps      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
