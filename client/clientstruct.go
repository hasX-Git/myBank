package client

type Transaction struct {
	Sum  uint32
	TrID string
}

type ClientInfo struct {
	Firstn string
	Lastn  string
	NID    string
}

type Account struct {
	Balance    uint32
	AID        string
	Trs        map[string]*Transaction
	Personinfo ClientInfo
}

var Transactions = make(map[string]*Transaction)
var Accounts = make(map[string]*Account)
var Clients = make(map[string]*ClientInfo)
