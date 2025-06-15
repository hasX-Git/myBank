package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	//initializing new client
	var newAccountRequest createAccountRequest
	if err := c.BindJSON(&newAccountRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message ": "Improper formatting"})
		return
	}

	//checking validity of NID
	if !checkValidityOfID(newAccountRequest.NID, 12) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message ": "Invalid NID"})
		return
	}

	//checking if the client already exists
	_, isPresent := clients[newAccountRequest.NID]
	if isPresent {
		c.IndentedJSON(http.StatusConflict, gin.H{"message ": "The account with such National ID already exists"})
		return
	}

	//creating Client and Account
	var newClient clientInfo
	newClient.Firstn = newAccountRequest.Firstn
	newClient.Lastn = newAccountRequest.Lastn
	newClient.NID = newAccountRequest.NID
	clients[newClient.NID] = &newClient

	var newAccount account
	newAccount.Personinfo = newClient
	newAccount.Balance = 0
	newAccount.Trs = make(map[string]*transaction)

	//generating unique Account ID(AID)
	var aid string
	for {
		aid = currentDateAsID(5)
		if _, isIn := accounts[aid]; !isIn {
			break
		}
	}
	newAccount.AID = aid
	accounts[newAccount.AID] = &newAccount

	//
	c.IndentedJSON(http.StatusCreated, newAccount)
}

func CreateTransaction(c *gin.Context) {
	var newTransactionRequest createTransactionRequest

	if err := c.BindJSON(&newTransactionRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Improper formatting"})
		return
	}

	_, isPresent := accounts[newTransactionRequest.Aid]
	if !isPresent {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The account with such AID doesn't exist"})
		return
	}

	var newtr transaction
	newtr.Sum = newTransactionRequest.Sum

	//checking if enough balance on account
	if newtr.Sum > accounts[newTransactionRequest.Aid].Balance {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "not enough balance"})
		return
	}

	newtr.TrID = currentDateAsID(5)
	transactions[newtr.TrID] = &newtr
	accounts[newTransactionRequest.Aid].Trs[newtr.TrID] = &newtr
	accounts[newTransactionRequest.Aid].Balance -= newtr.Sum

	c.IndentedJSON(http.StatusCreated, accounts[newTransactionRequest.Aid].Trs[newtr.TrID])
}

func GetAccountInfoByAID(c *gin.Context) {
	AID := c.Param("account")
	Account, err := findAccByAID(AID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, Account)
}

func GetTransactionInfoByTID(c *gin.Context) {
	TID := c.Param("transaction")
	Transaction, err := findTrByTID(TID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, Transaction)
}

func DepositMoney(c *gin.Context) {
	var newDepositRequest depositRequest

	if err := c.BindJSON(&newDepositRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Improper formatting"})
		return
	}

	if _, isIn := accounts[newDepositRequest.Aid]; !isIn {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account doesn't exist"})
		return
	}

	accounts[newDepositRequest.Aid].Balance += newDepositRequest.Sum
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"Deposited": newDepositRequest.Sum,
		"Balance":   accounts[newDepositRequest.Aid].Balance,
	})
}

func GetAccountsList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, accounts)
}

func GetClientsList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clients)
}

func GetTransacitonsList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, transactions)
}
