package entity

import "gorm.io/gorm"

type Playlist struct {
  gorm.Model
  Title string

  MemberID  *uint
  Member  Member  `gorm:"foreignKey:MemberID"`

  SoundPlaylists  []SoundPlaylist `gorm:"foreignKey:PlaylistID"`
}
