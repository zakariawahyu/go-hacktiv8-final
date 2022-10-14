package entity

import "time"

type Comment struct {
	ID        int64     `gorm:"primaryKey:auto_increment" json:"-"`
	UserID    int64     `gorm:"not null" json:"-"`
	User      User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"-"`
	PhotoID   int64     `gorm:"not null" json:"-"`
	Photo     Photo     `gorm:"foreignKey:PhotoID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"-"`
	Message   string    `gorm:"type:text" json:"-"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null, autoUpdateTime" json:"-"`
}
