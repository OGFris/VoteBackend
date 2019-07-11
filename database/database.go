//		Copyright (c) VoteBackend - All Rights Reserved
//
//	Unauthorized copying of this file, via any medium is strictly prohibited
//	Proprietary and confidential
//	Written by Ilyes Cherfaoui <ogfris@protonmail.com>, 2019

package database

import (
	"github.com/OGFris/VoteBackend/utils"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

var Instance *gorm.DB

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"Default:null" sql:"index" json:"deleted_at"`
}

func Init() {
	if Instance == nil {
		var d *gorm.DB

		sql := os.Getenv("DB_USER") + ":" +
			os.Getenv("DB_PASSWORD") + "@tcp(" +
			os.Getenv("DB_ADDRESS") + ")/" +
			os.Getenv("DB_NAME")

		d, err := gorm.Open("mysql", sql+"?charset=utf8&parseTime=True&loc=Local")
		utils.PanicError(err)
		Instance = d

		Instance.AutoMigrate(
			&Candidate{},
		)
	}
}
