package model

import (
	"time"
)

type Todo struct {
	Id        int       `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	Task      string    `gorm:"NOT NULL"`
	Done      bool      `gorm:"NOT NULL;DEFAULT:false"`
	CreatedAt time.Time `gorm:"NOT NULL"`
	UpdatedAt time.Time
}