package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Social struct {
	Id             uint       `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"not null;type:varchar(100)" json:"name" valid:"required~name is required"`
	SocialMediaUrl string     `gorm:"not null" json:"social_media_url" valid:"required~social_media_url is required"`
	UserId         uint       `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`

	User *User
}

func (s *Social) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		return errCreate
	}

	return
}
