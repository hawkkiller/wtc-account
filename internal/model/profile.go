package model

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model `json:"-"`
	Username   string `json:"username" validate:"required" gorm:"unique"`
	Password   string `json:"password" validate:"required"`
	Email      string `json:"email" validate:"required" gorm:"index:unique"`
	Sex        string `json:"sex" gorm:"default:undefined"`
}

type UserProfileDB struct {
	gorm.Model `json:"-"`
	Username   string `json:"username" validate:"required" gorm:"unique"`
	Password   string `json:"-" validate:"required"`
	Email      string `json:"email" validate:"required" gorm:"index:unique"`
	Sex        string `json:"sex" gorm:"default:undefined"`
}

func (profile UserProfile) ProfileToDB() UserProfileDB {
	return UserProfileDB{
		Model:    profile.Model,
		Username: profile.Username,
		Password: profile.Password,
		Email:    profile.Email,
		Sex:      profile.Sex,
	}
}
