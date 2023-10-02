package entity

import (
	"github.com/peempumrapee/sa-66-example/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-66.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	database.AutoMigrate(
		&Member{},
		&Creator{},
		&SoundType{},
		&Sound{},
		&History{},
		&Rating{},
		&Playlist{},
	)
	db = database

	// Mockup data
	member := Member{
		Username: "sa-test",
		Password: "test1234",
		Email:    "test@mail.com",
	}
	db.Model(&Member{}).Create(&member)

	creator := Creator{
		Username: "sa-test",
		Password: "test1234",
		Email:    "test@mail.com",
	}
	db.Model(&Creator{}).Create(&creator)

	sound1 := Sound{
		Title:   "test-sound",
		Path:    "path/to/file",
		Creator: creator,
	}
	db.Model(&Sound{}).Create(&sound1)

	sound2 := Sound{
		Title:   "test-sound2",
		Path:    "path/to/file",
		Creator: creator,
	}
	db.Model(&Sound{}).Create(&sound2)

	var sounds []entity.Sound
	append(sounds, sound1)
	append(sounds, sound2)
	playlist := Playlist{
		Title: "new-playlist",
		Sound: sounds,
	}
	db.Model(&Playlist{}).Create(&playlist)
}
