package entity

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey:auto_increment" json:"-"`
	Username  string    `gorm:"type:varchar(100)" json:"-"`
	Email     string    `gorm:"type:varchar(100)" json:"-"`
	Password  string    `gorm:"type:varchar(100)" json:"-"`
	Age       int       `gorm:"type:integer" json:"-"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null, autoUpdateTime" json:"-"`
}
