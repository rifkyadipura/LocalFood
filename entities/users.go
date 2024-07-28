package entities

import "time"

type Users struct {
	Id_users  uint
	Role      string
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
