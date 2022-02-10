package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id,string"`
	Username  string         `json:"username" gorm:"unique; not null; default:null; check:Username != '' "`
	Email     string         `json:"email" gorm:"unique; not null; default:null; check:Email != '' "`
	Password  string         `json:"password" gorm:"not null; default:null; check:Password != '' "`
	CreatedAt time.Time      `json:"-" gorm:"<-:create"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
