package entity

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	Title string

	MemberID *uint
	Member   Member `gorm:"foreignKey:MemberID"`

	Sound []Sound `gorm:"many2many:sound_playlist;"`
}
