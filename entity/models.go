package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string    `gorm:"type:varchar(55);not null"`
	Username string    `gorm:"type:varchar(55);not null;unique"`
	Email    string    `gorm:"type:varchar(35);not null;unique"`
	Phone    string    `gorm:"type:varchar(15);not null;unique"`
	Password string    `gorm:"type:varchar(255);not null"`
	Comments []Comment `gorm:"foreignkey:UserID"`
	Events   []Event   `gorm:"foreignkey:UserID"`
}

type Comment struct {
	gorm.Model
	Comment string
	UserID  uint
	EventID uint
}

type Event struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	Rules       string    `gorm:"type:text"`
	Organizer   string    `gorm:"type:varchar(55)"`
	DueDate     float64   `gorm:"type:number(15);not null"`
	BeginAt     time.Time `gorm:"type:time;not null"`
	Location    string    `gorm:"type:varchar(55);not null"`
	Ticket      int       `gorm:"type:int(5);not null"`
	Links       string    `gorm:"type:varchar(255)"`
	Banner      string    `gorm:"type:varchar(255);not null"`
	UserID      uint
	Comments    []Comment `gorm:"foreignkey:UserID"`
}
