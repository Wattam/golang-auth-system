package models

import (
	"time"

	"gorm.io/gorm"
)

type Shoe struct {
	ID        uint           `json:"id,string"`
	Name      string         `json:"name" gorm:"not null; default:null; check:Name != '' "`
	Style     string         `json:"style" gorm:"not null; default:null; check:Style != '' "`
	Colour    string         `json:"colour" gorm:"not null; default:null; check:Colour != '' "`
	Material  string         `json:"material" gorm:"not null; default:null; check:Material != '' "`
	Price     float64        `json:"price,string" gorm:"not null; default:null; check:Price > 0 "`
	CreatedAt time.Time      `json:"-" gorm:"<-:create"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
