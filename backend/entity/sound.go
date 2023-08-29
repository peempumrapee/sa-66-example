package entity

import "gorm.io/gorm"

type Sound struct {
  gorm.Model
  Title string
  Path  string

  SoundTypeID *uint
  // For Automate joining data purpose 
  SoundType SoundType `gorm:"foreignKey:SoundTypeID"`

  CreatorID *uint
  Creator Creator `gorm:"foreignKey:CreatorID"`

  Histories []History `gorm:"foreignKey:SoundID"`
  Ratings  []Rating `gorm:"foreignKey:SoundID"`
  SoundPlaylists []SoundPlaylist  `gorm:"foreignKey:SoundID"`
}
