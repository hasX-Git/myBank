package client

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	AID  string `gorm:"unique;column:AID"`
	Sum  uint32 `gorm:"column:Amount"`
	TrID string `gorm:"unique;column:TID"`
}

type ClientInfo struct {
	gorm.Model
	AID    string `gorm:"unique;column:AID"`
	Firstn string `gorm:"column:First_Name"`
	Lastn  string `gorm:"column:Last_Name"`
	NID    string `gorm:"unique;column:NID"`
}

type Account struct {
	gorm.Model
	Balance    uint32        `gorm:"column:Balance"`
	AID        string        `gorm:"unique;column:AID"`
	Trs        []Transaction `gorm:"foreignKey:AID;column:Transactions"`
	PersonInfo ClientInfo    `gorm:"foreignKey:AID;column:Personal_Info"`
}
