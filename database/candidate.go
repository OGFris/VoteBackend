//		Copyright (c) VoteBackend - All Rights Reserved
//
//	Unauthorized copying of this file, via any medium is strictly prohibited
//	Proprietary and confidential
//	Written by Ilyes Cherfaoui <ogfris@protonmail.com>, 2019

package database

type Candidate struct {
	Model
	Name        string `gorm:"Type:varchar(255);Column:name;unique;NOT NULL" json:"name"`
	Party       string `gorm:"Type:varchar(255);Column:party;NOT NULL" json:"party"`
	Description string `gorm:"Type:mediumtext;Column:description;NOT NULL" json:"description"`
}
