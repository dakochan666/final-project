package models

import (
	"time"

	"final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Id        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"not null;uniqueIndex" json:"username" valid:"required~username is required"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" valid:"required~email is required,email~Invalid format email"`
	Password  string     `gorm:"not null" json:"password" valid:"required~password is required,minstringlength(6)~password has to have minimum length of 6 characters"`
	Age       int        `gorm:"not null" json:"age" valid:"required~age is required,numeric~fill age with number,range(8|99)~minimum 8 years old"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Photo     []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comment   []Comment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Social    []Social   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		return errCreate
	}

	hash, err := helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hash
	return
}
