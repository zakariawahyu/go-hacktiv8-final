package entity

import "time"

type SocialMedia struct {
	ID             int64     `gorm:"primaryKey:auto_increment" json:"-"`
	Name           string    `gorm:"type:varchar(100)" json:"-"`
	SocialMediaUrl string    `gorm:"type:varchar(150)" json:"-"`
	UserID         int64     `gorm:"not null" json:"-"`
	User           User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"-"`
	CreatedAt      time.Time `gorm:"not null" json:"-"`
	UpdatedAt      time.Time `gorm:"not null, autoUpdateTime" json:"-"`
}
