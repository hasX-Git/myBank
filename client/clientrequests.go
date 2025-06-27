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
	var newClient ClientInfo
	newClient.Firstn = newAccountRequest.Firstn
	newClient.Lastn = newAccountRequest.Lastn
	newClient.NID = newAccountRequest.NID

	var newAccount Account
	newAccount.Personinfo = newClient

	//generating unique Account ID(AID)
	var aid string
	for {
		aid = "AID" + currentDateAsID(5)
		if _, isIn := accounts[aid]; !isIn {
			break
		}
	}
	newAccount.AID = aid
	newClient.AID = aid

	DB.Create(&newAccount)
	DB.Create(&newClient)
	clients[newClient.NID] = &newClient
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

	_, isPresent := accounts[newTransactionRequest.AID]
	if !isPresent {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The account with such AID doesn't exist"})
		return
	}

	//checking if enough balance on account
	if newTransactionRequest.Sum > accounts[newTransactionRequest.AID].Balance {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "not enough balance"})
		return
	}

	var newtr Transaction
	newtr.Sum = newTransactionRequest.Sum
	newtr.AID = newTransactionRequest.AID
	newtr.TrID = "TID" + currentDateAsID(5)

	DB.Create(&newtr)
	transactions[newtr.TrID] = &newtr

	newBalance := accounts[newTransactionRequest.AID].Balance - newtr.Sum
	accounts[newTransactionRequest.AID].Balance = newBalance
	DB.Model(&Account{}).Where("id = ?", newtr.AID).Update("Balance", newBalance)

	accounts[newTransactionRequest.AID].Trs = append(accounts[newTransactionRequest.AID].Trs, newtr)

	c.IndentedJSON(http.StatusCreated, newtr)
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

	if _, isIn := accounts[newDepositRequest.AID]; !isIn {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account doesn't exist"})
		return
	}

	newBalance := accounts[newDepositRequest.AID].Balance + newDepositRequest.Sum
	accounts[newDepositRequest.AID].Balance = newBalance
	DB.Model(&Account{}).Where("id = ?", newDepositRequest.AID).Update("Balance", newBalance)

	c.IndentedJSON(http.StatusOK, gin.H{
		"Deposited": newDepositRequest.Sum,
		"Balance":   accounts[newDepositRequest.AID].Balance,
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

func DBGetAccountInfoByAID(c *gin.Context) {
	var acc Account
	AID := c.Param("account")
	result := DB.First(&acc, AID)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, acc)
}

func DBGetTransactionInfoByTID(c *gin.Context) {
	var tr Transaction
	TID := c.Param("transaction")
	result := DB.First(&tr, TID)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, tr)
}
