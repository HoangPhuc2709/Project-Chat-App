package models

import (
	"time"
)

type User struct {
	UUID        string `gorm:"type:char(36);primaryKey"`
	ID          uint   `gorm:"unique;autoIncrement"`
	Username    string `gorm:"type:varchar(50);unique;index;check username NOT LIKE '% %'"`
	Password    string `gorm:"type:varchar(60)"`
	Gender      bool
	FirstName   string    `gorm:"type:nvarchar(50)"`
	LastName    string    `gorm:"type:nvarchar(50)"`
	Email       string    `gorm:"type:nvarchar(100);unique,check:email LIKE '%@%'"`
	PhoneNumber string    `gorm:"type:varchar(15);unique;check:phone_number REGEXP '^[0+][0-9]{6,}$'"`
	BirthDay    time.Time `gorm:"type:date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (user *User) GenderStr() string {
	if user.Gender {
		return "male"
	}
	return "female"
}
