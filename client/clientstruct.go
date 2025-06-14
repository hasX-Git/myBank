package client

type transaction struct {
	Sum  uint32
	TrID string
}

type clientInfo struct {
	Firstn string
	Lastn  string
	NID    string
}

type account struct {
	Balance    uint32
	AID        string
	Trs        map[string]*transaction
	Personinfo clientInfo
}

var transactions = make(map[string]*transaction)
var accounts = make(map[string]*account)
var clients = make(map[string]*clientInfo)
