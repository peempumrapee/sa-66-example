package entity

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Score int

	SoundID *uint
	Sound   Sound `gorm:"foreignKey:SoundID"`

	MemberID *uint
	Member   Member `gorm:"foreignKey:MemberID"`
}
