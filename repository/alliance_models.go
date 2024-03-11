package repository

import (
	"time"
)

type Satellite struct {
	ID       uint64 `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"type:varchar(255);unique;not null"`
	X        int    `gorm:"not null"`
	Y        int    `gorm:"not null"`
	Messages []Message
}

type Message struct {
	ID          uint64    `gorm:"primary_key;auto_increment"`
	Distance    float64   `gorm:"type:numeric(20,10);not null"`
	Date        time.Time `gorm:"not null"`
	SatelliteID uint64
	Satellite   Satellite `gorm:"foreignKey:SatelliteID"`
	Words       []Word
}

type Word struct {
	ID        uint64 `gorm:"primary_key;auto_increment"`
	Word      string `gorm:"type:varchar(255);not null"`
	Position  int    `gorm:"not null"`
	MessageID uint64
	Message   Message `gorm:"foreignKey:MessageID"`
}
