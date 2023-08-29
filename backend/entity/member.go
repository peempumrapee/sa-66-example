package entity

import "gorm.io/gorm"

type Member struct {
  gorm.Model
  Username  string  `gorm:"uniqueIndex"`
  Password  string
  Email     string  `gorm:"uniqueIndex"`

  Histories []History `gorm:"foreignKey:MemberID"`
  Ratings []Rating  `gorm:"foreignKey:MemberID"`
  Playlists []Playlist  `gorm:"foreignKey:MemberID"`
}
