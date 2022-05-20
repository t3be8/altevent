package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string    `gorm:"type:varchar(55);not null" json:"fullname"`
	Username string    `gorm:"type:varchar(55);not null;unique" json:"username"`
	Email    string    `gorm:"type:varchar(35);not null;unique" json:"email"`
	Phone    string    `gorm:"type:varchar(15);not null;unique" json:"phone"`
	Password string    `gorm:"type:varchar(255);not null" json:"password"`
	Comments []Comment `gorm:"foreignkey:UserID"`
	Events   []Event   `gorm:"foreignkey:UserID"`
}

type Comment struct {
	gorm.Model
	Comment string `json:"comment"`
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
}

type Event struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Rules       string    `gorm:"type:text" json:"rules"`
	Organizer   string    `gorm:"type:varchar(55)" json:"organizer"`
	DueDate     string    `json:"due_date"`
	BeginAt     string    `json:"begin_at"`
	Location    string    `gorm:"type:varchar(55);not null" json:"location"`
	Ticket      int       `gorm:"type:int(5);not null" json:"ticket"`
	Links       string    `gorm:"type:varchar(255)" json:"links"`
	Banner      string    `gorm:"type:varchar(255);not null" json:"banner"`
	UserID      uint      `json:"user_id"`
	Comments    []Comment `gorm:"foreignkey:EventID;references:id"`
}

type Attendee struct {
	gorm.Model
	Ticket  int  `json:"ticket"`
	UserID  uint `json:"user_id"`
	EventID uint `json:"event_id"`
}
