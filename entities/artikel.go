package entities

import "time"

type Artikel struct {
	Id_artikel uint
	Id_users   uint
	Judul      string
	Isi        string
	Foto       string
	Alamat     string
	Maps       string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
