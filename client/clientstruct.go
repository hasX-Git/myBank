package client

type Transaction struct {
	AID  string
	Sum  uint32
	TrID string `gorm:"primaryKey"`
}

type ClientInfo struct {
	AID    string `gorm:"unique"`
	Firstn string
	Lastn  string
	NID    string `gorm:"primaryKey"`
}

type Account struct {
	Balance    uint32
	AID        string        `gorm:"primaryKey"`
	Trs        []Transaction `gorm:"foreignKey:AID"`
	Personinfo ClientInfo    `gorm:"foreignKey:AID"`
}

var transactions = make(map[string]*Transaction)
var accounts = make(map[string]*Account)
var clients = make(map[string]*ClientInfo)
