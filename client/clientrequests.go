package client

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// POST
func POSTcreateAccount(c *gin.Context) {
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

	result := DB.First(&ClientInfo{}, "NID = ?", newAccountRequest.NID)

	if result.Error == nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Client already in database"})
		return
	}

	//creating Client and Account
	var newClient ClientInfo
	newClient.Firstn = newAccountRequest.Firstn
	newClient.Lastn = newAccountRequest.Lastn
	newClient.NID = newAccountRequest.NID

	var newAccount Account
	newAccount.PersonInfo = newClient

	//generating unique Account ID(AID)
	var aid string
	for {
		aid = "AID" + currentDateAsID(5)
		result = DB.First(&Account{}, "AID = ?", aid)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			break
		}
	}
	newAccount.AID = aid
	newClient.AID = aid

	result = DB.Create(&newClient)
	if result.Error != nil {
		log.Println("Error:", result)
		c.IndentedJSON(http.StatusNotImplemented, "updating client db failed")
		return
	}
	result = DB.Create(&newAccount)
	if result.Error != nil {
		log.Println("Error:", result)
		c.IndentedJSON(http.StatusNotImplemented, "updating account db failed")
		return
	}

	//
	c.IndentedJSON(http.StatusCreated, newAccount)
}

func POSTcreateTransaction(c *gin.Context) {
	var newTransactionRequest createTransactionRequest

	if Err := c.BindJSON(&newTransactionRequest); Err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Improper formatting"})
		return
	}

	result := DB.First(&Account{}, "AID = ?", newTransactionRequest.AID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The account with such AID doesn't exist"})
			log.Println("Error:", result)
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured when searching AID"})
			log.Println("Error:", result)
		}
		return
	}
	//checking if enough balance on account
	var balance uint32
	result = DB.Model(&Account{}).Select("balance").Where("AID = ?", newTransactionRequest.AID).Scan(&balance)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured when retrieving balance"})
		log.Println("Error:", result)
	}

	if newTransactionRequest.Sum > balance {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "not enough balance"})
		return
	}

	var newtr Transaction
	newtr.Sum = newTransactionRequest.Sum
	newtr.AID = newTransactionRequest.AID
	newtr.TrID = "TID" + currentDateAsID(5)

	DB.Create(&newtr)

	newBalance := balance - newTransactionRequest.Sum

	result = DB.Model(&Account{}).Where("AID = ?", newTransactionRequest.AID).Update("balance", newBalance)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured when updating Balance"})
		log.Println("Error:", result.Error)
		return
	}

	c.IndentedJSON(http.StatusCreated, newtr)
}

// GET
func GETaccountInfoByAID(c *gin.Context) {
	AID := c.Param("account")
	var acc Account

	result := DB.Preload("PersonInfo").Preload("Trs").First(&acc, "AID = ?", AID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured"})
			log.Println("Error:", result.Error)
		}
		return
	}

	c.IndentedJSON(http.StatusOK, acc)
}

func GETtransactionInfoByTID(c *gin.Context) {
	TID := c.Param("transaction")
	var tr Transaction

	result := DB.First(&tr, "TID = ?", TID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured"})
			log.Println("Error:", result.Error)
		}
		return
	}

	c.IndentedJSON(http.StatusOK, tr)
}

func GETclientInfoByNID(c *gin.Context) {
	NID := c.Param("id")
	var cl ClientInfo

	result := DB.First(&cl, "NID = ?", NID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Client not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured"})
			log.Println("Error:", result.Error)
		}
		return
	}

	c.IndentedJSON(http.StatusOK, cl)
}

func GETaccountsList(c *gin.Context) {
	var accounts []Account

	result := DB.Preload("PersonInfo").Preload("Trs").Find(&accounts)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured when getting list"})
		log.Println("Error:", result)
		return
	}

	c.IndentedJSON(http.StatusOK, accounts)
}

func GETclientsList(c *gin.Context) {
	var clients []ClientInfo

	result := DB.Find(&clients)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured when getting list"})
		log.Println("Error:", result)
		return
	}

	c.IndentedJSON(http.StatusOK, clients)
}

func GETtransactionsList(c *gin.Context) {
	var transactions []Transaction

	result := DB.Find(&transactions)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured when getting list"})
		log.Println("Error:", result)
		return
	}

	c.IndentedJSON(http.StatusOK, transactions)
}

//PATCH

func PATCHdepositMoney(c *gin.Context) {
	var newDepositRequest depositRequest

	if err := c.BindJSON(&newDepositRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Improper formatting"})
		return
	}

	result := DB.First(&Account{}, "AID = ?", newDepositRequest.AID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The account with such AID doesn't exist"})
			log.Println("Error:", result)
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured when searching AID"})
			log.Println("Error:", result)
		}
		return
	}

	var balance uint32
	result = DB.Model(&Account{}).Select("balance").Where("AID = ?", newDepositRequest.AID).Scan(&balance)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured when retrieving balance"})
		log.Println("Error:", result)
	}

	newBalance := balance + newDepositRequest.Sum
	result = DB.Model(&Account{}).Where("AID = ?", newDepositRequest.AID).Update("balance", newBalance)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occured when updating Balance"})
		log.Fatal("Error:", result.Error)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"Deposited": newDepositRequest.Sum,
		"Balance":   newBalance,
	})
}

func PATCHhardDeleteAll(c *gin.Context) {
	DB.Unscoped().Where("1 = 1").Delete(&Transaction{})
	DB.Unscoped().Where("1 = 1").Delete(&ClientInfo{})
	DB.Unscoped().Where("1 = 1").Delete(&Account{})
}
