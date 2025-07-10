package client

type Transaction struct {
	AID  string `gorm:"column:aid"`
	Sum  uint32 `gorm:"column:amount"`
	TrID string `gorm:"primaryKey;column:tid"`
}

type ClientInfo struct {
	AID    string `gorm:"unique;column:aid"`
	Firstn string `gorm:"column:first_name"`
	Lastn  string `gorm:"column:last_name"`
	NID    string `gorm:"primaryKey;column:nid"`
}

type Account struct {
	Balance    uint32        `gorm:"column:balance"`
	AID        string        `gorm:"primaryKey;column:aid"`
	Trs        []Transaction `gorm:"foreignKey:AID"`
	PersonInfo ClientInfo    `gorm:"foreignKey:AID;references:aid"`
}

type File struct {
	Filename string `gorm:"column:filename"`
	Hash     string `gorm:"column:hash;unique"`
}
