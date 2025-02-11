package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StudentID string `valid:"required~StudentID is required, matches(^[BMD]\\d{7}$)"`
	FirstName string `valid:"required~FirstName is required"`
	LastName  string `valid:"required~LastName is required"`
	Email     string `valid:"required~Email is required, email~Email is invalid"`
	Phone     string `valid:"required~Phone is required, stringlength(10|10)"`
	Profile   string `gorm:"type:longtext"`
	LinkedIn  string `valid:"required~LinkedIn is required, url~Url LinkedIn is invalid"`
	// GenderID ทำหน้าที่เป็น FK
	GenderID uint   `valid:"required~Gender is required"`
	Gender   Gender `gorm:"foreignKey:GenderID"`
}

type Gender struct {
	gorm.Model
	Name string
}


type Member struct {
	gorm.Model
	PhoneNumber string `gorm:"uniqueIndex" valid:"required~Phone is required, matches(^[0]\\d{9}$)~PhoneNumber length is not 10 digits."`
	Password string 
	Url string `valid:"url~Url is invalid"`
}

