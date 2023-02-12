package model

import "time"

type Cashier struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name" gorm:"unique"`
	Passcode  string    `json:"passcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
