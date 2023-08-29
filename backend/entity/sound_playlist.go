package entity

import "gorm.io/gorm"

type SoundPlaylist struct {
  gorm.Model

  PlaylistID  *uint
  Playlist  Playlist  `gorm:"foreignKey:PlaylistID"`

  SoundID *uint
  Sound Sound `gorm:"foreignKey:SoundID"`
}
