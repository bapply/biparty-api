package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	CreatedAt time.Time
}
