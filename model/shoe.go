package model

import (
	"time"

	"gorm.io/gorm"
)

type Shoe struct {
	ID        uint           `json:"id,string"`
	Name      string         `json:"name" gorm:"not null; default:null; check:Name != '' "`
	Color     string         `json:"color" gorm:"not null; default:null; check:Color != '' "`
	Price     float64        `json:"price,string" gorm:"not null; default:null; check:Price > 0 "`
	CreatedAt time.Time      `json:"-" gorm:"<-:create"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
