package entity

import "gorm.io/gorm"

type SoundType struct {
	gorm.Model
	Name string

	Sounds []Sound `gorm:"foreignKey:SoundTypeID"`
}
