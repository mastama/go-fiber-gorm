package model

import "time"

type Category struct {
	Id        int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT; primary_key;"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
