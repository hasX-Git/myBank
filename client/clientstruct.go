package client

type Transaction struct {
	Tr   uint32
	TrID string
}

type ClientInfo struct {
	Firstn string
	Lastn  string
	NID    string
}

type Account struct {
	Balance    int32
	AID        string
	Trs        []Transaction
	Personinfo ClientInfo
}

var Transactions []Transaction
var Accounts []Account
var Clients []ClientInfo
