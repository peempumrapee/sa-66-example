package entity

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	PlayedAt time.Time

	SoundID *uint
	Sound   Sound `gorm:"foreignKey:SoundID"`

	MemberID *uint
	Member   Member `gorm:"foreignKey:MemberID"`
}
