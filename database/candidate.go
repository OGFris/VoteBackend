//		Copyright (c) VoteBackend - All Rights Reserved
//
//	Unauthorized copying of this file, via any medium is strictly prohibited
//	Proprietary and confidential
//	Written by Ilyes Cherfaoui <ogfris@protonmail.com>, 2019

package database

import (
	"github.com/OGFris/VoteBackend/utils"
	"github.com/jinzhu/gorm"
)

type Candidate struct {
	Model
	Name        string `gorm:"Type:varchar(255);Column:name;unique;NOT NULL" json:"name"`
	Party       string `gorm:"Type:varchar(255);Column:party;NOT NULL" json:"party"`
	Description string `gorm:"Type:mediumtext;Column:description;NOT NULL" json:"description"`
}

func Exist(name string) bool {
	return !gorm.IsRecordNotFoundError(Instance.Find(&Candidate{}, &Candidate{Name: name}).Error)
}

func Create(name, party, description string) error {

	return Instance.Create(&Candidate{
		Name:        name,
		Party:       party,
		Description: description,
	}).Error
}

func AllCandidates() (c []Candidate) {
	utils.PanicError(Instance.Find(&c).Error)

	return
}
