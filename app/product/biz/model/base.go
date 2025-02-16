package model

import "time"

type Base struct {
	ID        int `gorm:"primarykey"`
	CreateAte time.Time
	Update    time.Time
}
